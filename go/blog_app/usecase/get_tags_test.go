package usecase

import (
	"blog_app/domain/model"
	mockrepo "blog_app/domain/repository/mock"
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_getTagsImpl_Execute(t *testing.T) {
	type args struct {
		ctx   context.Context
		input GetTagsInput
	}

	randomTags := genRandomTags(3)

	tests := []struct {
		name          string
		args          args
		buildStubsTag func(mockTag *mockrepo.MockTag)
		want          *GetTagsOutput
		wantErr       bool
	}{
		{
			"success/get-all-tags",
			args{
				context.Background(),
				GetTagsInput{[]uint64{}},
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Times(1).
					Return(randomTags, nil)
			},
			&GetTagsOutput{randomTags},
			false,
		},
		{
			"success/get-one-tag",
			args{
				context.Background(),
				GetTagsInput{[]uint64{randomTags[0].ID()}},
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq([]uint64{randomTags[0].ID()})).
					Times(1).
					Return([]model.Tag{randomTags[0]}, nil)
			},
			&GetTagsOutput{[]model.Tag{randomTags[0]}},
			false,
		},
		{
			"success/get-one-tag",
			args{
				context.Background(),
				GetTagsInput{[]uint64{randomTags[1].ID(), randomTags[1].ID()}},
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Eq([]uint64{randomTags[1].ID(), randomTags[1].ID()})).
					Times(1).
					Return([]model.Tag{randomTags[0], randomTags[1]}, nil)
			},
			&GetTagsOutput{[]model.Tag{randomTags[0], randomTags[1]}},
			false,
		},
		{
			"fail/tag-id-not-found",
			args{
				context.Background(),
				GetTagsInput{[]uint64{100000000000}},
			},
			func(mockTag *mockrepo.MockTag) {
				mockTag.EXPECT().
					Find(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil, gorm.ErrRecordNotFound)
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockTag := mockrepo.NewMockTag(ctrl)
			tt.buildStubsTag(mockTag)

			getTags := NewGetTags(mockTag)

			got, err := getTags.Execute(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTagsImpl.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTagsImpl.Execute() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
