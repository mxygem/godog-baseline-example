package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config contains data specific to the environment
// the suite needs to interact with.
type Config struct {
	AccountID string
}

// NewConfig loads data based on environment
func NewConfig() Config {
	// Manually setting env for the example's sake
	os.Setenv("ENV", "dev")
	env := strings.ToLower(os.Getenv("ENV"))

	viper.SetConfigName(fmt.Sprintf("config_%s", env))
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := Config{}
	c.AccountID = viper.GetString("account_id")

	return c
}
