package models

import "strings"

func GetTagsMap(tags []string) map[string]int{
	var tagsMap = make(map[string]int)
	for _,tag := range tags{
		tagLine := strings.Split(tag,"&")
		for _,val := range tagLine{
			tagsMap[val]++
		}
	}
	return tagsMap
}