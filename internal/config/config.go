package config

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type (
	Server struct {
		Addr string
		JWT  JWT
	}

	JWT struct {
		KeyPath   string
		PublicKey rsa.PublicKey
	}

	Database struct {
		DBConn string
	}

	Cache struct {
		Conn string
		Exp  time.Duration
	}

	Config struct {
		Srv   Server
		DB    Database
		Cache Cache
	}
)

func New() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		Srv: Server{
			Addr: viper.GetString("SERVER_ADDRESS"),
			JWT: JWT{
				KeyPath: viper.GetString("JWT_KEY_PATH"),
			},
		},
		DB: Database{
			DBConn: viper.GetString("POSTGRES_CONN"),
		},
		Cache: Cache{
			Conn: viper.GetString("REDIS_CONN"),
			Exp:  viper.GetDuration("CACHE_EXPIRATION"),
		},
	}

	publicKey, err := cfg.GetPublicKey(cfg.Srv.JWT.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get publicKey: %s", err)
	}
	cfg.Srv.JWT.PublicKey = *publicKey

	return cfg, nil
}

func (cfg *Config) GetPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	key, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read public.key file: %v", err)
	}

	return jwt.ParseRSAPublicKeyFromPEM(key)
}
