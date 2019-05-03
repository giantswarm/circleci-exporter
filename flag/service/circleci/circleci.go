package circleci

import (
	"github.com/giantswarm/circleci-exporter/flag/service/circleci/auth"
)

type CircleCI struct {
	Auth auth.Auth
}
