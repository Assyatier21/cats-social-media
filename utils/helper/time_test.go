package helper

import (
	"reflect"
	"testing"

	"github.com/backendmagang/project-1/models/entity"
)

func TestFormattedTime(t *testing.T) {
	type args struct {
		ts string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid timestamp",
			args: args{ts: "2023-06-27T12:34:56Z"},
			want: "2023-06-27 12:34:56",
		},
		{
			name: "Invalid timestamp",
			args: args{ts: "2023-06-27T12:34:56"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormattedTime(tt.args.ts)
			if got != tt.want {
				t.Errorf("FormattedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeArticleResponse(t *testing.T) {
	type args struct {
		article entity.ArticleResponse
	}
	tests := []struct {
		name string
		args args
		want entity.ArticleResponse
	}{
		{
			name: "Valid article response",
			args: args{
				article: entity.ArticleResponse{
					CreatedAt: "2023-06-27T12:34:56Z",
					UpdatedAt: "2023-06-28T09:12:34Z",
				},
			},
			want: entity.ArticleResponse{
				CreatedAt: "2023-06-27 12:34:56",
				UpdatedAt: "2023-06-28 09:12:34",
			},
		},
		{
			name: "Empty article response",
			args: args{
				article: entity.ArticleResponse{},
			},
			want: entity.ArticleResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTimeArticleResponse(tt.args.article)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatTimeArticleResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
