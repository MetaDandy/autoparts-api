package auth

// import (
// 	"context"
// 	"fmt"

// 	"github.com/gofiber/fiber/v2"
// 	"golang.org/x/oauth2"
// )

// var oauth2Cfg = &oauth2.Config{
// 	ClientID:     clientAppID,
// 	ClientSecret: clientSecret,
// 	Endpoint: oauth2.Endpoint{
// 		AuthURL:  fmt.Sprintf("https://%s.auth.%s.amazoncognito.com/oauth2/authorize", userPoolID, region),
// 		TokenURL: fmt.Sprintf("https://%s.auth.%s.amazoncognito.com/oauth2/token", userPoolID, region),
// 	},
// 	RedirectURL: "http://localhost:8000/auth/oauth/callback",
// 	Scopes:      []string{"openid", "email", "profile"},
// }

// func OAuthLogin(c *fiber.Ctx) error {
// 	state := "r4nd0m" // en prod, genera uno y guardalo en cookie/session
// 	authURL := oauth2Cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
// 	return c.Redirect(authURL)
// }

// func OAuthCallback(c *fiber.Ctx) error {
// 	code := c.Query("code")
// 	// state := c.Query("state") -> validar
// 	token, err := oauth2Cfg.Exchange(context.TODO(), code)
// 	if err != nil {
// 		return c.Status(400).SendString("OAuth exchange error: " + err.Error())
// 	}
// 	return c.JSON(fiber.Map{
// 		"access_token":  token.AccessToken,
// 		"refresh_token": token.RefreshToken,
// 		"id_token":      token.Extra("id_token"),
// 		"expiry":        token.Expiry,
// 	})
// }
