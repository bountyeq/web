package config

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jbsmith7741/toml"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
)

// DB represents a gorm DB
var DB *gorm.DB

// ContentDB represents a content gorm DB
var ContentDB *gorm.DB

// Config wraps all configuration options
type Config struct {
	Debug           bool     `toml:"debug" desc:"Enable debug mode"`
	Database        Database `toml:"database" desc:"database"`
	ContentDatabase Database `toml:"content_database" desc:"database"`
}

// Database contains connection info
type Database struct {
	Username string `toml:"username" desc:"username, e.g. eqemu"`
	Password string `toml:"password" desc:"password, e.g. eqemupass"`
	Host     string `toml:"host" desc:"host:port string, e.g. localhost:3306"`
	Database string `toml:"database" desc:"database name, e.g. eqemu"`
}

// New loads the config system
func New() (*Config, error) {
	var f *os.File
	cfg := Config{}
	path := "web.conf"

	isNewConfig := false
	fi, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("config info: %w", err)
		}
		f, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("create web.conf: %w", err)
		}
		fi, err = os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("new config info: %w", err)
		}
		isNewConfig = true
	}
	if !isNewConfig {
		f, err = os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open config: %w", err)
		}
	}

	defer f.Close()
	if fi.IsDir() {
		return nil, fmt.Errorf("web.conf is a directory, should be a file")
	}

	if isNewConfig {
		enc := toml.NewEncoder(f)
		enc.Encode(getDefaultConfig())

		fmt.Println("a new web.conf file was created. Please open this file and configure web, then run it again.")
		if runtime.GOOS == "windows" {
			option := ""
			fmt.Println("press a key then enter to exit.")
			fmt.Scan(&option)
		}
		os.Exit(0)
	}

	_, err = toml.DecodeReader(f, &cfg)
	if err != nil {
		return nil, fmt.Errorf("decode web.conf: %w", err)
	}
	fw, err := os.Create("web.conf")
	if err != nil {
		return nil, fmt.Errorf("web: %w", err)
	}
	defer fw.Close()

	enc := toml.NewEncoder(fw)
	err = enc.Encode(cfg)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	err = cfg.Verify()
	if err != nil {
		return nil, fmt.Errorf("verify: %w", err)
	}

	return &cfg, nil
}

func getDefaultConfig() *Config {
	cfg := Config{}
	cfg.Database.Database = "eqemu"
	cfg.Database.Username = "eqemu"
	cfg.Database.Password = "eqemupass"
	cfg.Database.Host = "127.0.0.1:3306"
	cfg.ContentDatabase.Database = "eqemu"
	cfg.ContentDatabase.Username = "eqemu"
	cfg.ContentDatabase.Password = "eqemupass"
	cfg.ContentDatabase.Host = "127.0.0.1:3306"
	return &cfg
}

// Verify config options
func (c *Config) Verify() error {
	return nil
}
