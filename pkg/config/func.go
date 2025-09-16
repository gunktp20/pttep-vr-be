package config

import (
	"os"
	"pttep-vr-api/pkg/constant/state"
	"regexp"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	vp  *viper.Viper
	App struct {
		Name     string     `mapstructure:"name" yaml:"name"`
		Version  string     `mapstructure:"version" yaml:"version"`
		State    state.Type `mapstructure:"state" yaml:"state"`
		Timezone string     `mapstructure:"timezone" yaml:"timezone"`
		Config   struct {
			Host   string `mapstructure:"host" yaml:"host"`
			Port   int    `mapstructure:"port" yaml:"port"`
			Path   string `mapstructure:"path" yaml:"path"`
			Allows struct {
				Origins  []string `mapstructure:"origins" yaml:"origins"`
				Response struct {
					Error bool `mapstructure:"error" yaml:"error"`
				} `mapstructure:"response" yaml:"response"`
			} `mapstructure:"allows" yaml:"allows"`
		} `mapstructure:"config" yaml:"config"`
	} `mapstructure:"app" yaml:"app"`
	Database struct {
		Host     string `mapstructure:"host" yaml:"host"`
		Port     int    `mapstructure:"port" yaml:"port"`
		Name     string `mapstructure:"name" yaml:"name"`
		Username string `mapstructure:"username" yaml:"username"`
		Password string `mapstructure:"password" yaml:"password"`
	} `mapstructure:"database" yaml:"database"`
}

var conf *Config

func Init(path string) error {

	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	for _, key := range viper.AllKeys() {
		if key == "app.config.allows.origins" {
			continue
		}
		value := replaceEnvInConfig([]byte(viper.GetString(key)))
		viper.Set(key, string(value))
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port != "" {
		p, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		conf.App.Config.Port = p
	}

	return nil
}

func Get() *Config {
	return conf
}

func replaceEnvInConfig(body []byte) []byte {
	search := regexp.MustCompile(`\$\{([^{}]+)\}`)
	replacedBody := search.ReplaceAllFunc(body, func(b []byte) []byte {
		group := search.ReplaceAllString(string(b), `$1`)
		envValue := os.Getenv(group)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		return []byte{}
	})

	return replacedBody
}
