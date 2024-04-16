package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"searchRest/types"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
}

type Indx string

func (indx Indx) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
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
			"size": %d
		}`, offset, limit)
	res, err := es.Search(
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithIndex("places"),
		es.Search.WithPretty(),
		es.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, 0, fmt.Errorf("cannot complete search: %v", err)
	}

	var responseMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseMap); err != nil {
		return nil, 0, fmt.Errorf("error decoding the response body: %v", err)
	}

	hits, ok := responseMap["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("'hits' missing in api response")
	}

	totalHits, ok := hits["total"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("'total' missing in api response")
	}

	totalNum, ok := totalHits["value"].(float64)
	if !ok {
		return nil, 0, fmt.Errorf("'value' missing in api's total hits")
	}

	totalNumber = int(totalNum)

	restPlaces, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("'hits' missing in api's hits")
	}

	for _, place := range restPlaces {
		placeData, ok := place.(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid hit data format")
		}

		sourcePlaceData, ok := placeData["_source"].(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("'_source' missing in hit data")
		}

		id, ok := sourcePlaceData["id"].(float64)
		if !ok {
			return nil, 0, fmt.Errorf("'id' invalid format")
		}
		name, ok := sourcePlaceData["name"].(string)
		if !ok {
			return nil, 0, fmt.Errorf("'name' invalid format")
		}
		address, ok := sourcePlaceData["address"].(string)
		if !ok {
			return nil, 0, fmt.Errorf("'address' invalid format")
		}
		phone, ok := sourcePlaceData["phone"].(string)
		if !ok {
			return nil, 0, fmt.Errorf("'phone' invalid format")
		}
		locationData, ok := sourcePlaceData["location"].(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("'location' missing in Place data")
		}
		lat, ok := locationData["lat"].(float64)
		if !ok {
			return nil, 0, fmt.Errorf("'latitude' missing in location")
		}
		lon, ok := locationData["lon"].(float64)
		if !ok {
			return nil, 0, fmt.Errorf("'longitude' missing in location")
		}

		places = append(places, types.Place{Name: name, Address: address, Phone: phone, Id: int(id), Location: types.Location{Latitude: lat, Longitude: lon}})
	}

	return places, totalNumber, nil
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
