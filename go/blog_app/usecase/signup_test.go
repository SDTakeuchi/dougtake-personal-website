package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model/uuid"
	mockrepo "blog_app/domain/repository/mock"
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_signupImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input SignupInput
	}
	testUsers := []struct {
		name     string
		email    string
		password string
	}{
		{
			name:     "John Valid",
			email:    "John12.34@example.com",
			password: "validPassword9845",
		},
		{
			name:     "Carly Invalid",
			email:    "carly@example.com-ish",
			password: "shortPS",
		},
	}
	testUUID := uuid.New()
	// randomUsers := testutil.GenRandomUsers(2)
	tests := []struct {
		name           string
		args           args
		buildStubsUser func(mockUser *mockrepo.MockUser)
		want           *SignupOutput
		wantErr        bool
	}{
		{
			"success",
			args{
				context.Background(),
				SignupInput{
					Name:        testUsers[0].name,
					Email:       testUsers[0].email,
					RawPassword: testUsers[0].password,
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(
						modelimpl.UserFromRecord(
							postgres.User{
								ID:       testUUID,
								Name:     testUsers[0].name,
								Email:    testUsers[0].email,
								Password: testUsers[0].password,
							},
						),
						nil,
					)
			},
			&SignupOutput{
				ID:    testUUID,
				Name:  testUsers[0].name,
				Email: testUsers[0].email,
			},
			false,
		},
		{
			"fail/password-less-than-8-chars",
			args{
				context.Background(),
				SignupInput{
					Name:        testUsers[1].name,
					Email:       testUsers[0].email,
					RawPassword: testUsers[1].password,
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			nil,
			true,
		},
		{
			"fail/invalid-email",
			args{
				context.Background(),
				SignupInput{
					Name:        testUsers[1].name,
					Email:       testUsers[1].email,
					RawPassword: testUsers[0].password,
				},
			},
			func(mockUser *mockrepo.MockUser) {
				mockUser.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockUser := mockrepo.NewMockUser(ctrl)
			tt.buildStubsUser(mockUser)

			signup := NewSignup(mockUser)

			got, err := signup.Execute(tt.args.ctx, tt.args.input)

			if (err != nil) != tt.wantErr {
				fmt.Printf("got.Email: %+v\n", got.Email)
				t.Errorf("signupImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("signupImpl.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
