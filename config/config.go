package config

import (
	"errors"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type Config struct {
	APIKey string
	TravelBaseURL string
	TravelOrigin string
	TravelTimeout time.Duration
}

func Load() (*Config, error) {
	c := &Config{
		APIKey: web.AppConfig.DefaultString("apikey", ""),
		TravelBaseURL: strings.TrimRight(web.AppConfig.DefaultString("travelapi_baseurl", ""), "/"),
		TravelOrigin: web.AppConfig.DefaultString("travelapi_origin", ""),
		TravelTimeout: time.Duration(web.AppConfig.DefaultInt("travelapi_timeout_seconds", 5)) * time.Second,
	}

	if c.APIKey == "" {
		return nil, errors.New("config: APP_APIKEY is required")
	}
	if c.TravelBaseURL == "" || c.TravelOrigin == "" {
		return nil, errors.New("config: travelapi_baseurl and travelapi_origin are required")
	}
	return c, nil
}