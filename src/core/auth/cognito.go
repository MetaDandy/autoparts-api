package auth

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func NewCognitoProvider(client *cognitoidentityprovider.Client, userPoolID, clientAppID string) *CognitoProvider {
	return &CognitoProvider{client: client, userPoolID: userPoolID, clientAppID: clientAppID}
}

func (p *CognitoProvider) SignUpUser(email, password, name, phone string) (string, error) {
	_, err := p.client.AdminCreateUser(context.TODO(), &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:    aws.String(p.userPoolID),
		Username:      aws.String(email),
		MessageAction: types.MessageActionTypeSuppress, // suprime el email
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
			{Name: aws.String("name"), Value: aws.String(name)},
			{Name: aws.String("phone_number"), Value: aws.String(phone)},
		},
	})
	if err != nil {
		return "", err
	}

	_, err = p.client.AdminSetUserPassword(context.TODO(), &cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: aws.String(p.userPoolID),
		Username:   aws.String(email),
		Password:   aws.String(password),
	})
	if err != nil {
		_, delErr := p.client.AdminDeleteUser(context.TODO(), &cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: aws.String(p.userPoolID),
			Username:   aws.String(email),
		})
		if delErr != nil {
			log.Printf("Error borrando usuario tras fallo de password: %v", delErr)
		}
		return "", err
	}

	getOut, err := p.client.AdminGetUser(context.TODO(), &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(p.userPoolID),
		Username:   aws.String(email),
	})
	if err != nil {
		return "", err
	}
	for _, attr := range getOut.UserAttributes {
		if *attr.Name == "sub" {
			return *attr.Value, nil
		}
	}
	return "", errors.New("sub attribute not found")
}

func (p *CognitoProvider) SignInUser(email, password string) (AuthTokens, error) {
	out, err := p.client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{"USERNAME": email, "PASSWORD": password},
		ClientId:       aws.String(p.clientAppID),
	})
	if err != nil {
		return AuthTokens{}, err
	}
	r := out.AuthenticationResult
	return AuthTokens{*r.AccessToken, *r.RefreshToken, *r.IdToken, r.ExpiresIn}, nil
}

func (p *CognitoProvider) RefreshToken(refreshToken string) (AuthTokens, error) {
	out, err := p.client.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeRefreshTokenAuth,
		AuthParameters: map[string]string{"REFRESH_TOKEN": refreshToken},
		ClientId:       aws.String(p.clientAppID),
	})
	if err != nil {
		return AuthTokens{}, err
	}
	r := out.AuthenticationResult

	return AuthTokens{*r.AccessToken, refreshToken, *r.IdToken, r.ExpiresIn}, nil
}

func (p *CognitoProvider) DeleteUser(email string) error {
	_, err := p.client.AdminDeleteUser(context.TODO(), &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(p.userPoolID),
		Username:   aws.String(email),
	})
	return err
}
