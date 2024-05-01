package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) InsertMatchCat(ctx context.Context, req entity.MatchCat) (err error) {
	query := `INSERT INTO match_cats (
        issued_by_id,
        target_user_id,
        match_cat_id,
        user_cat_id,
        message,
				status,
        created_at,
        updated_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING *`

	_, err = r.db.ExecContext(ctx,
		query,
		req.IssuedByID,
		req.TargetUserID,
		req.MatchCatID,
		req.UserCatID,
		req.Message,
		entity.MatchCatStatusPending,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][InsertMatchCat] failed to query, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) GetListMatchCat(ctx context.Context) (result []entity.MatchCat, err error) {
	query := `
		SELECT * FROM match_cats
	`

	err = r.db.SelectContext(ctx, &result, query)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Repository][MatchCat][GetListMatchCat] failed to get list match cat, err: ", err.Error())
		return
	}

	return
}
