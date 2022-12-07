package util

import "os"

func CreateDir(path string) (nil error) {
	err := os.Mkdir(path, 0755)

	if err != nil {
		return err
	}
	return nil
}
