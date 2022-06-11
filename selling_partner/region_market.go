package selling_partner

import (
	"strings"
)

type Region struct {
	Name         string
	Endpoint     string
	Region       string
	Marketplaces map[string]Marketplace
}

type Marketplace struct {
	Country, MarketplaceID, CountryCode string
}

// Compiled from:
// https://developer-docs.amazon.com/sp-api/docs/sp-api-endpoints
// https://developer-docs.amazon.com/sp-api/docs/marketplace-ids
var regions = map[string]Region{
	"NA": {
		Name:     "North America",
		Endpoint: "sellingpartnerapi-na.amazon.com",
		Region:   "us-west-1",
		Marketplaces: map[string]Marketplace{
			"CA": {"Canada", "A2EUQ1WTGCTBG2", "CA"},
			"US": {"United States of America", "ATVPDKIKX0DER", "US"},
			"MX": {"Mexico", "A1AM78C64UM0Y8", "MX"},
			"BR": {"Brazil", "A2Q3Y263D00KWC", "BR"},
		},
	},
	"EU": {
		Name:     "Europe",
		Endpoint: "sellingpartnerapi-eu.amazon.com",
		Region:   "eu-west-1",
		Marketplaces: map[string]Marketplace{
			"ES": {"Spain", "A1RKKUPIHCS9HS", "ES"},
			"UK": {"United Kingdom", "A1F83G8C2ARO7P", "UK"},
			"FR": {"France", "A13V1IB3VIYZZH", "FR"},
			"BE": {"Belgium", "AMEN7PMS3EDWL", "BE"},
			"NL": {"Netherlands", "A1805IZSGTT6HS", "NL"},
			"DE": {"Germany", "A1PA6795UKMFR9", "DE"},
			"IT": {"Italy", "APJ6JRA9NG5V4", "IT"},
			"SE": {"Sweden", "A2NODRKZP88ZB9", "SE"},
			"PL": {"Poland", "A1C3SOZRARQ6R3", "PL"},
			"EG": {"Egypt", "ARBP9OOSHTCHU", "EG"},
			"TR": {"Turkey", "A33AVAJ2PDY3EV", "TR"},
			"SA": {"Saudi Arabia", "A17E79C6D8DWNP", "SA"},
			"AE": {"United Arab Emirates", "A2VIGQ35RCS4UG", "AE"},
			"IN": {"India", "A21TJRUUN4KGV", "IN"},
		},
	},
	"FE": {
		Name:     "Far East",
		Endpoint: "sellingpartnerapi-fe.amazon.com",
		Region:   "us-west-2",
		Marketplaces: map[string]Marketplace{
			"SG": {"Singapore", "A19VAU5U5O7RUS", "SG"},
			"AU": {"Australia", "A39IBJ37TRP1C6", "AU"},
			"JP": {"Japan", "A1VC38T7YXB528", "JP"},
		},
	},
}

func MarketplaceID(m string) string {
	return CountryCodeToMarketplace(m).MarketplaceID
}

func CountryCodeToMarketplace(m string) Marketplace {
	for _, v := range regions {
		if mp, ok := v.Marketplaces[strings.ToUpper(m)]; ok {
			return mp
		}
	}
	return Marketplace{}
}
