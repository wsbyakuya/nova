package fetch

import (
	"sort"
	"strings"
)

func GetFullList(paras map[string][]string) []string {
	list := []string{}
	point := make(map[int]int)
	length := len(paras)
	lengths := make(map[int]int)

	keys := make([]string, 0, len(paras))
	for k, _ := range paras {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		lengths[i] = len(paras[k])
		point[i] = 0
	}
	p := 0
	for p < length {
		//combine
		str := []string{}
		for k, _ := range keys {
			str = append(str, keys[k]+"="+paras[keys[k]][point[k]])
		}
		list = append(list, strings.Join(str, "&"))

		//pointer control
		for p > 0 && point[p-1] == 0 {
			p--
		}
		point[p]++
		for point[p] == lengths[p] {
			point[p] = 0
			p++
			point[p]++
		}
	}
	return list
}
