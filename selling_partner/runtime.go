package selling_partner

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/amazon"
	"golang.org/x/oauth2/clientcredentials"
)

type option struct {
	debug       bool
	isGrantless bool
}

type Opt = func(*option)

func Debug() Opt {
	return func(o *option) {
		o.debug = true
	}
}

func Grantless() Opt {
	return func(o *option) {
		o.isGrantless = true
	}
}

func (sp *Client) ClientTransport(ctx context.Context, opts ...Opt) *client.Runtime {
	opt := option{}

	for _, m := range opts {
		m(&opt)
	}

	aws4Signer, err := sp.signer()
	if err != nil {
		return nil
	}

	if opt.debug {
		aws4Signer.Debug = *aws.LogLevel(aws.LogDebug | aws.LogDebugWithSigning)
		aws4Signer.Logger = aws.NewDefaultLogger()
	}

	src := sp.authTokenSource(ctx)
	if opt.isGrantless {
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

	cli := client.NewWithClient(sp.config.endpoint, "/", []string{"https"}, oauth2.NewClient(ctx, src))
	if opt.debug {
		cli.SetDebug(true)
	}
	return cli
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
