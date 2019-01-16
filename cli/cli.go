package cli

import (
	"flag"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/eduardo-matos/goversion/vchecker"
)

func parseVersion() (string, error) {
	var version string
	flag.StringVar(&version, "version", "", "Go version that will be checked")
	flag.Parse()

	if version == "" {
		return "", errors.Errorf("Version not provided: %s", version)
	}

	return version, nil
}

// Run runs version checker
func Run() {
	version, err := parseVersion()

	if err != nil {
		log.Fatal(err)
		return
	}

	exists, err := vchecker.Version(version, nil)
	vchecker.Output(os.Stdout, exists, err)
}
