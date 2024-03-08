package device

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func Run() error {

	log.Info("Device is running")

	for true {
		// Do something
		fmt.Println("Device is running")
		time.Sleep(1 * time.Second)
	}

	return errors.New("Device is not running")
}
