package main

import (
	"github.com/olivere/elastic"
	"lession/crawler_distributed/config"
	"lession/crawler_distributed/persist"
	"lession/crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(rpcServer(config.ItemSaverRPCAddr, config.ItemSaverRPCAddr))
}

func rpcServer(host, index string) error {
	client, err := elastic.NewSimpleClient(
		elastic.SetURL("http://127.0.0.1:32780"), // docker random port
		elastic.SetSniff(false),                  // must turn off sniff
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
