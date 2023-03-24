package version

import (
	"flag"
	"fmt"
	version "github.com/hashicorp/go-version"
	"io"
	"os"
)

var (
	Version       = "0.5.0"
	PreRelease    = ""
	BuildMetadata = ""
)

var (
	SemVer      *version.Version
	VersionFlag bool
)

func init() {
	ver := Version

	if PreRelease != "" {
		ver = fmt.Sprintf("%s-%s", Version, PreRelease)
	}

	if BuildMetadata != "" {
		ver = fmt.Sprintf("%s+%s", ver, BuildMetadata)
	}

	SemVer = version.Must(version.NewVersion(ver))

	flag.BoolVar(&VersionFlag, "version", false, "Print out Version")
	flag.BoolVar(&VersionFlag, "v", false, "Print out Version")

}

func String() string {
	return SemVer.String()
}

func PrintVersion(w io.Writer) {
	fmt.Fprintln(w, String())
	os.Exit(0)
}
