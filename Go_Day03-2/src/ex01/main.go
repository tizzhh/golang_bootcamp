package main

import (
	"fmt"
	"os"
	"searchRest/db"
	"searchRest/types"
	"strings"
	"text/template"
)

const INDEXNAME string = "places"

func main() {
	// es, err := elasticsearch.NewDefaultClient()
	// if err != nil {
	// 	fmt.Errorf("error creating an Elasticsearch client: %v", err)
	// 	os.Exit(1)
	// }
	// res, err := es.Search(
	// 	es.Search.WithBody(strings.NewReader(`

	// 	`)),
	// 	es.Search.WithPretty(),
	// )
	// res, err := es.Get(
	// 	INDEXNAME,
	// 	idValue,
	// 	es.Get.WithPretty(),
	// )
}
