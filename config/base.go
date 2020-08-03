package config

var (
	// BuildTag is a commit hash that will be injected in release mode
	BuildTag = "Unknown"
	// BuildTime is a time, when it build, that will be injected in release mode
	BuildTime = "Unknown"
	// Debug is a specific var of runtime mode
	Debug = false
	// File is a path, specify where the config file used
	File = ""
	// GoVersion is the version of go, will be injected at runtime
	GoVersion = "Unknown"
	// MakeVersion is the version of make in release, will be injected in release mode
	MakeVersion = "Unknown"
	// Version is the version of this program, will be injected in release mode
	Version = "development"
)
