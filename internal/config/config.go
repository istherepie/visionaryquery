package config

import (
	"fmt"
	"net/url"
)

// Config holds all the configurable parameters
// TODO: Should have validation to ensure that parameters are present and sane.
type Config struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Schema   string `yaml:"schema"`
	Table    string `yaml:"table"`
	Studio   string `yaml:"studio"`
	Dataset  string `yaml:"dataset"`
}

func (c *Config) ConnectionURI() string {
	query := url.Values{}
	query.Add("app name", "VisionaryQuery")
	query.Add("database", c.Database)

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}

	return u.String()
}
