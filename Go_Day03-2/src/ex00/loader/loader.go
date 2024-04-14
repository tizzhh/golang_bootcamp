package loader

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

const (
	ID = iota
	NAME
	ADDRESS
	PHONE
	LONGITUDE
	LATITUDE
)

func CreateIndex(indexName, mapping string) error {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}

	_, err = es.Indices.Delete(
		[]string{indexName},
	)
	if err != nil {
		return fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}

	createIndexReq, err := es.Indices.Create(
		indexName,
		es.Indices.Create.WithBody(strings.NewReader(mapping)),
	)

	if err != nil {
		return fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}
	if createIndexReq.IsError() {
		return fmt.Errorf("error creating an Elasticsearch client: %v", createIndexReq)
	}

	fmt.Println(createIndexReq)

	return nil
}

type Restaurant struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location GeoPoint `json:"location"`
}

type GeoPoint struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func LoadDataInDb(csv_path, indexName, mapping string) error {
	var buf bytes.Buffer
	file, err := os.Open(csv_path)
	if err != nil {
		return fmt.Errorf("error opening csv file: %v", err)
	}
	defer file.Close()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return fmt.Errorf("error creating an Elasticsearch client: %v", err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // missing ID in header fix
	reader.LazyQuotes = true    // " error
	reader.Comma = '\t'
	reader.Read() // header line
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error opening csv file: %v", err)

		}

		var rest Restaurant
		var location GeoPoint

		for i, value := range line {
			switch i {
			case ID:
				rest.Id, _ = strconv.Atoi(value)
			case NAME:
				rest.Name = value
			case ADDRESS:
				rest.Address = value
			case PHONE:
				rest.Phone = value
			case LONGITUDE:
				float_val, _ := strconv.ParseFloat(value, 64)
				location.Longitude = float_val
			case LATITUDE:
				float_val, _ := strconv.ParseFloat(value, 64)
				location.Latitude = float_val
			}
		}
		rest.Location = location

		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, rest.Id, "\n"))

		data, err := json.Marshal(rest)
		if err != nil {
			return fmt.Errorf("error marshalling csv file: %v", err)
		}

		data = append(data, "\n"...)
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}

	res, err := es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
	if err != nil {
		return fmt.Errorf("failure indexing: %v", err)
	}

	if res.IsError() {
		return fmt.Errorf("failure indexing: %v", err)
	}

	return nil
}
