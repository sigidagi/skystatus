package device

import (
	"fmt"
	"time"

	"github.com/sigidagi/skystatus/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

type Device struct {
	Port *serial.Port
	Tty  string
	Boud int
}

var device Device

func New() *Device {
	return &device
}

func (device Device) LightUp(led, color string) {

	dataToSend := fmt.Sprintf("%s_%s", led, color)
	bytesToSend := []byte(dataToSend)

	_, err := device.Port.Write(bytesToSend)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Has send " + dataToSend)
}

func blinkTest(dev *Device) {

	var indexArray = []string{"0", "1", "2", "3"}
	for _, index := range indexArray {
		dev.LightUp(index, "blue")
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)
}

func Setup(c config.Config) error {

	log.Info("Device is setting up")

	config := &serial.Config{
		Name:        c.Device.Name,
		Baud:        c.Device.Baud, // Set the baud rate according to your device's specifications.
		ReadTimeout: time.Second,
	}

	log.Info("Device baud rate is ", c.Device.Baud)
	log.Info("Device tty is ", c.Device.Name)

	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	device = Device{
		Tty:  c.Device.Name,
		Boud: c.Device.Baud,
		Port: port,
	}

	return nil
}

func Run() error {

	log.Info("Device is running")

	blinkTest(&device)
	return nil
}
