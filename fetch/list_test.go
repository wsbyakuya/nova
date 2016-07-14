package fetch

import (
	"encoding/json"
	"os"
	"testing"
)

var (
	ConstParas []string
	API        ApiInfo
)

type ApiInfo struct {
	Address string      `json:"api"`
	Paras   []Parameter `json:"parameters"`
}

func TestList(t *testing.T) {
	loadParas()
	t.Log(GetFullList(API.Paras))
}

func loadParas() {
	f, err := os.Open("api.json")
	if err != nil {
		panic("无法打开参数文件")
	}
	defer f.Close()
	temp := ApiInfo{Paras: []Parameter{}}
	d := json.NewDecoder(f)
	err = d.Decode(&temp)
	if err != nil {
		panic("JSON文件解析失败")
	}
	API.Address = temp.Address
	for _, v := range temp.Paras {
		l := len(v.Values)
		if l == 1 {
			ConstParas = append(ConstParas, v.Para+"="+v.Values[0])
		} else if l > 1 {
			API.Paras = append(API.Paras, v)
		}
	}
}
