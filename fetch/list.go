package fetch

import (
	"strings"
)

func GetFullList(paras [][]string) []string {
	list := []string{}
	point := make(map[int]int)
	length := len(paras)
	lengths := make(map[int]int)

	keys := make([]string, 0, len(paras))
	for i := range paras {
		if len(paras[i]) > 0 {
			keys = append(keys, paras[i][0])
			paras[i] = paras[i][1:]
		}
	}

	for i := range keys {
		lengths[i] = len(paras[i])
		point[i] = 0
	}
	p := length - 1
	for p >= 0 {
		//combine
		str := []string{}
		for i := range keys {
			str = append(str, keys[i]+"="+paras[i][point[i]])
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

func GetRandomList(paras [][]string) []string {
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
