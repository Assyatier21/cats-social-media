package lib

import "errors"

var (
	ERROR_FORMAT_ID         = "id must be an integer"
	ERROR_FORMAT_LIMIT      = "limit must be an integer"
	ERROR_FORMAT_OFFSET     = "offset must be an integer"
	ERROR_FORMAT_ORDER_BY   = "order by must be a boolean"
	ERROR_FORMAT_SLUG       = "incorrect slug format"
	ERROR_FORMAT_CATEGORYID = "categoryid must be an integer"
	ERROR_FORMAT_METADATA   = "incorrect metadata format"

	ERROR_EMPTY_ID          = "id failed to be empty"
	ERROR_EMPTY_LIMIT       = "limit failed to be empty"
	ERROR_EMPTY_OFFSET      = "offset failed to be empty"
	ERROR_EMPTY_TITLE       = "title failed to be empty"
	ERROR_EMPTY_SLUG        = "slug failed to be empty"
	ERROR_EMPTY_HTMLCONTENT = "htmlcontent failed to be empty"
	ERROR_EMPTY_CATEGORYID  = "categoryid failed to be empty"
	ERROR_EMPTY_METADATA    = "metadata failed to be empty"
	ERROR_BAD_REQUEST       = "bad request"

	ERROR_FORMAT_EMPTY_ID         = "id must be an integer and failed to be empty"
	ERROR_FORMAT_EMPTY_SLUG       = "incorrect slug format or slug failed to be empty"
	ERROR_FORMAT_EMPTY_CATEGORYID = "incorrect categoryid format or categoryid failed to be empty"
	ERROR_FORMAT_EMPTY_METADATA   = "incorrect metadata format or metadata failed to be empty"

	ERR_VALIDATION_ID    = "id must be an integer and can't be empty"
	ERR_PHONE_REGISTERED = "phone number already registered"
	ERR_DATA_NOT_FOUND   = "record not found"
	ERR_PHONE_OR_PASS    = "phone or password are incorrect"
	ERR_ROLE_FORMAT      = "role chosen aren't defined"
	ERR_GENERATE_JWT     = "failed to generate jwt token"
	ERR_EMPTY_TOKEN      = "token is empty"
	ERR_EMPTY_PAYLOAD    = "payload can't be empty"
	ERR_BINDING          = "failed to bind request"
	ERR_TOKEN_EXPIRED    = "token already expired"

	ErrorNotFound       = errors.New("data not found")
	ErrorNoRowsAffected = errors.New("no rows affected")
	ErrorNoRowsResult   = errors.New("no rows in result set")
)
