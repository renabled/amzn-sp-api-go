package selling_partner

import (
	"strings"
)

type Client struct {
	config       *Config
	marketplaces map[string]Marketplace
}

func New(conf *Config) *Client {
	return &Client{
		config:       conf,
		marketplaces: regions[strings.ToUpper(conf.region)].Marketplaces,
	}
}
