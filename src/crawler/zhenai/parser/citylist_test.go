package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470
	exceptedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	exceptedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if resultSize != len(result.Requests) {
		t.Errorf("result  should have %d,but had %d", resultSize, len(result.Requests))
	}

	for i, url := range exceptedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("excepted url #%d %s,but url is %s",
				i, url, result.Requests[i].Url)
		}
	}
	if resultSize != len(result.Items) {
		t.Errorf("result  should have %d,but had %d", resultSize, len(result.Items))
	}
	for i, city := range exceptedCities {
		if city != result.Items[i].(string) {
			t.Errorf("excepted city #%d %s,but was is %s",
				i, city, result.Items[i].(string))
		}
	}
}
