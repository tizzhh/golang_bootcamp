package reader

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func GetData(idValue, indexName string) error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}
	res, err := es.Get(
		indexName,
		idValue,
		es.Get.WithPretty(),
	)
	if err != nil {
		return fmt.Errorf("error getting an index info: %v", err)
	}

	fmt.Println(res)

	return nil
}

func GetInfo(value, indexName string) error {
	switch value {
	case "help":
		fmt.Println("Usage: ./loadElastic [help/info/id]")
	case "info":
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			return fmt.Errorf("error creating an Elasticsearch client: %v", err)
		}
		res, err := es.Indices.Get(
			[]string{indexName},
			es.Indices.Get.WithPretty(),
		)
		if err != nil {
			return fmt.Errorf("error getting an index info: %v", err)
		}
		fmt.Println(res)
	}

	return nil
}
