package signer

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const TokenIssuer = "Avila Tek, C.A."

const EmailTypeClaimsKey = "email_typ"

type UserRole uint8

const (
	_ UserRole = iota
	StandardUserRole
	AdminUserRole
)

var userRoleStrings = map[string]UserRole{
	"user":  StandardUserRole,
	"admin": AdminUserRole,
}

func (l UserRole) String() string {
	for s, v := range userRoleStrings {
		if l == v {
			return s
		}
	}
	return "invalid"
}

func ParseUserRole(s string) UserRole {
	return userRoleStrings[s]
}

type ResendEmailClaims struct {
	ExpiresAt int64  `json:"exp,omitempty"`
	EmailType string `json:"email_typ,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type JWTHandler struct {
	SecretKey []byte
}

func NewJWTHandler(secretKey []byte) JWTHandler {
	return JWTHandler{
		SecretKey: secretKey,
	}
}

func (h *JWTHandler) Sign(claims jwt.Claims) (string, error) {
	var (
		tk     *jwt.Token
		signed string
	)

	tk = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := tk.SignedString(h.SecretKey)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func (h *JWTHandler) validate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to "SigningMethodHMAC" or your expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return h.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Validates token string and then extracts and return subject from claims. If all is good, err will be nil
func (h *JWTHandler) ValidateAndGetSubject(tokenString string) (string, error) {
	token, err := h.validate(tokenString)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return (claims)["sub"].(string), nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}

// Validates token string and then extracts and return map claims. If all is good, err will be nil
func (h *JWTHandler) ValidateAndGetClaimsMap(tokenString string) (jwt.MapClaims, error) {

	token, err := h.validate(tokenString)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, fmt.Errorf("invalid token")
	}
}
