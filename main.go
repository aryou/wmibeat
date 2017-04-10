package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/singlehopllc/wmibeat/beater"
)

func main() {
	err := beat.Run("wmibeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
