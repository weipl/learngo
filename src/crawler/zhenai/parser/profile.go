package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)
//1、()表示我们需要的 2、[\d] + 表示匹配年龄	3、[^<]+ 匹配字符串
var ageRe = regexp.MustCompile( `<div class="m-btn purple" data-v-bff6f798="">([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([\d]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">月收入:([^<]+)千</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name
	age ,err := strconv.Atoi(
		extreactString(contents,ageRe))
	if err != nil{
		profile.Age = age
	}
		height ,err := strconv.Atoi(
		extreactString(contents,heightRe))
	if err != nil{
		profile.Height = height
	}
		weight ,err := strconv.Atoi(
		extreactString(contents,weightRe))
	if err != nil{
		profile.Weight = weight
	}

	profile.Marriage = extreactString(contents,marriageRe)
	profile.Xinzuo = extreactString(contents,xinzuoRe)
	profile.Income = extreactString(contents,incomeRe)
	profile.Education = extreactString(contents,educationRe)
	profile.Hokou = extreactString(contents,hokouRe)
	profile.Occupation = extreactString(contents,occupationRe)
	profile.House = extreactString(contents,houseRe)
	profile.Car = extreactString(contents,carRe)


	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return  result
}
func extreactString(contents []byte,re *regexp.Regexp)string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2{
		return string(match[1])
		}else {
			return ""
		}

}
