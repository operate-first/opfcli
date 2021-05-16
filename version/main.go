// Package version contains variables modified at compile time to
// embed version information in the binary.
package version

// Version information used by the 'opfcli version' command. Values
// are substituted at build time using ldflags.
var (
	Name         = "opfcli"
	BuildDate    string
	BuildHash    string
	BuildVersion string
)
