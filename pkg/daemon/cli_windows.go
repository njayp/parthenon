//go:build windows

package daemon

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func LoadAndLaunch(binaryName string) {
	klog.Fatal("Windows userd not impl")
}
