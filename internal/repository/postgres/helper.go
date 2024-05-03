package postgres

import (
	"strings"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/spf13/cast"
)

func buildQueryGetListCats(req entity.GetListCatRequest) (string, []interface{}) {
	var queryBuilder strings.Builder
	var args []interface{}

	queryBuilder.WriteString("SELECT * FROM cats WHERE 1=1")

	if req.ID != "" {
		queryBuilder.WriteString(" AND id = ?")
		args = append(args, cast.ToInt(req.ID))
	}

	if req.Race != "" {
		queryBuilder.WriteString(" AND race = ?")
		args = append(args, req.Race)
	}

	if req.Sex != "" {
		queryBuilder.WriteString(" AND sex = ?")
		args = append(args, req.Sex)
	}

	if req.Match != "" {
		queryBuilder.WriteString(" AND is_already_matched = ?")
		args = append(args, cast.ToBool(req.Match))
	}

	if req.Age != "" {
		switch req.AgeOperator {
		case "=>":
			queryBuilder.WriteString(" AND age > ?")
		case "<=":
			queryBuilder.WriteString(" AND age < ?")
		case "=":
			queryBuilder.WriteString(" AND age = ?")
		}

		args = append(args, cast.ToInt(req.AgeValue))
	}

	if req.Owned != "" {
		if cast.ToBool(req.Owned) {
			queryBuilder.WriteString(" AND user_id = ?")
		} else {
			queryBuilder.WriteString(" AND user_id != ?")
		}
		args = append(args, req.UserID)
	}

	if req.Search != "" {
		queryBuilder.WriteString(" AND name = ?")
		args = append(args, req.Search)
	}

	queryBuilder.WriteString(" AND deleted_at IS NULL")
	queryBuilder.WriteString(" ORDER BY created_at DESC LIMIT ? OFFSET ?")
	args = append(args, cast.ToInt(req.Limit), cast.ToInt(req.Offset))

	return queryBuilder.String(), args
}
