// testES
package main

import (
	"encoding/json"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	fmt.Println("Hello World!")
	boolQuery := elastic.NewBoolQuery()
	idents := []interface{}{"aaaa"}
	nestedQuery := elastic.NewNestedQuery("clients", elastic.NewTermsQuery("clients.develop_ident", idents...))
	boolQuery = boolQuery.Must(nestedQuery)
	tags := []interface{}{"bbb"}
	nestQuery := elastic.NewTermsQuery("tags", tags...)
	boolQuery = boolQuery.Must(nestQuery)
	src, _ := boolQuery.Source()
	jsonQuery, _ := json.Marshal(src)
	fmt.Println(string(jsonQuery))
}
