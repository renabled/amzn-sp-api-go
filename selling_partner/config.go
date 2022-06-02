package selling_partner

import (
	"strings"

	"github.com/google/uuid"
)

type Config struct {
	// Your AWS access key Id.
	accessKeyId string
	// Your AWS secret access key.
	secretKey string
	// The AWS region to which you are directing your call.
	// options: na, eu or fe
	region string
	// The ARN of the IAM role.
	roleArn string
	// An identifier for the session that you define. (UUID)
	roleSessionName string
	// Your LWA client identifier.
	clientID string
	// Your LWA client secret.
	clientSecret string
	// The LWA refresh token. Get this value when the selling partner authorizes your application.
	// Include refreshToken if the operation that you call in the following step requires selling
	// partner authorization. All operations that are not grantless operations require selling
	// partner authorization. If you include refreshToken, do not include withScopes.
	refreshToken string
	// The scope of the LWA authorization grant.
	// Include withScopes if the operation that you call in the following step
	// is a grantless operation. If you include withScopes, do not include refreshToken.
	withScopes []string
	//	The LWA authentication server URI.
	endpoint string
}

func NewConfig(
	accessKeyId, secretKey, roleArn, refreshToken, clientID, clientSecret, region string,
) *Config {
	return &Config{
		accessKeyId:     accessKeyId,
		secretKey:       secretKey,
		region:          Regions[strings.ToUpper(region)].Region,
		roleArn:         roleArn,
		roleSessionName: uuid.New().String(),
		clientID:        clientID,
		clientSecret:    clientSecret,
		refreshToken:    refreshToken,
		withScopes: []string{
			"sellingpartnerapi::notifications",
			"sellingpartnerapi::migration",
		},
		endpoint: Regions[strings.ToUpper(region)].Endpoint,
	}
}
