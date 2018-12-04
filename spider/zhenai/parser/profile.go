package parser

import (
	"shikenian/learn/spider/engine"
	"regexp"
	"shikenian/learn/spider/model"
)
const profileRege = `<div class="m-btn purple" data-v-ff544c08="">([^<]+)</div>`

func ParseProfile(content []byte) engine.ParseResult{
	reg := regexp.MustCompile(profileRege)
	matches := reg.FindAllSubmatch(content,-1)
	profile := model.Profile{}
	if matches != nil {
		for _,match := range matches {
			profile.Marriage =string(match[0])
			profile.Age = string(match[1])
			profile.Constellation = string(match[2])
			profile.Height = string(match[3])
			profile.Weight=string(match[4])
			profile.WorkingSpace=string(match[5])
			profile.Salary=string(match[6])
			profile.Job=string(match[7])
			profile.Degree=string(match[8])
 		}
	}
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result

}
