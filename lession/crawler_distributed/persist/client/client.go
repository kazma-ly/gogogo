package client

import (
	"lession/crawler/engine"
	"lession/crawler_distributed/config"
	"lession/crawler_distributed/rpcsupport"
	"log"
)

// ItemSaver 保存item
func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("save: #%d, item: %+v\n", itemCount, item)
			itemCount++

			// call rpc
			var result string
			err := client.Call(config.ItemSaverRPCMethod, item, &result)
			if err != nil {
				log.Printf("Error: %v save error: %v", item, err)
				continue
			}
		}
	}()
	return out, nil
}
