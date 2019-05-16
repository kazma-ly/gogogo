package persist

import (
	"context"
	"errors"
	"lession/crawler/engine"
	"log"

	"github.com/olivere/elastic"
)

// ItemSaver 保存item
func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	client, err := elastic.NewSimpleClient(
		elastic.SetURL("http://127.0.0.1:32771"), // docker random port
		elastic.SetSniff(false),                  // must turn off sniff
	)

	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("save: #%d, item: %+v\n", itemCount, item)
			itemCount++

			_, err := Save(item, index, client)
			if err != nil {
				log.Printf("Error: %v save error: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}

// Save 保存item
func Save(item engine.Item, index string, client *elastic.Client) (id string, err error) {
	if item.Type == "" {
		return "", errors.New("must supply type")
	}

	service := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		service.Id(item.Id)
	}

	response, err := service.Do(context.Background())
	if err != nil {
		return "", nil
	}
	return response.Id, nil
}
