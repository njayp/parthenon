package daemon

// Daemon defines a service
type Daemon struct {
	Label     string
	Program   string
	KeepAlive bool
	RunAtLoad bool
	plistPath string
}
