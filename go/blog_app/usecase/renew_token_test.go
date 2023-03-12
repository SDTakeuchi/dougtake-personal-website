package usecase

import (
	"blog_app/adapter/config"
	authimpl "blog_app/adapter/domain_impl/model/auth"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/auth"
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_renewTokenImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		user  model.User
		input RenewTokenInput
	}
	users := testutil.GenRandomUsers(2)

	tokenIssuer, _ := authimpl.NewJWTIssuer("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	accessToken, _, _ := tokenIssuer.Create(
		users[0].ID(),
		auth.AccessToken,
		config.Get().Token.AccessTokenDuration,
	)
	refreshToken, refreshTokenPayload, _ := tokenIssuer.Create(
		users[0].ID(),
		auth.RefreshToken,
		config.Get().Token.RefreshTokenDuration,
	)
	expiredRrefreshToken, expiredRrefreshTokenPayload, _ := tokenIssuer.Create(
		users[0].ID(),
		auth.RefreshToken,
		0,
	)
	now := time.Now()

	tests := []struct {
		name              string
		args              args
		buildStubsSession func(mockUser *mockrepo.MockSession)
		tokenIssuer       auth.TokenIssuer
		want              *RenewTokenOutput
		wantErr           bool
	}{
		{
			"success",
			args{
				context.Background(),
				users[0],
				RenewTokenInput{
					refreshToken,
				},
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Get(gomock.Any(), refreshTokenPayload.ID().String()).
					Times(1).
					Return(
						authimpl.SessionFromRecord(
							postgres.Session{
								ID:           refreshTokenPayload.ID().String(),
								UserID:       users[0].ID().String(),
								RefreshToken: refreshToken,
								ExpiresAt:    now.Add(config.Get().Token.RefreshTokenDuration),
								CreatedAt:    now,
							},
						), nil)
			},
			tokenIssuer,
			&RenewTokenOutput{AccessTokenExpiresAt: now},
			false,
		},
		{
			"fail/access-token-used",
			args{
				context.Background(),
				users[0],
				RenewTokenInput{
					accessToken,
				},
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Get(gomock.Any(), gomock.Any()).
					Times(0)
			},
			tokenIssuer,
			nil,
			true,
		},
		{
			"fail/token-expired",
			args{
				context.Background(),
				users[0],
				RenewTokenInput{
					expiredRrefreshToken,
				},
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Get(gomock.Any(), expiredRrefreshTokenPayload.ID().String()).
					Times(1).
					Return(
						authimpl.SessionFromRecord(
							postgres.Session{
								ID:           expiredRrefreshTokenPayload.ID().String(),
								UserID:       users[0].ID().String(),
								RefreshToken: refreshToken,
								ExpiresAt:    now.Add(config.Get().Token.RefreshTokenDuration),
								CreatedAt:    now,
							},
						), nil)
			},
			tokenIssuer,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockSession := mockrepo.NewMockSession(ctrl)
			tt.buildStubsSession(mockSession)

			renewToken := NewRenewToken(tt.tokenIssuer, mockSession)

			got, err := renewToken.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("renewTokenImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isValidRenewTokenResponse(got, tt.want) {
				t.Errorf("renewTokenImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isValidRenewTokenResponse(got, want *RenewTokenOutput) bool {
	// if both got and want are nil, comparison is not needed anymore
	if got == nil && want == nil {
		return true
	}

	if want.AccessTokenExpiresAt.After(got.AccessTokenExpiresAt) &&
		got.AccessTokenExpiresAt.Add(time.Second*3).After(want.AccessTokenExpiresAt) {
		fmt.Printf(
			"different expiration date, got: %v, want: %v\n",
			got.AccessTokenExpiresAt,
			want.AccessTokenExpiresAt,
		)
		return false
	}

	if len(got.AccessToken) < 30 {
		fmt.Printf(
			"access token too short: %v\n",
			got.AccessToken,
		)
		return false
	}
	return true
}
