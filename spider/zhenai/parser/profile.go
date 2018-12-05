package parser

import (
	"regexp"
	"shikenian/learn/spider/engine"
	"shikenian/learn/spider/model"
)

const profileRege = `<div class="m-btn purple" data-v-ff544c08>([^<]+)</div>`
const nameRege = `<h1 class="nickName" data-v-312fdcc4>([^<]+)</h1>`
var reg = regexp.MustCompile(profileRege)
var regName = regexp.MustCompile(nameRege)
func ParseProfile(content []byte) engine.ParseResult {
	matches := reg.FindAllSubmatch(content, -1)
	var value []string
	for _, match := range matches {
		value = append(value, string(match[1]))
	}
	if len(value) != 9 {
		return engine.ParseResult{}
	}
	profile := model.Profile{}
	profile.Marriage = value[0]
	profile.Age = value[1]
	profile.Constellation = value[2]
	profile.Height =  value[3]
	profile.Weight =  value[4]
	profile.WorkingSpace =  value[5]
	profile.Salary =  value[6]
	profile.Job =  value[7]
	profile.Degree =  value[8]

	name := regName.FindSubmatch(content)
	profile.Name = string(name[1])


	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
