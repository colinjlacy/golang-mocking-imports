package create

import (
	"errors"
	"fmt"
	"os"
)

var (
	osStat   = os.Stat
	osCreate = os.Create
)

func CreateFile(filename, contents string) error {
	_, err := osStat(filename)
	if err == nil {
		return fmt.Errorf("file already exists")
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file system error: %s", err)
	}
	f, err := osCreate(filename)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("could not create file: %s", err)
	}
	_, err = f.Write([]byte(contents))
	return err
}
