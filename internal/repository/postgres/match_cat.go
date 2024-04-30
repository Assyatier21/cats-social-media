package postgres

import (
	"context"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error) {
	query := `INSERT INTO match_cats (
        issued_by_id,
        target_user_id,
        match_cat_id,
        user_cat_id,
        message,
        created_at,
        updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING *`

	_, err = r.db.ExecContext(ctx,
		query,
		req.IssuedByID,
		req.TargetUserID,
		req.MatchCatID,
		req.UserCatID,
		req.Message,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][InsertMatchCat] failed to query, err: %s", err.Error())
		return
	}

	return
}
