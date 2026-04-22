package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.0-k3s1"
	gitCommit    = "48526dacb8698953b928563213eb2d6d88f42ae3"
	gitTreeState = "clean"
	buildDate = "2026-04-22T20:53:17Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)
