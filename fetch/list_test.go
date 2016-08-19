package fetch

import (
	"testing"
)

var paras [][]string

func init() {
	paras = append(paras, append([]string{}, "user_type", "0", "1"))
	paras = append(paras, append([]string{}, "user_role", "0", "1", "4", "6"))
	paras = append(paras, append([]string{}, "page", "1", "2", "3", "4", "5", "6", "7"))
}

func TestFullList(t *testing.T) {
	list := GetFullList(paras)
	for _, v := range list {
		t.Log(v)
	}
}

func TestRandomList(t *testing.T) {
	list := GetRandomList(paras)
	for _, v := range list {
		t.Log(v)
	}
}
