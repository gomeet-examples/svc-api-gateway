// The main package for the svc-echo executable.
package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/gomeet-examples/svc-echo/cmd"
)

// Main manages command execution
func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
