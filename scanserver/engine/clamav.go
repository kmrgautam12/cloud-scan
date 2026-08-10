package engine

import (
	"bytes"

	"github.com/dutchcoders/go-clamd"
)

func ScanFileWithClamav(b []byte) {
	client := clamd.NewClamd("tcp://127.0.0.1:3310")
	abort := make(chan bool)
	reader := bytes.NewReader(b)
	result, err := client.ScanStream(reader, abort)
	for res := range result {
		res.
	}
}
