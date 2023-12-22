package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	starry_conf     *StarryConfig
	raw_starry_conf *viper.Viper
)

type StarryConfig struct {
	Player player `mapstructure:"player"`
	Meta   meta   `mapstructure:"meta"`
	Login  login  `mapstructure:"login"`
}
type player struct {
	Volume   int    `mapstructure:"volume"`
	LastSong string `mapstructure:"last_song"`
}
type login struct {
	Cookies string `mapstructure:"cookies"`
}
type meta struct {
	IsLoggedIn bool   `mapstructure:"is_logged_in"`
	JSONLogs   bool   `mapstructure:"json_logs"`
	LogLevel   string `mapstructure:"log_level"`
}

func GetConfig() *StarryConfig {
	return starry_conf
}

func SetRawConfig(key string, value any) {
	raw_starry_conf.Set(key, value)
}

func GetRawConfig() *viper.Viper {
	return raw_starry_conf
}

func SaveRawConfig() error {
	os.Remove(raw_starry_conf.ConfigFileUsed())

	err := raw_starry_conf.SafeWriteConfig()
	if err != nil {
		return err
	}
	Initialize()
	return nil
}

func validateIntRange(in *int, lower int, upper int) {
	*in = max(min(lower, *in), upper)
}

func UnmarshalToStruct() {
	starry_conf = &StarryConfig{}

	err := raw_starry_conf.Unmarshal(starry_conf)
	if err != nil {
		log.Printf("unable to decode into config struct, %v", err)
	}
	validateIntRange(&starry_conf.Player.Volume, 0, 100)
	// validateLoginHeaders(starry_conf.Login.Headers)
}

func Initialize() {
	raw_starry_conf = viper.New()
	raw_starry_conf.SetConfigName("starry_config")
	raw_starry_conf.AddConfigPath(".")
	raw_starry_conf.SetConfigType("toml")
	raw_starry_conf.SetDefault("player.volume", 75)

	err := raw_starry_conf.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	UnmarshalToStruct()
}

func init() {
	Initialize()
}
