package v1handlers

import (
	"flag"
)

// MountedVolume is the volume mounted to this program
var MountedVolume string

// Port is the TCP port for the service
var Port string

// DefaultPort is the default TCP port.
// It can be overwritten by the command line
const DefaultPort = "8080"

// RequestVersion tells if a version flag is provided
var RequestVersion bool

func init() {
	flag.StringVar(&MountedVolume, "volume", "", "You must specify a path of which the directory is mounted. It should be an absolute path which starts with slash (\"/\")")
	flag.StringVar(&Port, "port", DefaultPort, "The TCP port of RMDASHRF.")
	flag.BoolVar(&RequestVersion, "version", false, "Print the current version.")

	flag.Parse()
}
