package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Usable-Security-and-Privacy-Lab/lets-auth-ca/certs"
	"github.com/Usable-Security-and-Privacy-Lab/lets-auth-ca/config"
)

func main() {
	// process command line arguments
	signRoot := flag.Bool("root", false, "Resigns the root certificate. Mutually exclusive with other operating flags.")
	configDir := flag.String("configDir", "configs", "configuration directory")
	configMode := flag.String("config", "development", "configuration mode")
	flag.Parse()

	config.Init(*configDir, *configMode)
	cfg := config.Get()
	fmt.Println(cfg.Name)

	if *signRoot {
		certs.ReSignRootCert()
		os.Exit(0)
	}
	// else continue with normal ca operations
}
