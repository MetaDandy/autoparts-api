package auth

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
	ExpiresIn    int32
}

type CognitoProvider struct {
	client      *cognitoidentityprovider.Client
	userPoolID  string
	clientAppID string
}
