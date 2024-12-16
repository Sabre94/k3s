package version

const (
	// DefaultKubeBinaryVersion is the hard coded k8 binary version based on the latest K8s release.
	// It is supposed to be consistent with gitMajor and gitMinor, except for local tests, where gitMajor and gitMinor are "".
	// Should update for each minor release!
	DefaultKubeBinaryVersion = "1.31"
)

var (
	gitMajor = "1"
	gitMinor = "31"
	gitVersion   = "v1.31.4-k3s1"
	gitCommit    = "275bb129204e71758cec0cebd7a01f791c1f4bc3"
	gitTreeState = "clean"
	buildDate = "2024-12-11T13:43:23Z"
)
