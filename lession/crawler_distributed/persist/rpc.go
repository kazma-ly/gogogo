package persist

import (
	"github.com/olivere/elastic"
	"lession/crawler/engine"
	"lession/crawler/persist"
)

// ItemSaverService rpcn服务提供
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save 保存item
func (iss *ItemSaverService) Save(item engine.Item, result *string) error {
	_, err := persist.Save(item, iss.Index, iss.Client)
	if err == nil {
		*result = "ok"
	}
	return err
}
