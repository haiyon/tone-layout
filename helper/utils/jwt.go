package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// ErrNeedTokenProvider - need token provider
	ErrNeedTokenProvider = "can not sign token without token provider"
	// ErrInvalidToken - invalid token
	ErrInvalidToken = "invalid token"
	// ErrTokenParsing - token parsing error
	ErrTokenParsing = "token parsing error"
)

// Token - token body
type Token struct {
	ID      string         `json:"id"`
	Payload map[string]any `json:"payload"`
	Subject string         `json:"subject"`
	Expire  time.Duration  `json:"expire"`
}

// generateToken - Generate token
func generateToken(key string, token *Token) (string, error) {
	claims := &jwt.MapClaims{
		// "token_id": token.ID,
		"subject": token.Subject,
		"payload": token.Payload,
		"expire":  time.Now().Add(token.Expire).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(key))
	if err != nil {
		return "", errors.New(ErrNeedTokenProvider)
	}
	return tokenString, nil
}

// ValidateToken - Validate Token
func ValidateToken(key, token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, errors.New(ErrTokenParsing)
	}
	return t, nil
}

// DecodeToken - Decode token
func DecodeToken(key, token string) (map[string]any, error) {
	t, err := ValidateToken(key, token)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New(ErrInvalidToken)
	}
	return t.Claims.(jwt.MapClaims), nil
}

// GenerateToken - Generate token
func GenerateToken(key string, payload map[string]any, tokenID ...string) (string, string) {
	// access token payload
	accessPayload := payload
	// access token
	accessToken, _ := GenerateAccessToken(key, accessPayload)

	// refresh token
	if len(tokenID) > 0 {
		refreshPayload := map[string]any{
			"user_id":  payload["user_id"].(string),
			"token_id": tokenID[0],
		}
		refreshToken, _ := GenerateRefreshToken(key, refreshPayload)
		return accessToken, refreshToken
	}

	return accessToken, ""
}

// RefreshToken - Refresh token
func RefreshToken(key string, payload map[string]any, tokenID string, originalRefreshToken string, refreshTokenExpire int64) (string, string) {
	now := time.Now().Unix()
	diff := refreshTokenExpire - now

	refreshToken := originalRefreshToken
	accessPayload := payload
	accessToken, _ := GenerateAccessToken(key, accessPayload)
	if diff < 60*60*24*15 {
		refreshPayload := map[string]any{
			"user_id":  payload["user_id"].(string),
			"token_id": tokenID,
		}

		refreshToken, _ = GenerateRefreshToken(key, refreshPayload)
	}

	return accessToken, refreshToken
}

// RegisterToken - Register token
func RegisterToken(key, userID string, tokenID ...string) string {
	// register token payload
	payload := map[string]any{
		"user_id": userID,
	}
	registerToken, _ := GenerateRegisterToken(key, payload)

	// refresh token
	if len(tokenID) > 0 {
		payload = map[string]any{
			"user_id":  userID,
			"token_id": tokenID[0],
		}
		registerToken, _ = GenerateRegisterToken(key, payload)
		return registerToken
	}

	return registerToken
}

// GenerateAccessToken - Generate access token, 360 minute
func GenerateAccessToken(key string, payload map[string]any, subject ...string) (string, error) {
	defaultSubject := "access_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(key, &Token{
		Payload: payload,
		Subject: defaultSubject,
		Expire:  time.Hour * 24,
	})
}

// GenerateRegisterToken - Generate register, 60 minute
func GenerateRegisterToken(key string, payload map[string]any, subject ...string) (string, error) {
	defaultSubject := "register_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(key, &Token{
		Payload: payload,
		Subject: defaultSubject,
		Expire:  time.Hour * 1,
	})
}

// GenerateRefreshToken - Generate refresh token, 10080 minute, 7 day
func GenerateRefreshToken(key string, payload map[string]any, subject ...string) (string, error) {
	defaultSubject := "refresh_token"
	if len(subject) > 0 {
		defaultSubject = subject[0]
	}
	return generateToken(key, &Token{
		Payload: payload,
		Subject: defaultSubject,
		Expire:  time.Hour * 24 * 7,
	})
}
