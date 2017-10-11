package middleware

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/rinkbase/grinder"
)

// JWTConfig holds token information
type JWTConfig struct {
	SigningKey    interface{}
	SigningMethod string // Defaults to HS256
	ParseFrom     string // Defaults to query (options: header, query)
	Claims        jwt.Claims
}

const (
	// AlgoHS256 is 256 bit encryption
	AlgoHS256 = "HS256"

	// AlgoHS512 is 512 bit encryption
	AlgoHS512 = "HS512"

	// AlgoHS1024 is 1024 bit encryptions
	AlgoHS1024 = "HS1024"
)

// TokenParser parses out token
type TokenParser func(grinder.Context) (string, error)

// DefaultJWT is the default settings for the JWT
var DefaultJWT = JWTConfig{
	SigningMethod: AlgoHS256,
}

// JWTError returns a grinder Handler when an error is occured
func JWTError(c grinder.Context) error {
	return c.JSON(500, "JWT Error")
}

// JWT default json web token handler
func JWT(c grinder.Context, handler grinder.Handler) grinder.Handler {
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	j := DefaultJWT
	j.SigningKey = []byte(config["JWT_SECRET"])

	parser := parseFromQuery()
	switch j.ParseFrom {
	case "query":
		parser = parseFromQuery()
	}

	parsed, err := parser(c)
	if err != nil {
		log.Println(err)
	}

	token := new(jwt.Token)
	if _, ok := j.Claims.(jwt.MapClaims); ok {
		token, err = jwt.Parse(parsed, func(token *jwt.Token) (interface{}, error) {
			return []byte("AllYourBase"), nil
		})
	}

	if token.Valid && err == nil {
		return handler
	}

	return JWTError
}

// expects http://domain.com?token=<HASH>
func parseFromQuery() TokenParser {
	return func(c grinder.Context) (string, error) {
		token := c.GetParam("token")
		if token == "" {
			return "", errors.New("JWT token is missing")
		}

		return token, nil
	}
}
