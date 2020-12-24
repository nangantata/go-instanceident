package instanceident

import (
	"os"
	"strconv"
)

// ProcessDerivedStringIdentity generate identity string from host name and PID of running process.
func ProcessDerivedStringIdentity() (identity string, err error) {
	hostname, err := os.Hostname()
	if nil != err {
		return
	}
	pidVal := os.Getpid()
	pidText := strconv.FormatInt(int64(pidVal), 10)
	identity = hostname + "#" + pidText
	return identity, nil
}
