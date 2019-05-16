package model

import "encoding/json"

type Profile struct {
	Url        string
	Id         string
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	IncomeLow  int
	IncomeUp   int
	Marriage   string // 婚姻状况
	Education  string // 教育
	Occupation string
	Job        string // 职业
	Xinzuo     string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, nil
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
