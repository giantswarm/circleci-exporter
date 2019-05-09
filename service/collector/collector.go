package collector

const (
	namespace = "circleci_exporter"
	subsystem = "ci"
)

const (
	labelBranch = "branch"
	labelRepo   = "repo"
	labelState  = "state"
)

const (
	numBuilds    = 100
	bucketStart  = 15000
	bucketFactor = 2
	numBuckets   = 8
)
