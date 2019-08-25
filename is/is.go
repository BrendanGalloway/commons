package is

import (
	"os"
	"reflect"
)

//Slice returns true if the argument is a slice
func Slice(arg interface{}) bool {
	return reflect.ValueOf(arg).Kind() == reflect.Slice
}

func TTY() bool {
	fi, _ := os.Stdout.Stat()
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func File(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
