package lib

import "log"

func CheckErrorAndLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
