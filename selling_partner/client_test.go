package selling_partner_test

import (
	"testing"

	"github.com/renabled/amzn-sp-api-go/selling_partner"
	"github.com/stretchr/testify/require"
)

func Test_Init(t *testing.T) {
	var cli *selling_partner.Client
	var config *selling_partner.Config

	config = selling_partner.NewConfig("", "", "", "", "", "", "na")
	cli = selling_partner.New(config)
	require.NotNil(t, cli)

	config = selling_partner.NewConfig("", "", "", "", "", "", "eu")
	cli = selling_partner.New(config)
	require.NotNil(t, cli)

	config = selling_partner.NewConfig("", "", "", "", "", "", "fe")
	cli = selling_partner.New(config)
	require.NotNil(t, cli)
}

func Test_MarketplaceToID(t *testing.T) {
	require.Equal(t, "A1RKKUPIHCS9HS", selling_partner.MarketplaceID("ES"))
	require.Equal(t, "A2NODRKZP88ZB9", selling_partner.MarketplaceID("SE"))
	require.Equal(t, "", selling_partner.MarketplaceID("IL"))
}

func Test_CountryCodeToMarketplace(t *testing.T) {
	require.Equal(t, "Spain", selling_partner.CountryCodeToMarketplace("ES").Country)
	require.Equal(t, "Sweden", selling_partner.CountryCodeToMarketplace("SE").Country)
	require.Empty(t, selling_partner.CountryCodeToMarketplace("IL"))
}
