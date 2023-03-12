package auth

import (
	"blog_app/domain/model/auth"
	testutil "blog_app/util/test"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTIssuer(t *testing.T) {
	tokenIssuer, err := NewJWTIssuer(testutil.GenRandomChars(200))
	require.NoError(t, err)

	user := testutil.GenRandomUsers(1)[0]
	duration := time.Minute

	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	token, payload, err := tokenIssuer.Create(user.ID(), auth.AccessToken, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, token)

	payload, err = tokenIssuer.Verify(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID())
	require.Equal(t, user.ID(), payload.UserID())
	require.WithinDuration(t, issuedAt, payload.IssuedAt(), time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiresAt(), time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	tokenIssuer, err := NewJWTIssuer(testutil.GenRandomChars(32))
	require.NoError(t, err)

	user := testutil.GenRandomUsers(1)[0]
	token, payload, err := tokenIssuer.Create(user.ID(), auth.AccessToken, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, token)

	payload, err = tokenIssuer.Verify(token)
	require.Error(t, err)
	require.EqualError(t, err, auth.ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	user := testutil.GenRandomUsers(1)[0]
	payload := NewPayload(user.ID(), auth.AccessToken, time.Minute)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, ClaimsFromPayload(payload))
	// UnsafeAllowNoneSignatureType should be only used for testing purposes because it is not safe
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	tokenIssuer, err := NewJWTIssuer(testutil.GenRandomChars(32))
	require.NoError(t, err)

	payload, err = tokenIssuer.Verify(token)
	require.Error(t, err)
	require.EqualError(t, err, auth.ErrInvalidToken.Error())
	require.Nil(t, payload)
}
