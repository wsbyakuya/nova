package fetch

import (
	"testing"
)

func TestFullList(t *testing.T) {
	paras := make(map[string][]string)
	paras["user_type"] = []string{"0", "1"}
	paras["user_role"] = []string{"0", "1", "4", "6"}
	paras["page"] = []string{"1", "2", "3", "4", "5", "6", "7"}
	list := GetFullList(paras)
	for _, v := range list {
		t.Log(v)
	}
}

func TestRandomList(t *testing.T) {
	paras := make(map[string][]string)
	paras["user_type"] = []string{"0", "1"}
	paras["user_role"] = []string{"0", "1", "4", "6"}
	paras["page"] = []string{"1", "2", "3", "4", "5", "6", "7"}
	list := GetRandomList(paras)
	for _, v := range list {
		t.Log(v)
	}
}
