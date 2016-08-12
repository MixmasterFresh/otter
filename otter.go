package otter

import (
	"flag"
	"fmt"
	"os"
)

const (
	// VERSION describes the otter version being developed
	VERSION      = "0.1.0"
	helpUsage    = "prints the help message"
	versionUsage = "prints the version of otter you are currently using"
	portUsage    = "The port to expose the otter server on"
	keyUsage     = "provide a strong key here for authenticating your server"
	helpText     = "text goes here" //TODO: fix this
)

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, helpUsage)

	var port int
	flag.IntVar(&port, "port", 1729, portUsage)
	flag.IntVar(&port, "p", 1729, portUsage)

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, versionUsage)

	var key string
	flag.StringVar(&key, "key", "", keyUsage)
	flag.StringVar(&key, "k", "", keyUsage)

	flag.Parse()

	if help {
		fmt.Printf("%s\n", helpText)
		os.Exit(0)
	}
	if versionFlag {
		fmt.Printf("%s\n", VERSION)
		os.Exit(0)
	}

	if port < 0 || port > 65535 {
		panic("Port must be a valid number between 1 and 65535 and must be free\n")
	}

	server(port, key)
}
