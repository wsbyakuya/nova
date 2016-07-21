package main

import (
	"testing"
)

func TestArgs(t *testing.T) {
	args := []string{"abd", "sddd", "here", "hello"}
	aa := Args(args)
	if !aa.Contains("a", "b", "c") {
		t.Log("PASS1")
	} else {
		t.Fail()
	}
	if aa.Contains("a", "b", "here") {
		t.Log("PASS2")
	} else {
		t.Fail()
	}
	test := []string{"a", "b", "here"}
	if aa.Contains(test...) {
		t.Log("PASS3")
	} else {
		t.Fail()
	}
}

func TestArgs2(t *testing.T) {
	args := []string{"abd", "sddd", "here", "hello"}
	aa := Args(args)
	tests := []struct {
		sep  []string
		want bool
	}{
		{[]string{"a", "b", "c"}, false},
		{[]string{"a", "b", "here"}, true},
		{[]string{"a", "here", "hello"}, true},
	}

	for i, v := range tests {
		if aa.Contains(v.sep...) == v.want {
			t.Log("PASS ", i)
		} else {
			t.Fail()
		}
	}
}

func BenchmarkArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsArgsContians(b)
	}
}

func IsArgsContians(b *testing.B) {
	args := []string{"abd", "sddd", "here", "hello"}
	aa := Args(args)
	tests := []struct {
		sep  []string
		want bool
	}{
		{[]string{"a", "b", "c"}, false},
		{[]string{"a", "b", "here"}, true},
		{[]string{"a", "here", "hello"}, true},
	}
	for i, v := range tests {
		if aa.Contains(v.sep...) == v.want {
			b.Log("PASS ", i)
		} else {
			b.Fail()
		}
	}
}

func BenchmarkParse(b *testing.B) {
	tests := []struct {
		sep  string
		want []byte
	}{
		{"-oe", []byte{'o', 'e'}},
		{"-abc", []byte{'a', 'b', 'c'}},
		{"-h", []byte{'h'}},
		{"-", []byte{}},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			temp := []string{}
			temp = append(temp, v.sep)
			argsMaps := Args(temp).Parse()
			if len(argsMaps) != len(v.sep)-1 {
				b.Fail()
			}
			for _, k := range v.want {
				if !argsMaps[k] {
					b.Fail()
				}
			}
		}
	}
}
