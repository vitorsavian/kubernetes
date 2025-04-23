package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "33"
	gitVersion   = "v1.33.0-k3s1"
	gitCommit    = "214100f340db3c691f6de7aaca8e3c2308b45df7"
	gitTreeState = "clean"
	buildDate = "2025-04-23T23:07:05Z"
)
