package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"searchRest/types"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Total struct {
	Value int `json:"value"`
}

type Hit struct {
	Source types.Place `json:"_source"`
}

type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type OuterHits struct {
	OuterHits Hits `json:"hits"`
}

type Store interface {
	GetPlaces(limit int, offset int, mode string) ([]types.Place, int, error)
}

type Indx string

const (
	MODE_PAGE = iota
	MODE_CLOSEST
)

func (indx Indx) GetPlaces(limit, offset int, lat, lon float64, mode int) ([]types.Place, int, error) {
	var places []types.Place

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, 0, fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}

	var query string
	switch mode {
	case MODE_PAGE:
		query = fmt.Sprintf(`
			{
				"from": %d,
				"size": %d
			}`, offset, limit)
	case MODE_CLOSEST:
		query = fmt.Sprintf(`{
			"sort": [
				{
				"_geo_distance": {
					"location": {
					"lat": %f,
					"lon": %f
					},
					"order": "asc",
					"unit": "km",
					"mode": "min",
					"distance_type": "arc",
					"ignore_unmapped": true
				}
				}
			],
			"size": %d
		  }`, lat, lon, limit)
	}
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithIndex(string(indx)),
		es.Search.WithPretty(),
		es.Search.WithTrackTotalHits(true),
	)
	defer res.Body.Close()

	if err != nil {
		return nil, 0, fmt.Errorf("cannot complete search: %v", err)
	}

	var hitsReponse OuterHits
	if err := json.NewDecoder(res.Body).Decode(&hitsReponse); err != nil {
		return nil, 0, fmt.Errorf("error decoding the response body: %v", err)
	}

	for _, place := range hitsReponse.OuterHits.Hits {
		places = append(places, place.Source)
	}

	return places, hitsReponse.OuterHits.Total.Value, nil
}

func IncreaseMaxEntries() error {
	url := "http://localhost:9200/places/_settings"
	payload := `{
		"index": {
			"max_result_window": 20000
		}
	}`

	req, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	return nil
}
