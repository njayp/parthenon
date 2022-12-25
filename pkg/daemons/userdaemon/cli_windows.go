//go:build windows

package userdaemon

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func LoadAndLaunch(binaryName string) {
	log.Fatal("Windows userd not impl")
}
