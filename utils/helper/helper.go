package helper

import (
	"os"
)

func GetStringFromEnv(t string) string {
	value, _ := os.LookupEnv(t)
	return value
}
