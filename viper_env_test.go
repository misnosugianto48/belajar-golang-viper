package belajargolangviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestENVFile(t *testing.T) {
	var config *viper.Viper = viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Misno Sugianto", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DB_PORT"))
	assert.True(t, config.GetBool("DB_SHOW_SQL"))
}

func TestENVEnvironment(t *testing.T) {
	var config *viper.Viper = viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "belajar-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "Misno Sugianto", config.GetString("APP_AUTHOR"))
	assert.Equal(t, 3306, config.GetInt("DB_PORT"))
	assert.True(t, config.GetBool("DB_SHOW_SQL"))

	assert.Equal(t, "hello", config.GetString("FROM_ENV"))
}
