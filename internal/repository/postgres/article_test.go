package postgres

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/backendmagang/project-1/models/entity"
	"github.com/lib/pq"
)

func Test_repository_GetArticles(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.GetArticlesRequest
	}
	tests := []struct {
		name string
		args args
		want []entity.ArticleResponse
		mock func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.GetArticlesRequest{Limit: 5, Offset: 0, SortBy: "asc", OrderBy: "", OrderByBool: true},
			},
			want: []entity.ArticleResponse{
				{ID: "a-1", Title: "Article 1", Slug: "article-1", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryList: []entity.CategoryResponse{{ID: 1, Title: "Category 1", Slug: "category-1"}}},
				{ID: "a-2", Title: "Article 2", Slug: "article-2", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryList: []entity.CategoryResponse{{ID: 1, Title: "Category 1", Slug: "category-1"}}},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow("a-1", "Article 1", "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`).
					AddRow("a-2", "Article 2", "article-2", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`)
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at, json_agg(c) AS categories FROM articles a JOIN categories c ON c.id = ANY(a.category_id) GROUP BY a.id LIMIT $1 OFFSET $2`)).
					WillReturnRows(rows)
			},
		},
		{
			name: "failed to scan rows",
			args: args{
				ctx: ctx,
				req: entity.GetArticlesRequest{Limit: 5, Offset: 0, SortBy: "asc", OrderBy: "", OrderByBool: true},
			},
			want: []entity.ArticleResponse{},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow(nil, "Article 1", "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`).
					RowError(1, errors.New("error while scanning rows"))
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).
					WillReturnRows(rows)
			},
		},
		{
			name: "failed to exec query",
			args: args{
				ctx: ctx,
				req: entity.GetArticlesRequest{Limit: 5, Offset: 0, SortBy: "asc", OrderBy: "", OrderByBool: true},
			},
			want: []entity.ArticleResponse{},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow(nil, "Article 1", "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`).
					RowError(1, errors.New("erro"))
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECTED WRONG QUERY`)).
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
			got, _ := r.GetArticles(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nrepository.GetArticles() \nGOT  = %v, \nWANT = %v\n", got, tt.want)
			}
		})
	}
}

func Test_repository_GetArticleDetails(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.GetArticleDetailsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.ArticleResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.GetArticleDetailsRequest{ID: "a-1"},
			},
			want:    entity.ArticleResponse{ID: "a-1", Title: "article 1", Slug: "article-1", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryList: []entity.CategoryResponse{{ID: 1, Title: "Category 1", Slug: "category-1"}}},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow("a-1", "article 1", "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`)

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at, ARRAY_AGG(json_build_object('id', c.id, 'title', c.title, 'slug', c.slug)) as categories FROM articles a JOIN categories c ON c.id = ANY(a.category_id) WHERE a.id = $1 GROUP BY a.id`)).
					WillReturnRows(rows)
			},
		},
		{
			name: "failed to run query",
			args: args{
				ctx: ctx,
				req: entity.GetArticleDetailsRequest{ID: "a-1"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at,	ARRAY_AGG(json_build_object('id', c.id, 'title', c.title, 'slug', c.slug)) as categories FROM articles a JOIN categories c ON c.id = ANY(a.category_id) WHERE a.id = $1 GROUP BY a.id`)).
					WithArgs(1).WillReturnError(errors.New("failed to run query"))
			},
		},
		{
			name: "failed to run scan column",
			args: args{
				ctx: ctx,
				req: entity.GetArticleDetailsRequest{ID: "a-1"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow("a-1", nil, "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`)

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at, ARRAY_AGG(json_build_object('id', c.id, 'title', c.title, 'slug', c.slug)) as categories FROM articles a JOIN categories c ON c.id = ANY(a.category_id) WHERE a.id = $1 GROUP BY a.id`)).
					WithArgs(1).WillReturnRows(rows).WillReturnError(errors.New("failed to scan column"))
			},
		},
		{
			name: "failed to unmarshal categories",
			args: args{
				ctx: ctx,
				req: entity.GetArticleDetailsRequest{ID: "a-1"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "slug", "html_content", "metadata", "created_at", "updated_at", "categories"}).
					AddRow("a-1", "article 1", "article-1", "content", `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z", `[{"id": 1, "title": "Category 1", "slug": "category-1", "created_at": "2023-06-18T19:36:03.994148", "updated_at": "2023-06-18T19:36:03.994148"}]`)

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT a.id, a.title, a.slug, a.html_content, a.metadata, a.created_at, a.updated_at, ARRAY_AGG(json_build_object('id', c.id, 'title', c.title, 'slug', c.slug)) as categories FROM articles a JOIN categories c ON c.id = ANY(a.category_id) WHERE a.id = $1 GROUP BY a.id`)).
					WithArgs(1).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.GetArticleDetails(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetArticleDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetArticleDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_InsertArticle(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.InsertArticleRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.ArticleResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: entity.InsertArticleRequest{ID: "a-1", Title: "article 1", Slug: "article-1", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryIDs: []int{1}},
			},
			want:    entity.ArticleResponse{ID: "a-1", Title: "article 1", Slug: "article-1", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryList: []entity.CategoryResponse{{ID: 1, Title: "Category 1", Slug: "category-1"}}},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO articles (id, title, slug, html_content, category_id, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)).
					WithArgs("a-1", "article 1", "article-1", "content", pq.Array([]int{1}), `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
					WillReturnResult(sqlmock.NewResult(1, 1))

				rows := sqlmock.NewRows([]string{"id", "title", "slug"}).
					AddRow(1, "Category 1", "category-1")

				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, slug FROM categories WHERE id IN (
  								SELECT unnest($1::integer[]));`)).
					WithArgs(pq.Array([]int{1})).
					WillReturnRows(rows)
			},
		},
		{
			name: "failed to insert article",
			args: args{
				ctx: ctx,
				req: entity.InsertArticleRequest{ID: "a-1", Title: "article 1", Slug: "article-1", HTMLContent: "content", Metadata: `{"key": "value2"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:29:00Z", CategoryIDs: []int{1}},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`INSERT INTO articles (id, title, slug, html_content, category_id, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)).
					WithArgs("a-1", "article 1", "article-1", "content", pq.Array([]int{1}), `{"key": "value2"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:29:00Z").
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
			got, err := r.InsertArticle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_UpdateArticle(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx     context.Context
		article entity.UpdateArticleRequest
	}
	tests := []struct {
		name    string
		args    args
		want    entity.ArticleResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				article: entity.UpdateArticleRequest{ID: "a-1", Title: "Updated Article", Slug: "updated-article", HTMLContent: "updated content", CategoryIDs: []int{1, 2}, Metadata: `{"key": "value"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:30:00Z"},
			},
			want:    entity.ArticleResponse{ID: "a-1", Title: "Updated Article", Slug: "updated-article", HTMLContent: "updated content", Metadata: `{"key": "value"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:30:00Z"},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, created_at = $6, updated_at = $7 WHERE id = $8`)).
					WithArgs("Updated Article", "updated-article", "updated content", pq.Array([]int{1, 2}), `{"key": "value"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:30:00Z", "a-1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed to update article",
			args: args{
				ctx:     ctx,
				article: entity.UpdateArticleRequest{ID: "a-1", Title: "Updated Article", Slug: "updated-article", HTMLContent: "updated content", CategoryIDs: []int{1, 2}, Metadata: `{"key": "value"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:30:00Z"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, created_at = $6, updated_at = $7 WHERE id = $8`)).
					WithArgs("Updated Article", "updated-article", "updated content", pq.Array([]int{1, 2}), `{"key": "value"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:30:00Z", "a-1").
					WillReturnError(errors.New("failed to update article"))
			},
		},
		{
			name: "failed to update article",
			args: args{
				ctx:     ctx,
				article: entity.UpdateArticleRequest{ID: "a-1", Title: "Updated Article", Slug: "updated-article", HTMLContent: "updated content", CategoryIDs: []int{1, 2}, Metadata: `{"key": "value"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:30:00Z"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, created_at = $6, updated_at = $7 WHERE id = $8`)).
					WithArgs("Updated Article", "updated-article", "updated content", pq.Array([]int{1, 2}), `{"key": "value"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:30:00Z", "a-1").
					WillReturnError(errors.New("failed to update article"))
			},
		},
		{
			name: "error no rows affected",
			args: args{
				ctx:     ctx,
				article: entity.UpdateArticleRequest{ID: "a-1", Title: "Updated Article", Slug: "updated-article", HTMLContent: "updated content", CategoryIDs: []int{1, 2}, Metadata: `{"key": "value"}`, CreatedAt: "2022-12-01T20:29:00Z", UpdatedAt: "2022-12-01T20:30:00Z"},
			},
			want:    entity.ArticleResponse{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`UPDATE articles SET title = $1, slug = $2, html_content = $3, category_id = $4, metadata = $5, created_at = $6, updated_at = $7 WHERE id = $8`)).
					WithArgs("Updated Article", "updated-article", "updated content", pq.Array([]int{1, 2}), `{"key": "value"}`, "2022-12-01T20:29:00Z", "2022-12-01T20:30:00Z", "a-1").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			got, err := r.UpdateArticle(tt.args.ctx, tt.args.article)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.UpdateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_DeleteArticle(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		req entity.DeleteArticleRequest
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
				req: entity.DeleteArticleRequest{
					ID: "a-1",
				},
			},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`DELETE FROM articles WHERE id = $1`)).
					WithArgs("a-1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "failed to exec query",
			args: args{
				ctx: ctx,
				req: entity.DeleteArticleRequest{
					ID: "a-1",
				},
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`DELETE FROM articles WHERE id = $1`)).
					WithArgs("a-1").
					WillReturnError(errors.New("failed to exec query"))
			},
		},
		{
			name: "no rows affected",
			args: args{
				ctx: ctx,
				req: entity.DeleteArticleRequest{
					ID: "a-1",
				},
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(regexp.QuoteMeta(`DELETE FROM articles WHERE id = $1`)).
					WithArgs("a-1").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				db: db,
			}
			err := r.DeleteArticle(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
