package report

import (
	"encoding/json"
)

//DealMeta用于提取JSON结构中的id
type DealMetaId struct {
	Meta struct {
		HasNext bool   `json:"has_next"`
		Count   int    `json:"count"`
		Version string `json:"version,omitempty"`
	} `json:"meta"`
	Objects []struct {
		Id int `json:"id"`
	} `json:"objects,omitempty"`
	Ids []int `json:"ids,omitempty"`
}

//FormatDealIds提取返回JSON中的id
func FormatDealIds(body string) string {
	deals := DealMetaId{}
	var err error
	err = json.Unmarshal([]byte(body), &deals)
	if err != nil {
		return ""
	}
	deals.Ids = []int{}
	for _, v := range deals.Objects {
		deals.Ids = append(deals.Ids, v.Id)
	}
	deals.Objects = nil

	var ss []byte
	ss, err = json.Marshal(&deals)
	if err != nil {
		return ""
	} else {
		return string(ss)
	}
}
