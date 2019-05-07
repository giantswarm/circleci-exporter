package collector

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/jszwedko/go-circleci"
	"github.com/prometheus/client_golang/prometheus"
)

// BuildConfig is this collector's configuration struct.
type BuildConfig struct {
	CircleCIClient *circleci.Client
	Logger         micrologger.Logger
}

// Build is the main struct for this collector.
type Build struct {
	circleCIClient *circleci.Client
	logger         micrologger.Logger
}

// NewAppResource creates a new Build metrics collector
func NewBuild(config BuildConfig) (*Build, error) {
	if config.CircleCIClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.CircleCIClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	c := &Build{
		circleCIClient: config.CircleCIClient,
		logger:         config.Logger,
	}

	return c, nil
}

// Collect is the main metrics collection function.
func (c *Build) Collect(ch chan<- prometheus.Metric) error {
	// TODO
	return nil
}

// Describe emits the description for the metrics collected here.
func (c *Build) Describe(ch chan<- *prometheus.Desc) error {
	// TODO
	return nil
}
