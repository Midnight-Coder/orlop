// Code generated by winch. DO NOT EDIT.
package version

import (
	"fmt"
	"runtime"
)

const (
	Name        = "orlop"
	Description = "Orlop is the base deck of the whole friggin ship"
	ReleaseName = "responsible leghorn"
	Version     = "1.14.0"
	Prerelease  = ""
)

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}
	return fmt.Sprintf("%s %s", Version, s)
}
