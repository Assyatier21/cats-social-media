package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/backend-magang/cats-social-media/models/entity"
	"github.com/backend-magang/cats-social-media/utils/pkg"
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
		updated_at = $7,
		deleted_at = $8
	WHERE id = $9
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
		req.DeletedAt,
		req.ID,
	)

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][UpdateMatchCat] failed to query, err: %s", err.Error())
		return
	}

	return
}

func (r *repository) FindMatchCatByID(ctx context.Context, id int) (result entity.MatchCat, err error) {
	query := `SELECT * FROM match_cats WHERE id = $1`

	err = r.db.QueryRowxContext(ctx, query, id).StructScan(&result)
	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][FindMatchCatByID] failed to query, err: %s", err.Error())
		return result, err
	}

	return result, err
}

func (r *repository) ApproveMatch(ctx context.Context, matchCatId int) (err error) {
	query := `UPDATE match_cats
		SET status = $1
		WHERE id = $2 RETURNING *`

	var result sql.Result

	tx, _ := pkg.ExtractTx(ctx)
	if tx != nil {
		result, err = tx.ExecContext(ctx, query, entity.MatchCatStatusApproved, matchCatId)
	} else {
		result, err = r.db.ExecContext(ctx, query, entity.MatchCatStatusApproved, matchCatId)
	}

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][ApproveMatch] failed to query, err: %s", err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][ApproveMatch] failed to query, err: %s", err.Error())
		return
	}

	if rowsAffected == 0 {
		r.logger.Errorf("[Repository][MatchCat][ApproveMatch] failed to query, err: no rows effected")
		return
	}

	return nil
}

func (r *repository) UpdateCatsMatch(ctx context.Context, matchCat entity.MatchCat) (err error) {
	query := `UPDATE cats
		SET is_already_matched = TRUE 
		WHERE id = $1 OR id = $2`

	var result sql.Result

	tx, _ := pkg.ExtractTx(ctx)
	if tx != nil {
		result, err = tx.ExecContext(ctx, query, matchCat.UserCatID, matchCat.MatchCatID)
	} else {
		result, err = r.db.ExecContext(ctx, query, matchCat.UserCatID, matchCat.MatchCatID)
	}

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][UpdateCatsMatch] failed to query, err: %s", err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][UpdateCatsMatch] failed to query, err: %s", err.Error())
		return
	}

	if rowsAffected == 0 {
		r.logger.Errorf("[Repository][MatchCat][UpdateCatsMatch] failed to query, err: no rows effected")
		return
	}

	return nil
}

func (r *repository) DeleteOtherMatch(ctx context.Context, catId int, matchCatId int) (err error) {
	query := `UPDATE match_cats
		SET deleted_at = $1 
		WHERE (user_cat_id = $2 OR match_cat_id = $2) AND id != $3 AND status = 'pending'`

	var (
		result sql.Result
		now    = time.Now()
	)

	tx, _ := pkg.ExtractTx(ctx)
	if tx != nil {
		result, err = tx.ExecContext(ctx, query, now, catId, matchCatId)
	} else {
		_, err = r.db.ExecContext(ctx, query, now, catId, matchCatId)
	}

	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][DeleteOtherMatch] failed to query, err: %s", err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Errorf("[Repository][MatchCat][DeleteOtherMatch] failed to query, err: %s", err.Error())
		return
	}

	if rowsAffected == 0 {
		r.logger.Errorf("[Repository][MatchCat][DeleteOtherMatch] failed to query, err: no rows effected")
		return
	}

	return nil
}
