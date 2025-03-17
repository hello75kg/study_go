package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.0.249:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	q := elastic.NewMatchQuery("address", "street")
	src, err := q.Source()
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	got := string(data)
	fmt.Println(got)
	// -------------------
	elastic.NewMatchQuery("address", "street")
	result, err := client.Search().Index("user").Query(q).Do(context.Background())
	if err != nil {
		panic(err)
	}
	total := result.Hits.TotalHits.Value
	fmt.Println(total)

	for _, value := range result.Hits.Hits {
		if jsonData, err := value.Source.MarshalJSON(); err == nil {
			fmt.Println(string(jsonData))
		} else {
			fmt.Println(err)
		}
	}
}
