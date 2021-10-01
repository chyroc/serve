package internal

import (
	"os"
)

var (
	home        string
	workDir     string
	workDirCert string
)

func init() {
	home, _ = os.UserHomeDir()
	workDir = home + "/.chyroc-serve"
	workDirCert = workDir + "/cert"
	err := os.MkdirAll(workDirCert, 0o777)
	if err != nil {
		panic(err)
	}
}
