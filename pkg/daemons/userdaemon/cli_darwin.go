//go:build darwin

package userdaemon

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func LoadAndLaunch(binaryName string) {
	data := struct {
		Label     string
		Program   string
		KeepAlive bool
		RunAtLoad bool
	}{
		Label:     fmt.Sprintf("com.njayp.%s", binaryName), // Reverse-DNS naming convention
		Program:   fmt.Sprintf("%s/bin/%s", os.Getenv("GOPATH"), binaryName),
		KeepAlive: true,
		RunAtLoad: true,
	}

	plistPath := fmt.Sprintf("%s/Library/LaunchAgents/%s.plist", os.Getenv("HOME"), data.Label)
	f, err := os.Open(plistPath)
	if err != nil {
		log.Fatalf("Opening pList path failed: %s", err)
	}
	t := template.Must(template.New("launchdConfig").Parse(pListTemplate))
	err = t.Execute(f, data)
	if err != nil {
		log.Fatalf("Template generation failed: %s", err)
	}
}
