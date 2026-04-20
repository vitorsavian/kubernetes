package version

var (
	gitMajor = "1"
	gitMinor = "36"
	gitVersion   = "v1.36.0-rc.1-k3s1"
	gitCommit    = "5164e240b5af76b89582b5281b897a10566852cf"
	gitTreeState = "clean"
	buildDate = "2026-04-20T17:32:01Z"
)

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.36"
)
