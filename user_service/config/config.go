package config

type Config struct {
	AppPort            string `envconfig:"APP_PORT" default:"5000"`
	HealthPort         string `envconfig:"HEALTH_PORT" default:"50000"`
	GracefulPeriod     string `envconfig:"GRACEFUL_PERIOD" default:"10"`
	GoogleCloudProject string `envconfig:"GOOGLE_CLOUD_PROJECT" default:""`
}

func (c *Config) GetAppPort() string {
	if c != nil {
		return c.AppPort
	}
	return ""
}

func (c *Config) GetHealthPort() string {
	if c != nil {
		return c.HealthPort
	}
	return ""
}
func (c *Config) GetGracefulPeriod() string {
	if c != nil {
		return c.GracefulPeriod
	}
	return ""
}

func (c *Config) GetGoogleCloudProject() string {
	if c != nil {
		return c.GoogleCloudProject
	}
	return ""
}
