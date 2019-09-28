package config

type Config struct {
	GoogleCloudProject string `envconfig:"GOOGLE_CLOUD_PROJECT" default:""`
}
