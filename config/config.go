package config

import "github.com/cristalhq/aconfig"

type Config struct {
	Workdir string `default:"workdir" env:"DB" flag:"db" usage:"bbolt working directory path"`
	DB      string `default:"lapunder.bboltdb" env:"DB" flag:"db" usage:"bbolt db file name"`
}

func Load() (Config, error) {
	cfg := Config{}
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		EnvPrefix: "LAPUNDER",
	})

	if err := loader.Load(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
