package initializers

import "github.com/spf13/viper"

type Config struct {
	DBHost 		   string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
    // Set the configuration path
    viper.AddConfigPath(path)
    viper.SetConfigType("env")
    viper.SetConfigName("app")

    // Automatically bind environment variables with the Config struct
    viper.AutomaticEnv()

    // Read the configuration file
    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    // Unmarshal the configuration data into the Config struct
    err = viper.Unmarshal(&config)
    return
}
