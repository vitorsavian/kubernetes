package version

var (
	gitMajor = "1"
	gitMinor = "37"
	gitVersion   = "v1.37.0-k3s1"
	gitCommit    = "02ebfdb895be969cc5a92f282f0c8b60a5f87d52"
	gitTreeState = "clean"
	buildDate = "2026-08-26T18:09:40Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.37"
)
