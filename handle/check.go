package handle

import (
	"bytes"
	"encoding/json"
)

type Items []interface{}

type UnknownJSON struct {
	Object []interface{} `json:"objects"`
}

//返回查询结果检测 16/05/16 update
func ItemsNum(body []byte) int {
	if len(body) == 0 {
		return 0
	} else if string(body) == "[]" || string(body) == "{}" {
		return 0
	}
	body = bytes.ToLower(body)
	if bytes.Contains(body, []byte("<html>")) {
		return 0
	}
	if bytes.Contains(body, []byte("objects")) {
		return objectsNum(body)
	}
	if num := itemsNum(body); num != 0 {
		return num
	}
	if len(body) > 2 {
		return 1
	}
	return 0
}

//返回结果是否为空值
func IsEmpty(body []byte) bool {
	if ItemsNum(body) != 0 {
		return false
	} else {
		return true
	}
}

func objectsNum(body []byte) int {
	objs := UnknownJSON{}
	json.Unmarshal(body, &objs)
	return len(objs.Object)
}

func itemsNum(body []byte) int {
	items := Items{}
	json.Unmarshal(body, &items)
	return len(items)
}
