//go:build windows

package daemon

import "k8s.io/klog/v2"

func NewDaemon(binaryName string) *Daemon {
	klog.Fatal("Windows userd not impl")
	return nil
}

func LoadAndLaunch(binaryName string) {
	klog.Fatal("Windows userd not impl")
}

func (d *Daemon) Stop() error {
	klog.Fatal("Windows userd not impl")
	return nil
}
