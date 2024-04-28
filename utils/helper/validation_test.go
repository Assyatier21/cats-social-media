package helper

import (
	"testing"

	"github.com/backendmagang/project-1/models/entity"
	"github.com/labstack/echo/v4"
)

type MockContext struct {
	echo.Context
}

func (mc *MockContext) Bind(i interface{}) error {
	req := i.(*entity.InsertArticleRequest)
	req.Title = "example"

	return nil
}

func TestIsValidSlug(t *testing.T) {
	validSlugs := []string{
		"hello-world",
		"foo-bar",
		"123",
		"abc-123",
		"abc-def-ghi",
	}

	for _, slug := range validSlugs {
		if !IsValidSlug(slug) {
			t.Errorf("Expected slug '%s' to be valid", slug)
		}
	}
}

func TestBuildSortByQuery(t *testing.T) {
	// Test case 1: When sort_by is "title"
	result := BuildSortByQuery("title")
	expected := "title.keyword"
	if result != expected {
		t.Errorf("BuildSortByQuery('title') returned incorrect result, got: %s, expected: %s", result, expected)
	}

	// Test case 2: When sort_by is "slug"
	result = BuildSortByQuery("slug")
	expected = "slug.keyword"
	if result != expected {
		t.Errorf("BuildSortByQuery('slug') returned incorrect result, got: %s, expected: %s", result, expected)
	}

	// Test case 3: When sort_by is "html_content"
	result = BuildSortByQuery("html_content")
	expected = "html_content.keyword"
	if result != expected {
		t.Errorf("BuildSortByQuery('html_content') returned incorrect result, got: %s, expected: %s", result, expected)
	}

	// Test case 4: When sort_by is any other value
	result = BuildSortByQuery("updated_at")
	expected = "updated_at"
	if result != expected {
		t.Errorf("BuildSortByQuery('updated_at') returned incorrect result, got: %s, expected: %s", result, expected)
	}
}

func TestBuildOrderByQuery(t *testing.T) {
	// Test case 1: When order_by is "asc"
	result := BuildOrderByQuery("asc")
	expected := true
	if result != expected {
		t.Errorf("BuildOrderByQuery('asc') returned incorrect result, got: %t, expected: %t", result, expected)
	}

	// Test case 2: When order_by is "desc"
	result = BuildOrderByQuery("desc")
	expected = false
	if result != expected {
		t.Errorf("BuildOrderByQuery('desc') returned incorrect result, got: %t, expected: %t", result, expected)
	}

	// Test case 3: When order_by is any other value
	result = BuildOrderByQuery("random")
	expected = false
	if result != expected {
		t.Errorf("BuildOrderByQuery('random') returned incorrect result, got: %t, expected: %t", result, expected)
	}
}

func TestGeArticleRequestValidation(t *testing.T) {
	// Test case 1: When req.SortBy is "title" and req.OrderBy is "asc"
	req := entity.GetArticlesRequest{
		SortBy:  "title",
		OrderBy: "asc",
		Limit:   0,
	}

	expected := entity.GetArticlesRequest{
		SortBy:      "title.keyword",
		OrderByBool: true,
		Limit:       10,
	}

	result := GeArticleRequestValidation(req)
	if result.SortBy != expected.SortBy {
		t.Errorf("GeArticleRequestValidation returned incorrect SortBy value, got: %s, expected: %s", result.SortBy, expected.SortBy)
	}
	if result.OrderByBool != expected.OrderByBool {
		t.Errorf("GeArticleRequestValidation returned incorrect OrderByBool value, got: %t, expected: %t", result.OrderByBool, expected.OrderByBool)
	}
	if result.Limit != expected.Limit {
		t.Errorf("GeArticleRequestValidation returned incorrect Limit value, got: %d, expected: %d", result.Limit, expected.Limit)
	}

	// Test case 2: When req.SortBy is "slug" and req.OrderBy is "desc"
	req = entity.GetArticlesRequest{
		SortBy:  "slug",
		OrderBy: "desc",
		Limit:   5,
	}

	expected = entity.GetArticlesRequest{
		SortBy:      "slug.keyword",
		OrderByBool: false,
		Limit:       5,
	}

	result = GeArticleRequestValidation(req)
	if result.SortBy != expected.SortBy {
		t.Errorf("GeArticleRequestValidation returned incorrect SortBy value, got: %s, expected: %s", result.SortBy, expected.SortBy)
	}
	if result.OrderByBool != expected.OrderByBool {
		t.Errorf("GeArticleRequestValidation returned incorrect OrderByBool value, got: %t, expected: %t", result.OrderByBool, expected.OrderByBool)
	}
	if result.Limit != expected.Limit {
		t.Errorf("GeArticleRequestValidation returned incorrect Limit value, got: %d, expected: %d", result.Limit, expected.Limit)
	}

	// Test case 3: When req.SortBy and req.OrderBy have invalid values
	req = entity.GetArticlesRequest{
		SortBy:  "invalid",
		OrderBy: "invalid",
		Limit:   0,
	}

	expected = entity.GetArticlesRequest{
		SortBy:      "updated_at",
		OrderByBool: false,
		Limit:       10,
	}

	result = GeArticleRequestValidation(req)
	if result.SortBy != expected.SortBy {
		t.Errorf("GeArticleRequestValidation returned incorrect SortBy value, got: %s, expected: %s", result.SortBy, expected.SortBy)
	}
	if result.OrderByBool != expected.OrderByBool {
		t.Errorf("GeArticleRequestValidation returned incorrect OrderByBool value, got: %t, expected: %t", result.OrderByBool, expected.OrderByBool)
	}
	if result.Limit != expected.Limit {
		t.Errorf("GeArticleRequestValidation returned incorrect Limit value, got: %d, expected: %d", result.Limit, expected.Limit)
	}
}
