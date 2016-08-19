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
	p := length - 1
	for p >= 0 {
		//combine
		str := []string{}
		for k, _ := range keys {
			str = append(str, keys[k]+"="+paras[keys[k]][point[k]])
		}
		list = append(list, strings.Join(str, "&"))

		//pointer control
		for p < length-1 && point[p+1] == 0 {
			p++
		}
		point[p]++
		for point[p] == lengths[p] {
			point[p] = 0
			p--
			point[p]++
		}
	}
	return list
}

func GetRandomList(paras map[string][]string) []string {
	list := GetFullList(paras)
	length := len(list)
	if length <= 3 {
		return list
	} else {
		mid := length / 2
		nlist := make([]string, 0, 3)
		nlist = append(nlist, list[0], list[mid], list[length-1])
		return nlist
	}
}
