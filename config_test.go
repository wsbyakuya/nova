package main

import (
	"fmt"
	"github.com/wsbyakuya/nova/fetch"
	"testing"
)

func TestGetParas(t *testing.T) {
	fmt.Println(Api)
	list := fetch.GetFullList(Paras)
	for _, v := range list {
		fmt.Println(v)
	}
}
