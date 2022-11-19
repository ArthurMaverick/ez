package util

import (
	"io/ioutil"
	"log"
)

func CopyTemplate(src string, dest string) (nil error) {
	byteRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = ioutil.WriteFile(dest, byteRead, 0755)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
