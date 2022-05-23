package helpers

import "github.com/labstack/gommon/log"

func Check(err error) {
	if err != nil {
		log.Error(err)
	}
}
