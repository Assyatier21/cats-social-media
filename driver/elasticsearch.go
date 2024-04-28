package driver

import (
	"context"
	"log"

	"github.com/backendmagang/project-1/config"
	"github.com/olivere/elastic/v7"
)

func InitElasticClient(cfg config.ElasticConfig) *elastic.Client {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(cfg.Address),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	log.Println("[Elasticsearch] initialized...")
	if err != nil {
		log.Println("[Elasticsearch] failed to connect to elasticsearch: ", err)
		return nil
	}

	info, _, err := client.Ping(cfg.Address).Do(ctx)
	if err != nil {
		log.Println("[Elasticsearch] error ping, err: ", err)
	}
	log.Printf("[Elasticsearch] successfully connected. running version %s\n", info.Version.Number)
	return client
}
