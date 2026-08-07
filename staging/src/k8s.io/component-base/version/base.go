package version

var (
	gitMajor = "1"
	gitMinor = "37"
	gitVersion   = "v1.37.0-rc.0-k3s1"
	gitCommit    = "1aa5250f174b60a9ea01293bfd58218966efa689"
	gitTreeState = "clean"
	buildDate = "2026-08-07T13:57:28Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.37"
)
