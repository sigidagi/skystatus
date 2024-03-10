package device

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

func lightUp(port *serial.Port, led, color string) {

	dataToSend := fmt.Sprintf("%s_%s", led, color)
	bytesToSend := []byte(dataToSend)

	_, err := port.Write(bytesToSend)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Has send " + dataToSend)
}

func Run() error {

	log.Info("Device is running")

	portName := "/dev/ttyACM2"

	config := &serial.Config{
		Name:        portName,
		Baud:        9600, // Set the baud rate according to your device's specifications.
		ReadTimeout: time.Second,
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer port.Close()

	// 0 LED - cpp check,
	lightUp(port, "0", "blue")

	time.Sleep(300 * time.Millisecond)

	lightUp(port, "1", "blue")

	time.Sleep(300 * time.Millisecond)

	lightUp(port, "2", "blue")

	time.Sleep(300 * time.Millisecond)

	lightUp(port, "3", "blue")
	/*
	 *    go func(failed bool) {
	 *        blink(port, 30, "2", failed)
	 *    }(false)
	 *
	 *    go func(failed bool) {
	 *        blink(port, 30, "3", failed)
	 *    }(true)
	 */

	// main loop
	for {
		time.Sleep(1 * time.Second)
	}

	//return errors.New("device is not running")
}
