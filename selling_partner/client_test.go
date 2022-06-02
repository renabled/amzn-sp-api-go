package selling_partner_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xamandar/amzn-sp-api-go/selling_partner"
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
