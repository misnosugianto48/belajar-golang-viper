package belajargolangviper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestJson(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	// read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	// mengambil value
	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, "Misno Sugianto", config.GetString("app.author"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.True(t, config.GetBool("database.show_sql"))
}
