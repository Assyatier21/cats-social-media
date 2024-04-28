package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/backend-magang/cats-social-media/models/entity"
)

func (r *repository) GetListCat(ctx context.Context, req entity.GetListCatRequest) (result []entity.Cat, err error) {
	query, args := buildQueryGetListCats(req)
	query = r.db.Rebind(query)

	err = r.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		log.Println("[Repository][Cat][GetList] failed to query, err: ", err.Error())
		err = fmt.Errorf("failed to query: %s", err.Error())
		return
	}

	return
}