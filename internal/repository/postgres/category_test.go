package postgres

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/lib/pq"
)

func Test_repository_GetCategoryTree(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.GetCategoriesRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.GetCategoriesRequest{
					Limit:  10,
					Offset: 0,
				},
			},
			want: []entity.Category{
				{ID: 1, Title: "Category 1", Slug: "category-1", CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z"},
				{ID: 2, Title: "Category 2", Slug: "category-2", CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z"},
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "Category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					AddRow(2, "Category 2", "category-2", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug, created_at, updated_at FROM categories ORDER BY id LIMIT $1 OFFSET $2`)).
					WithArgs(10, 0).
					WillReturnRows(rows)
			},
		},
		{
			name: "error exec query",
			args: args{
				ctx: ctx,
				req: entity.GetCategoriesRequest{
					Limit:  10,
					Offset: 0,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug, created_at, updated_at FROM categories LIMIT $1 OFFSET $2`)).
					WithArgs(10, 0).
					WillReturnError(errors.New("failed to query categories"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryTree(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetCategoryDetails(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.GetCategoryDetailsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Category
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.GetCategoryDetailsRequest{
					ID: 1,
				},
			},
			want: entity.Category{
				ID:        1,
				Title:     "Category 1",
				Slug:      "category-1",
				CreatedAt: "2022-12-01T20:29:00Z",
				UpdatedAt: "2022-12-01T20:29:00Z",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "created_at", "updated_at"}).
					AddRow(1, "Category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z")

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug, created_at, updated_at FROM categories WHERE id = $1`)).
					WithArgs(1).
					WillReturnRows(rows)
			},
		},
		{
			name: "error scanning category",
			args: args{
				ctx: ctx,
				req: entity.GetCategoryDetailsRequest{
					ID: 1,
				},
			},
			want:    entity.Category{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug, created_at, updated_at FROM categories WHERE id = $1`)).
					WithArgs(1).
					WillReturnError(errors.New("failed to scan category"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetCategoryDetails(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoryDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoryDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetCategoriesByIDs(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx         context.Context
		categoryIDs []int
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.CategoryResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:         ctx,
				categoryIDs: []int{1, 2},
			},
			want: []entity.CategoryResponse{
				{ID: 1, Title: "Category 1", Slug: "category-1"},
				{ID: 2, Title: "Category 2", Slug: "category-2"},
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug"}).
					AddRow(1, "Category 1", "category-1").
					AddRow(2, "Category 2", "category-2")

				sqlMock.ExpectQuery(`SELECT id, title, slug FROM categories WHERE id IN (.*)`).
					WithArgs(pq.Array([]int{1, 2})).
					WillReturnRows(rows)
			},
		},
		{
			name: "error getting categories",
			args: args{
				ctx:         ctx,
				categoryIDs: []int{1, 2},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug FROM categories WHERE id IN (.*)`)).
					WithArgs(pq.Array([]int{1, 2})).
					WillReturnError(errors.New("failed to get categories"))
			},
		},
		{
			name: "error scanning category",
			args: args{
				ctx:         ctx,
				categoryIDs: []int{1, 2},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug"}).
					AddRow(1, "Category 1", "category-1").
					AddRow(2, "Category 2", "category-2").
					RowError(1, errors.New("failed to scan category"))

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug FROM categories WHERE id IN (.*)`)).
					WithArgs(pq.Array([]int{1, 2})).
					WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetCategoriesByIDs(tt.args.ctx, tt.args.categoryIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetCategoriesByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetCategoriesByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_InsertCategory(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx      context.Context
		category entity.InsertCategoryRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.InsertCategoryRequest
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				category: entity.InsertCategoryRequest{
					Title:     "Category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want: entity.InsertCategoryRequest{
				ID:        1,
				Title:     "Category 1",
				Slug:      "category-1",
				CreatedAt: "2022-12-01T20:29:00Z",
				UpdatedAt: "2022-12-01T20:29:00Z",
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				sqlMock.ExpectQuery(`INSERT INTO categories (.*)`).
					WithArgs("Category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					WillReturnRows(rows)
			},
		},
		{
			name: "failed to insert article",
			args: args{
				ctx: ctx,
				category: entity.InsertCategoryRequest{
					Title:     "Category 1",
					Slug:      "category-1",
					CreatedAt: "2022-12-01T20:29:00Z",
					UpdatedAt: "2022-12-01T20:29:00Z",
				},
			},
			want:    entity.InsertCategoryRequest{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(`INSERT INTO categories (.*)`).
					WithArgs("Category 1", "category-1", "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					WillReturnError(errors.New("failed to insert article"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.InsertCategory(tt.args.ctx, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_UpdateCategory(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx      context.Context
		category entity.UpdateCategoryRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				category: entity.UpdateCategoryRequest{
					ID:        1,
					Title:     "Updated Category",
					Slug:      "updated-category",
					UpdatedAt: "2022-12-02T10:00:00Z",
				},
			},
			wantErr: false,
			mock: func() {
				result := sqlmock.NewResult(1, 1)

				sqlMock.ExpectExec(`UPDATE categories SET (.*) WHERE id = (.*)`).
					WithArgs("Updated Category", "updated-category", "2022-12-02T10:00:00Z", 1).
					WillReturnResult(result)
			},
		},
		{
			name: "error execution",
			args: args{
				ctx: ctx,
				category: entity.UpdateCategoryRequest{
					ID:        1,
					Title:     "Updated Category",
					Slug:      "updated-category",
					UpdatedAt: "2022-12-02T10:00:00Z",
				},
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`UPDATE categories SET (.*) WHERE id = (.*)`).
					WithArgs("Updated Category", "updated-category", "2022-12-02T10:00:00Z", 1).
					WillReturnError(fmt.Errorf("database error"))
			},
		},
		{
			name: "no rows affected",
			args: args{
				ctx: ctx,
				category: entity.UpdateCategoryRequest{
					ID:        1,
					Title:     "Updated Category",
					Slug:      "updated-category",
					UpdatedAt: "2022-12-02T10:00:00Z",
				},
			},
			wantErr: true,
			mock: func() {
				result := sqlmock.NewResult(0, 0)

				sqlMock.ExpectExec(`UPDATE categories SET (.*) WHERE id = (.*)`).
					WithArgs("Updated Category", "updated-category", "2022-12-02T10:00:00Z", 1).
					WillReturnResult(result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			err := r.UpdateCategory(tt.args.ctx, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_repository_DeleteCategory(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx     context.Context
		request entity.DeleteCategoryRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				request: entity.DeleteCategoryRequest{
					ID: 1,
				},
			},
			wantErr: false,
			mock: func() {
				result := sqlmock.NewResult(1, 1)

				sqlMock.ExpectExec(`DELETE FROM categories WHERE id = (.*)`).
					WithArgs(1).
					WillReturnResult(result)
			},
		},
		{
			name: "failed to exec query",
			args: args{
				ctx: ctx,
				request: entity.DeleteCategoryRequest{
					ID: 1,
				},
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`DELETE FROM categories WHERE id = (.*)`).
					WithArgs(1).
					WillReturnError(errors.New("failed to exec query"))
			},
		},
		{
			name: "no rows affected",
			args: args{
				ctx: ctx,
				request: entity.DeleteCategoryRequest{
					ID: 1,
				},
			},
			wantErr: true,
			mock: func() {
				result := sqlmock.NewResult(0, 0)

				sqlMock.ExpectExec(`DELETE FROM categories WHERE id = (.*)`).
					WithArgs(1).
					WillReturnResult(result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			err := r.DeleteCategory(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
