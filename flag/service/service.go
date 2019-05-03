package service

import (
	"github.com/giantswarm/circleci-exporter/flag/service/circleci"
)

type Service struct {
	CircleCI circleci.CircleCI
}
