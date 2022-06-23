package util

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/auth"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func GenerateTokenAndAddCookies(ctx context.Context, value string) error {
	jwt, err := GenerateJWTToken(value)
	if err != nil {
		return fmt.Errorf(
			"generate jwt token err=%s", err.Error(),
		)
	}

	c := cookies.CreateCookie(jwt)
	cookies.AddServerCookieToCtx(ctx, c)

	return nil
}

func GenerateJWTToken(value string) (string, error) {
	secretKey, err := GetDotEnvVariable(JWTSecretEnvName)
	if err != nil {
		return "", err
	}

	expirationMinuteString, err := GetDotEnvVariable(JWTExpirationMinutesEnvName)
	if err != nil {
		return "", err
	}

	expirationMinute, err := strconv.Atoi(expirationMinuteString)
	if err != nil {
		return "", err
	}

	jwtToken, err := auth.GenerateJWTToken(value, secretKey, expirationMinute)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ParseJWTToken(tokenString string) (*auth.Claims, error) {
	secretKey, err := GetDotEnvVariable(JWTSecretEnvName)
	if err != nil {
		return &auth.Claims{}, err
	}

	return auth.ParseJWTToken(tokenString, secretKey)
}
