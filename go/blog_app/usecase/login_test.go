package usecase

import (
	authimpl "blog_app/adapter/domain_impl/model/auth"
	"blog_app/adapter/log"
	"blog_app/domain/model/auth"
	mockrepo "blog_app/domain/repository/mock"
	testutil "blog_app/util/test"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_loginImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input LoginInput
	}

	randomUsers := testutil.GenRandomUsers(3)
	tokenIssuer, _ := authimpl.NewJWTIssuer("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	logger := log.NewLogger()

	tests := []struct {
		name              string
		args              args
		buildStubsUser    func(mockUser *mockrepo.MockUser)
		buildStubsSession func(mockUser *mockrepo.MockSession)
		tokenIssuer       auth.TokenIssuer
		want              *LoginOutput
		wantErr           bool
	}{
		{
			"success",
			args{
				ctx: context.Background(),
				input: LoginInput{
					Email:       randomUsers[0].Email(),
					RawPassword: testutil.SamplePassword,
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByEmail(gomock.Any(), randomUsers[0].Email()).
					Times(1).
					Return(randomUsers[0], nil)
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					// returns session record in reality
					Return(nil, nil)
			},
			tokenIssuer,
			&LoginOutput{
				UserID: randomUsers[0].ID(),
			},
			false,
		},
		{
			"fail/user-not-exit",
			args{
				ctx: context.Background(),
				input: LoginInput{
					Email:       "not.exixt@mail.com",
					RawPassword: testutil.SamplePassword,
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByEmail(gomock.Any(), "not.exixt@mail.com").
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			tokenIssuer,
			nil,
			true,
		},
		{
			"fail/incorrect-password",
			args{
				ctx: context.Background(),
				input: LoginInput{
					Email:       randomUsers[1].Email(),
					RawPassword: "incorrectPassword1122",
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					GetByEmail(gomock.Any(), randomUsers[1].Email()).
					Times(1).
					Return(randomUsers[1], nil)
			},
			func(mockUser *mockrepo.MockSession) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			tokenIssuer,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockUser := mockrepo.NewMockUser(ctrl)
			tt.buildStubsUser(mockUser)
			mockSession := mockrepo.NewMockSession(ctrl)
			tt.buildStubsSession(mockSession)

			login := NewLogin(tt.tokenIssuer, mockUser, mockSession, logger)

			got, err := login.Execute(tt.args.ctx, tt.args.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("loginImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isValidLoginResponse(got, tt.want) {
				t.Errorf("loginImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isValidLoginResponse(got, want *LoginOutput) bool {
	// if both got and want are nil, comparison is not needed anymore
	if got == nil && want == nil {
		return true
	}

	idMatches := got.UserID == want.UserID
	if !idMatches {
		fmt.Printf("got: %v, want: %v\n", got, want)
	}
	accessTokenNotEmpty := got.AccessToken != ""
	if !accessTokenNotEmpty {
		fmt.Printf("empty AccessToken: %v\n", got.AccessToken)
	}
	accessTokenIsFuture := got.AccessTokenExpiresAt.After(time.Now())
	if !accessTokenIsFuture {
		fmt.Printf("AccessToken not future: %v\n", got.AccessTokenExpiresAt)
	}
	refreshTokenNotEmpty := got.RefreshToken != ""
	if !accessTokenNotEmpty {
		fmt.Printf("empty RefreshToken: %v\n", got.RefreshToken)
	}
	refreshTokenIsFuture := got.RefreshTokenExpiresAt.After(time.Now())
	if !accessTokenIsFuture {
		fmt.Printf("RefreshToken not future: %v\n", got.RefreshTokenExpiresAt)
	}
	return idMatches && accessTokenNotEmpty && accessTokenIsFuture && refreshTokenNotEmpty && refreshTokenIsFuture
}
