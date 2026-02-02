package config

import (
	"den-den-mushi-v2/pkg/config"
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Load(path string) *Config {
	var cfg Config

	dir := filepath.Dir(path)
	filename := filepath.Base(path)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType("json")
	v.AddConfigPath(dir)
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.BindSsl(v)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return &cfg
}

func HotReload(path string, cfg *Config) {
	v := viper.New()
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
}
