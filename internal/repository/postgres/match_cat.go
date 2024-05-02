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

func (r *repository) GetListMatchCat(ctx context.Context, req entity.GetListMatchCatRequest) (result []entity.GetListMatchCatQueryResponse, err error) {
	query := `
		SELECT
			match_cats.id AS match_id,
			match_cats.message,
			match_cats.created_at AS match_created_at,
			issued_by.name AS issuedBy_name,
			issued_by.email AS issuedBy_email,
			issued_by.created_at AS issuedBy_createdAt,
			match_cat.id AS match_cat_id,
			match_cat.name AS match_cat_name,
			match_cat.race AS match_cat_race,
			match_cat.sex AS match_cat_sex,
			match_cat.description AS match_cat_description,
			match_cat.age AS match_cat_age,
			match_cat.images AS match_cat_images,
			match_cat.is_already_matched AS match_cat_hasMatched,
			match_cat.created_at AS match_cat_createdAt,
			user_cat.id AS user_cat_id,
			user_cat.name AS user_cat_name,
			user_cat.race AS user_cat_race,
			user_cat.sex AS user_cat_sex,
			user_cat.description AS user_cat_description,
			user_cat.age AS user_cat_age,
			user_cat.images AS user_cat_images,
			user_cat.is_already_matched AS user_cat_hasMatched,
			user_cat.created_at AS user_cat_createdAt
		FROM
			match_cats
		JOIN
			cats AS match_cat ON match_cats.match_cat_id = match_cat.id
		JOIN
			cats AS user_cat ON match_cats.user_cat_id = user_cat.id
		JOIN
			users AS issued_by ON match_cats.issued_by_id = issued_by.id
		WHERE
    	match_cats.issued_by_id = $1 OR match_cats.target_user_id = $1
		ORDER BY
			match_cats.created_at DESC; -- ordered by newest first
	`

	err = r.db.SelectContext(ctx, &result, query, req.UserID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("[Repository][MatchCat][GetListMatchCat] failed to get list match cat, err: ", err.Error())
		return
	}

	return
}
