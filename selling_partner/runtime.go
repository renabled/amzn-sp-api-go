package selling_partner

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"golang.org/x/oauth2/clientcredentials"
)

func (sp *Client) ClientTransport(ctx context.Context, isGrantless bool) *client.Runtime {
	aws4Signer, err := sp.signer()
	if err != nil {
		return nil
	}

	src := sp.authTokenSource(ctx)
	if isGrantless {
		src = sp.grantlessTokenSource(ctx)
	}

	ctx = context.WithValue(ctx, oauth2.HTTPClient,
		&http.Client{
			Transport: &oauth2.Transport{
				Source: src,
				Base: &transport{
					aws4Signer: aws4Signer,
					region:     sp.config.region,
				},
			},
		},
	)
	cli := oauth2.NewClient(ctx, src)

	return client.NewWithClient(sp.config.endpoint, "/", []string{"https"}, cli)
}

func (sp *Client) grantlessTokenSource(ctx context.Context) oauth2.TokenSource {
	conf := &clientcredentials.Config{
		ClientID:     sp.config.clientID,
		ClientSecret: sp.config.clientSecret,
		TokenURL:     amazon.Endpoint.TokenURL,
		Scopes:       sp.config.withScopes,
	}
	return conf.TokenSource(ctx)
}

func (sp *Client) authTokenSource(ctx context.Context) oauth2.TokenSource {
	conf := &oauth2.Config{
		ClientID:     sp.config.clientID,
		ClientSecret: sp.config.clientSecret,
		Endpoint:     amazon.Endpoint,
	}
	return conf.TokenSource(ctx, &oauth2.Token{RefreshToken: sp.config.refreshToken})
}
