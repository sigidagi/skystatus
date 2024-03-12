package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/sigidagi/skystatus/internal/config"
	"github.com/sigidagi/skystatus/internal/device"
	"github.com/sigidagi/skystatus/internal/runners"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string // config file
var version string

var rootCmd = &cobra.Command{
	Use:   "skystatus",
	Short: "SkyStatus: project status for IUI developers ",
	Long: `SkyStatus: project status for IUI developers . 
	> documentation & support: https://www.skyaalborg.io/
	> source & copyright information: https://skyaalborg.io/skystatus`,
	RunE: run,
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetDefault("general.log_level", 4)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

// Execute executes the root command.
func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		b, err := os.ReadFile(cfgFile)
		if err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
	} else {
		viper.SetConfigName("skystatus")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/skystatus/")
		viper.AddConfigPath("/etc/skystatus/")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
			default:
				log.WithError(err).Fatal("read configuration file error")
			}
		}
	}

	for _, pair := range os.Environ() {
		d := strings.SplitN(pair, "=", 2)
		if strings.Contains(d[0], ".") {
			log.Warning("Using dots in env variable is illegal and deprecated. Please use double underscore `__` for: ", d[0])
			underscoreName := strings.ReplaceAll(d[0], ".", "__")
			// Set only when the underscore version doesn't already exist.
			if _, exists := os.LookupEnv(underscoreName); !exists {
				os.Setenv(underscoreName, d[1])
			}
		}
	}

	viperBindEnvs(config.C)

	if err := viper.Unmarshal(&config.C); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
	}
}

func run(cmd *cobra.Command, args []string) error {

	tasks := []func() error{
		printStartMessage,
		setLogLevel,
		setupDevice,
		runTest,
		setupRunners,
		runProject,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
	log.Warning("shutting down service")

	return nil
}

func printStartMessage() error {
	log.WithFields(log.Fields{
		"version": version,
		"docs":    "https://skyaalborg.io/skystatus/",
	}).Info("starting SkyStatus service")
	return nil
}

func setLogLevel() error {
	log.SetLevel(log.Level(uint8(config.C.General.LogLevel)))
	return nil
}

func setupDevice() error {
	if err := device.Setup(config.C); err != nil {
		return errors.New(fmt.Sprintf("Device setup error: '%s'", err.Error()))
	}
	return nil
}

func setupRunners() error {
	if err := runners.Setup(config.C); err != nil {
		return errors.New(fmt.Sprintf("Runners setup error: '%s'", err.Error()))
	}
	return nil
}

func runTest() error {

	if err := device.Run(); err != nil {
		return err
	}
	return nil
}

func runProject() error {
	if err := runners.Run(); err != nil {
		return err
	}
	return nil
}

func viperBindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tv = strings.ToLower(t.Name)
		}
		if tv == "-" {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			viperBindEnvs(v.Interface(), append(parts, tv)...)
		default:
			// Bash doesn't allow env variable names with a dot so
			// bind the double underscore version.
			keyDot := strings.Join(append(parts, tv), ".")
			keyUnderscore := strings.Join(append(parts, tv), "__")
			viper.BindEnv(keyDot, strings.ToUpper(keyUnderscore))
		}
	}
}
