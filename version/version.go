package version

import (
	"fmt"
	"io"
	"os"
)

var (
	// Version is Clamp's version.
	Version = "unknown version"
	// Commit is the hash of the commit used to build Clamp.
	Commit = "unknown commit"
	// Date is when Clamp was built.
	Date = "unknown build date"
)

// Print writes Clamp's version information to standard output.
func Print() {
	Fprint(os.Stdout)
}

// Fprint writes Clamp's version information to w.
func Fprint(w io.Writer) {
	fmt.Fprint(w, "Clamp version %s build %s (at %s)\n", Version, Commit, Date)
}
