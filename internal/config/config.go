package config

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Log     Log     `json:"log" toml:"log"`
	MongoDB MongoDB `json:"mongodb" toml:"mongodb"`
	Etcd    Etcd    `json:"etcd" toml:"etcd"`
	Server  Server  `json:"server" toml:"server"`
	Collect Collect `json:"collect" toml:"collect"`
}

type Log struct {
	DisableTimestamp bool   `json:"disable-timestamp" toml:"disable-timestamp"`
	Level            string `json:"level" toml:"level"`
	Format           string `json:"format" toml:"format"`
	FileName         string `json:"filename" toml:"filename"`
	MaxSize          int    `json:"maxsize" toml:"maxsize"`
}

type MongoDB struct {
	DSN string `json:"dsn" toml:"dsn"`
}

type Etcd struct {
	Endpoints []string `json:"endpoints" toml:"endpoints"`
}

type Server struct {
	Host string `json:"host" toml:"host"`
	Port int    `json:"port" toml:"port"`
}

type Collect struct {
	CodeList []string `json:"code-list" toml:"code-list"`
}

func (c *Config) Load(path string, override func(cfg *Config)) error {
	if path == "" {
		return nil
	}

	if _, err := toml.DecodeFile(path, c); err != nil {
		return err
	}
	return nil
}

func (cg *Config) String() string {
	buf, _ := json.Marshal(cg)
	return string(buf)
}

var GlobalConfig = &Config{
	Log: Log{
		DisableTimestamp: false,
		Level:            "info",
		Format:           "text",
		FileName:         "/tmp/robber-datasource/data.log",
		MaxSize:          20,
	},
	MongoDB: MongoDB{
		DSN: "mongodb://localhost:27017",
	},
	Server: Server{
		Host: "0.0.0.0",
		Port: 19090,
	},
	Collect: Collect{
		CodeList: []string{
			"sh688***",
			"sh605***",
			"sh603***",
			"sh601***",
			"sh600***",
			"sz300***",
			"sz0030**",
			"sz002***",
			"sz001**",
			"sz000***",
		},
	},
}
