package config

import (
	"time"

	"go.uber.org/zap/zapcore"
)

const (
	UUIDLabel          = "buildkite.com/job-uuid"
	BuildURLAnnotation = "buildkite.com/build-url"
	JobURLAnnotation   = "buildkite.com/job-url"
	DefaultNamespace   = "default"
	DefaultAgentImage  = "ghcr.io/buildkite/agent-k8s:latest"
)

type Config struct {
	AgentTokenSecret string `mapstructure:"agent-token-secret" validate:"required"`
	BuildkiteToken   string `mapstructure:"buildkite-token" validate:"required"`
	Debug            bool
	Image            string        `validate:"required"`
	JobTTL           time.Duration `mapstructure:"job-ttl"`
	MaxInFlight      int           `mapstructure:"max-in-flight" validate:"min=0"`
	Namespace        string        `validate:"required"`
	Org              string        `validate:"required"`
	Tags             stringSlice   `validate:"min=1"`
	ProfilerAddress  string        `mapstructure:"profiler-address" validate:"omitempty,hostname_port"`
	ClusterUUID      string        `mapstructure:"cluster-uuid" validate:"omitempty"`
}

type stringSlice []string

func (s stringSlice) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, x := range s {
		enc.AppendString(x)
	}
	return nil
}

func (c Config) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("agent-token-secret", c.AgentTokenSecret)
	enc.AddBool("debug", c.Debug)
	enc.AddString("image", c.Image)
	enc.AddDuration("job-ttl", c.JobTTL)
	enc.AddInt("max-in-flight", c.MaxInFlight)
	enc.AddString("namespace", c.Namespace)
	enc.AddString("org", c.Org)
	enc.AddString("profiler-address", c.ProfilerAddress)
	enc.AddString("cluster-uuid", c.ClusterUUID)
	return enc.AddArray("tags", c.Tags)
}
