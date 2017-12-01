package main

// Deliberately uninitialized. See below.
var buildVersion string

// versionInfo returns a string containing the version information of the
// current build. It's empty by default, but can be included as part of the
// build process by setting the main.buildVersion variable.
func versionInfo() string {
	if buildVersion == "" {
		return "unknown"
	}
	return buildVersion
}
