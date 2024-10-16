package config

import "time"

type Auth struct {
	JWTSettings JWTSettings `envPrefix:"JWT_SETTINGS_"`
}

type JWTSettings struct {
	Secret   string `env:"SECRET,required"`
	Lifetime struct {
		Access  time.Duration `env:"ACCESS,required"`
		Refresh time.Duration `env:"REFRESH,required"`
	} `envPrefix:"LIFETIME_"`
}
