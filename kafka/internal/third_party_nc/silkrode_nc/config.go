package silkrode_nc

type Config struct {
	AppID     string `yaml:"app_id" mapstructure:"app_id"`
	SecretKey string `yaml:"secret_key" mapstructure:"secret_key"`
	Hostname  string
}
