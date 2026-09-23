package version

var (
	gitMajor = "1"
	gitMinor = "37"
	gitVersion   = "v1.37.1-k3s1"
	gitCommit    = "59321a1aa8a3decc1e46590c954e1920495da126"
	gitTreeState = "clean"
	buildDate = "2026-09-23T20:27:17Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.37"
)
