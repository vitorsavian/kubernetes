package version

var (
	gitMajor = "1"
	gitMinor = "37"
	gitVersion   = "v1.37.0-rc.1-vs1"
	gitCommit    = "df9edbbbf40e410762b3c3fe4488b5165146457f"
	gitTreeState = "clean"
	buildDate = "2026-08-22T18:49:43Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.37"
)
