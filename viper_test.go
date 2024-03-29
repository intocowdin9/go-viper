package go_viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read file
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper", config.GetString("app.name"))
	assert.Equal(t, "Muhammad Rafli", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestYAML(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("json")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	// read file
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper", config.GetString("app.name"))
	assert.Equal(t, "Muhammad Rafli", config.GetString("app.author"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestENVFile(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("json")
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	// for environment-variable
	config.AutomaticEnv()

	// read file
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Muhammad Rafli", config.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))

	// from environment-variable
	// fisrt do, 'export FROM_ENV=Hello' in your terminal
	// if want to remove, do 'unset FROM_ENV'
	assert.Equal(t, "Hello", config.GetString("FROM_ENV"))

	fmt.Println(config.GetString("FROM_ENV"))
}
