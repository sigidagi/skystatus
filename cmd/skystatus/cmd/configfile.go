package cmd

import (
	"html/template"
	"os"

	"github.com/sigidagi/skystatus/internal/config"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const configTemplate = `
[general]
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}
# Log to syslog.
#
# When set to true, log messages are being written to syslog.
log_to_syslog={{ .General.LogToSyslog }}

# Gateway backend configuration.
[hub]
publish={{ .Hub.Publish }}
subscribe={{ .Hub.Subscribe }}

[mqtt]
broker={{ .Mqtt.Broker }}
port={{ .Mqtt.Port }}
topic={{ .Mqtt.Topic }}
device={{ .Mqtt.Device }}
`

var configCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the Hub-WebSocket service configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
