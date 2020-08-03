package flag

import (
	pflag "github.com/spf13/pflag"
)

// Parse params
func Parse() {
	// Register Flag mapping
	registerVersionFlag()
	registerHelpFlag()
	registerDebugFlag()
	registerConfigFlag()

	// Parse Flag
	pflag.Parse()

	// Handle Flag event
	handleVersionFlag()
	handleHelpFlag()
}
