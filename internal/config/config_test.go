package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCfg(t *testing.T) {
	_assert := assert.New(t)
	var (
		path = "../../tests/config/config-test.toml"
		cfg  = &Config{}
	)

	err := cfg.Load(path, nil)
	_assert.Nil(err)

	_assert.Equal(true, cfg.Log.DisableTimestamp)
	_assert.Equal("test", cfg.Log.Level)
	_assert.Equal("test", cfg.Log.Format)
	_assert.Equal("/tmp/robber-datasource/data.log", cfg.Log.FileName)
	_assert.Equal(2, cfg.Log.MaxSize)
	_assert.Equal("mongodb://127.0.0.1:27017", cfg.MongoDB.DSN)
	_assert.Equal("0.0.0.1", cfg.Server.Host)
	_assert.Equal(29090, cfg.Server.Port)
	_assert.Equal(3, len(cfg.Collect.CodeList))
}
