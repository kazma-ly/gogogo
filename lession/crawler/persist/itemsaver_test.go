package persist

import (
	"context"
	"encoding/json"
	"lession/crawler/engine"
	"lession/crawler/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSaver(t *testing.T) {
	profile := engine.Item{
		Url:  "http://123.com",
		Type: "zhenai",
		Id:   "1121223",
		Payload: model.Profile{
			Name:       "斌豪09",
			Age:        100,
			Gender:     "女",
			Height:     170,
			Weight:     120,
			IncomeLow:  8000,
			IncomeUp:   9000,
			Marriage:   "已婚",
			Education:  "本科",
			Occupation: "",
			Job:        "职业",
			Xinzuo:     "不知道",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	id, err := Save(profile, "xq_test", client)

	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("xq_test").Type("zhenai").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	var item engine.Item
	err = json.Unmarshal([]byte(*result.Source), &item)
	if err != nil {
		panic(err)
	}
	actual, _ := model.FromJsonObj(item.Payload)
	item.Payload = actual

	t.Log(item)
	if item != profile {
		t.Errorf("got %v; but wanna %v", item, profile)
	}
}
