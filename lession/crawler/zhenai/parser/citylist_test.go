package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d, got %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d, got %d", resultSize, len(result.Items))
	}
}
