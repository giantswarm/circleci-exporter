package collector

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/jszwedko/go-circleci"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	buildLabelDesc *prometheus.Desc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "build"),
		"circle CI build labels.",
		[]string{
			labelBranch,
			labelRepo,
			labelState,
		},
		nil,
	)
)

var (
	buildDesc = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    prometheus.BuildFQName(namespace, subsystem, "labels_buildtime"),
			Help:    "circle CI build time duration.",
			Buckets: prometheus.ExponentialBuckets(bucketStart, bucketFactor, numBuckets),
		},
		[]string{
			labelBranch,
			labelRepo,
			labelState,
		},
	)
)

func init() {
	prometheus.MustRegister(buildDesc)
}

type key struct {
	Branch string
	Repo   string
	Status string
}

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
	builds, err := c.circleCIClient.ListRecentBuilds(100, 0)
	if err != nil {
		return microerror.Mask(err)
	}

	repoCount := map[key]int{}
	for _, build := range builds {
		branch := build.Branch
		repo := build.Reponame
		status := build.Status

		buildDesc.WithLabelValues(branch, repo, status).Observe(float64(*build.BuildTimeMillis))

		{
			k := key{
				Branch: build.Branch,
				Repo:   build.Reponame,
				Status: build.Status,
			}
			repoCount[k] = repoCount[k] + 1
		}

	}

	for k, v := range repoCount {
		ch <- prometheus.MustNewConstMetric(
			buildLabelDesc,
			prometheus.GaugeValue,
			float64(v),
			k.Branch,
			k.Repo,
			k.Status,
		)
	}

	return nil
}

// Describe emits the description for the metrics collected here.
func (c *Build) Describe(ch chan<- *prometheus.Desc) error {
	ch <- buildLabelDesc
	return nil
}
