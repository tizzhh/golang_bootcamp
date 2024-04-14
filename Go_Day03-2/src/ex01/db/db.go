package db

import (
	"fmt"
	"searchRest/types"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Store interface {
	// returns a list of items, a total number of hits and (or) an error in case of one
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
}

type Aboba int

func (a Aboba) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	var (
		places      []types.Place
		totalNumber int
	)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}

	query := fmt.Sprintf(`
		{
			"from": %d,
			"size": %d,	
		}
	`, offset, limit)
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithPretty(),
	)

	if err != nil {
		return nil, 0, fmt.Errorf("cannot complete search: %v", err)
	}

	return places, totalNumber, nil
}
