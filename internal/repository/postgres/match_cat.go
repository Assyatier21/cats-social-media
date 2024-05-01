package postgres

import (
	"context"
	"database/sql"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) FindMatchByID(ctx context.Context, id int) (entity.MatchCat, error) {
	result := entity.MatchCat{}
	query := `SELECT * FROM match_cats WHERE id = $1`

	err := r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Errorf("[Repository][MatchCat][FindMatchByID] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

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
		req.Status,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][InsertMatchCat] failed to query, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) UpdateMatchCat(ctx context.Context, req entity.MatchCat) (err error) {
	query := `UPDATE match_cats
	SET
		issued_by_id = $1,
		target_user_id = $2,
		match_cat_id = $3,
		user_cat_id = $4,
		message = $5,
		status = $6,
		updated_at = $7
	WHERE id = $8
	RETURNING *`

	_, err = r.db.ExecContext(ctx,
		query,
		req.IssuedByID,
		req.TargetUserID,
		req.MatchCatID,
		req.UserCatID,
		req.Message,
		req.Status,
		req.UpdatedAt,
		req.ID,
	)

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][UpdateMatchCat] failed to query, err: %s", err.Error())
		return
	}

	return
}
