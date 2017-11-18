package version

import (
	"fmt"
)

// Build information constants used, retrieved via ldflags.
var (
	appName   = "unknown"
	tag       = "unknown"
	branch    = "unknown"
	revision  = "unknown"
	goVersion = "unknown"
	buildDate = "unknown"
	buildUser = "unknown"
)

// Version is a structure with build information about Gosea.
type Version struct {
	AppName   string `json:"app_name"`
	Tag       string `json:"tag"`
	Branch    string `json:"branch"`
	Revision  string `json:"revision"`
	GoVersion string `json:"go_version"`
	BuildDate string `json:"build_date"`
	BuildUser string `json:"build_user"`
}

// New returns a new Version object that contains the current build information.
func New() Version {
	return Version{
		AppName:   appName,
		Tag:       tag,
		Branch:    branch,
		Revision:  revision,
		GoVersion: goVersion,
		BuildDate: buildDate,
		BuildUser: buildUser,
	}
}

// String returns the application name with the most recent tag.
func (v Version) String() string {
	return fmt.Sprintf("%s - tag %s\n", v.AppName, v.Tag)
}

// Print outputs detailed and formatted build information about the current version.
func (v Version) Print() {
	fmt.Printf("%s\n", v.AppName)
	fmt.Printf("  tag: \t\t%s\n", v.Tag)
	fmt.Printf("  branch: \t%s\n", v.Branch)
	fmt.Printf("  revision: \t%s\n", v.Revision)
	fmt.Printf("  build date: \t%s\n", v.BuildDate)
	fmt.Printf("  build user: \t%s\n", v.BuildUser)
	fmt.Printf("  go version: \t%s\n", v.GoVersion)
}
