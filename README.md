# Amazon Selling Partner API in Go

## Installation

```
go get github.com/xamandar/amzn-sp-api-go
```

## Example

```golang
import (
  "fmt"

  "github.com/xamandar/amzn-sp-api-go/api/catalogItems_2022-04-01/catalog_items_2022_04_01_client"
  "github.com/xamandar/amzn-sp-api-go/api/catalogItems_2022-04-01/catalog_items_2022_04_01_client/catalog"
  "github.com/xamandar/amzn-sp-api-go/selling_partner"
)

// Initilize your config
config := selling_partner.NewConfig(
  "accessKeyId",
  "secretKey",
  "roleArn",
  "refreshToken",
  "clientID",
  "clientSecret",
  "region",
)

// Create a selling partner client providing the config
cli := selling_partner.New(config)

// Transport that require auth
transport := cli.ClientTransport(context.Background(), false /* true, if grantless ops */)

// Call catalog_items_2022_04_01
catalogCli := catalog_items_2022_04_01_client.New(transport, nil)

// Make the actual request
res, err := catalogCli.Catalog.GetCatalogItem(
  &catalog.GetCatalogItemParams{
    MarketplaceIds: []string{selling_partner.Regions["EU"].Marketplaces["ES"].MarketplaceID},
    Asin:           "B01GI7NONI",
    IncludedData:   []string{"summaries"},
    Context:        context.Background(),
  },
)

fmt.Println(res, err)
```
