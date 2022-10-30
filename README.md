# Amazon Selling Partner API in Go

[Selling Partner API or SP-API](https://developer-docs.amazon.com/sp-api/) is the next generation suite of API-based automation functionality for Amazon's Selling Partners and is an evolution of (the legacy) [Amazon Marketplace Web Service (Amazon MWS) APIs](https://docs.developer.amazonservices.com/en_US/dev_guide/index.html]), which have been offering sellers programmatic access to critical Amazon features for more than 10 years. SP-API is available in production in all the marketplaces where Amazon MWS is available.

This package provides an easy way to communicate with SP-API.

## How does this work

Amazon provides [swagger](https://swagger.io/) models for each API, and using that you will be able to form requests to the API. Some more steps are needed to have a successful call to the API, including authentication, signing the request, etc.

This package will make it easier to make authenticated requests to SP-API. The HTTP client(s) are generated from the swagger models specs using [go-swagger](https://goswagger.io/).

To keep the client(s) up to date, there is a schedule job that runs to fetch the Selling Partner API [models](https://github.com/amzn/selling-partner-api-models/tree/main/models) from Amazon's Github account, regenerate the client(s), and commit the changes to a new branch called `next`.

When we think that `next` branch is stable enough to be used, we will merge it into the `main` branch and create a new version out of the new or updated models.

## Installation

```
go get github.com/xamandar/amzn-sp-api-go
```

## Example

```golang
import (
  "fmt"

  "github.com/xamandar/amzn-sp-api-go/api/sellers/sellers_client/sellers"
  "github.com/xamandar/amzn-sp-api-go/selling_partner"
)

// Initilize your config
config := selling_partner.NewConfig(
  "_accessKeyId_",
  "_secretKey_",
  "_roleArn_",
  "_refreshToken_",
  "_clientID_",
  "_clientSecret_",
  "region", // NA, EU or FE
)

// Create a selling partner client providing the config
cli := selling_partner.New(config)

// Transport that require auth
transport := cli.ClientTransport(
  context.Background(),
  /* selling_partner.Grantless(), selling_partner.Debug() */)

// Call sellers
sellersCli := sellers.New(transport, nil)

// Make the actual request
res, err := sellersCli.GetMarketplaceParticipations(sellers.NewGetMarketplaceParticipationsParams())

fmt.Println(res, err)
```
