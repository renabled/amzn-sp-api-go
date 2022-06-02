package selling_partner

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/sts"
)

type transport struct {
	aws4Signer        *v4.Signer
	region            string
	originalTransport http.RoundTripper
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	var body io.ReadSeeker
	if r.Body != nil {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(payload))
		body = bytes.NewReader(payload)
	}

	r.Header.Add("User-Agent", "xmndr/amznspapi-go/0.0")
	if r.Header.Get("Authorization")[:7] == "Bearer " {
		r.Header.Add("X-Amz-Access-Token", r.Header.Get("Authorization")[7:])
	}

	_, err := t.aws4Signer.Sign(r, body, "execute-api", t.region, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return t.base().RoundTrip(r)
}

func (t *transport) base() http.RoundTripper {
	if t.originalTransport != nil {
		return t.originalTransport
	}
	return http.DefaultTransport
}

func (sp *Client) signer() (*v4.Signer, error) {
	awsSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			sp.config.accessKeyId,
			sp.config.secretKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	role, err := sts.New(awsSession).AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(sp.config.roleArn),
		RoleSessionName: aws.String(sp.config.roleSessionName),
	})
	if err != nil {
		return nil, err
	}

	signer := v4.NewSigner(
		credentials.NewStaticCredentials(
			*role.Credentials.AccessKeyId,
			*role.Credentials.SecretAccessKey,
			*role.Credentials.SessionToken,
		),
		func(s *v4.Signer) {
			s.DisableURIPathEscaping = true
		},
	)

	return signer, nil
}
