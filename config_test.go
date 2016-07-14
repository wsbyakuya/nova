package main

import (
	"fmt"
	"github.com/wsbyakuya/nova/fetch"
	"testing"
)

func TestGetParas(t *testing.T) {
	loadParas()
	fmt.Println(Api)
	list := fetch.GetFullList(Paras)
	for _, v := range list {
		fmt.Println(v)
	}
}
