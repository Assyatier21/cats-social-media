package elasticsearch

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/backendmagang/project-1/models/entity"
	"github.com/olivere/elastic/v7"
)

func (r *elasticRepository) GetCategoryTree(ctx context.Context, req entity.GetCategoriesRequest) ([]entity.Category, error) {
	var (
		categories = []entity.Category{}
	)

	res, err := r.es.Search().Index(r.cfg.IndexCategory).From(req.Offset).Size(req.Limit).Do(ctx)
	if err != nil {
		return categories, err
	}

	if res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var category entity.Category
			err = json.Unmarshal(hit.Source, &category)
			if err != nil {
				log.Println("[Repository][Elastic][GetCategoryDetails] failed to unmarshal category, err: ", err)
				return categories, err
			}
			categories = append(categories, category)
		}
	}

	return categories, err
}

func (r *elasticRepository) GetCategoryDetails(ctx context.Context, query elastic.Query) (entity.Category, error) {
	var (
		category = entity.Category{}
	)

	res, err := r.es.Search().Index(r.cfg.IndexCategory).Query(query).Do(ctx)
	if err != nil {
		return category, err
	}

	if res.Hits.TotalHits.Value > 0 {
		err = json.Unmarshal(res.Hits.Hits[0].Source, &category)
		if err != nil {
			log.Println("[Repository][Elastic][GetCategoryDetails] failed to unmarshal category, err: ", err)
			return category, err
		}
	}

	return category, err
}

func (r *elasticRepository) InsertCategory(ctx context.Context, category entity.InsertCategoryRequest) error {
	categoryJSON, err := json.Marshal(category)
	if err != nil {
		log.Println("[Repository][Elastic][InsertCategory] failed to marshal category, err: ", err)
		return err
	}

	_, err = r.es.Index().Index(r.cfg.IndexCategory).Id(strconv.Itoa(category.ID)).BodyJson(string(categoryJSON)).Do(ctx)
	if err != nil {
		log.Println("[Repository][Elastic][InsertCategory] failed to insert category, err: ", err)
		return err
	}

	return nil
}

func (r *elasticRepository) UpdateCategory(ctx context.Context, category entity.UpdateCategoryRequest) error {
	_, err := r.es.Update().Index(r.cfg.IndexCategory).Id(strconv.Itoa(category.ID)).Doc(category).Do(ctx)
	if err != nil {
		log.Println("[Repository][Elastic][UpdateCategory] failed to update category, err: ", err)
		return err
	}

	return nil
}

func (r *elasticRepository) DeleteCategory(ctx context.Context, req entity.DeleteCategoryRequest) error {
	_, err := r.es.Delete().Index(r.cfg.IndexCategory).Id(strconv.Itoa(req.ID)).Do(ctx)
	if err != nil {
		log.Println("[Repository][Elastic][DeleteCategory] failed to delete category, err: ", err)
		return err
	}

	return nil
}
