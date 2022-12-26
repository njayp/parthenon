//go:build darwin

package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"k8s.io/klog/v2"
)

const pListTemplate = `
<?xml version='1.0' encoding='UTF-8'?>
 <!DOCTYPE plist PUBLIC \"-//Apple Computer//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\" >
 <plist version='1.0'>
   <dict>
     <key>Label</key><string>{{.Label}}</string>
     <key>Program</key><string>{{.Program}}</string>
     <key>StandardOutPath</key><string>/tmp/{{.Label}}.out.log</string>
     <key>StandardErrorPath</key><string>/tmp/{{.Label}}.err.log</string>
     <key>KeepAlive</key><{{.KeepAlive}}/>
     <key>RunAtLoad</key><{{.RunAtLoad}}/>
   </dict>
</plist>
`

// Daemon defines a service
type Daemon struct {
	Label     string
	Program   string
	KeepAlive bool
	RunAtLoad bool
	plistPath string
}

// NewDaemon constructs a launchd service.
func NewDaemon(binaryName string) Daemon {
	home := os.Getenv("HOME")
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(home, "go")
	}
	label := fmt.Sprintf("com.njayp.%s", binaryName)
	return Daemon{
		Label:     label, // Reverse-DNS naming convention
		Program:   filepath.Join(gopath, "bin", binaryName),
		KeepAlive: true,
		RunAtLoad: true,
		plistPath: filepath.Join(home, fmt.Sprintf("Library/LaunchAgents/%s.plist", label)),
	}
}

func (d Daemon) Install() error {
	f, err := os.OpenFile(d.plistPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("error opening path %s: %s", d.plistPath, err)
	}
	t := template.Must(template.New("launchdConfig").Parse(pListTemplate))
	err = t.Execute(f, d)
	if err != nil {
		return fmt.Errorf("error writing file: %s", err)
	}
	return nil
}

// Start will start the service.
func (d Daemon) Start() error {
	klog.Infof("Starting %s", d.Label)
	// We start using load -w on plist file
	output, err := exec.Command("/bin/launchctl", "load", "-w", d.plistPath).CombinedOutput()
	klog.Infof("Output (launchctl load): %s", string(output))
	return err
}

func (d Daemon) Stop() error {
	// We stop by removing the job. This works for non-demand and demand jobs.
	output, err := exec.Command("/bin/launchctl", "unload", d.plistPath).CombinedOutput()
	klog.Infof("Output (launchctl remove): %s", string(output))
	return err
}

func BuildTagTest() {
	klog.Info("darwin")
}
