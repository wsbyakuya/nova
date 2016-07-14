package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Host1   string
	Host2   string
	Hosts   []string
	Port    string
	MaxDiff float64
	Timeout int
)

var (
	ConstParas []string
	Api        string
	Paras      map[string][]string
)

func init() {
	loadConfig("")
	loadParas("")
}

func loadConfig(filename string) {
	filename = GlobalPath + "env.cfg"
	f, err := os.Open(filename)
	if err != nil {
		panic("无法打开配置文件")
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		line = strings.Trim(line, "\r\n")
		if line != "" {
			if line[0] == '#' {
				continue
			}
			k, v := parseKeyValue(line)
			switch k {
			case "host1":
				Host1 = v
			case "host2":
				Host2 = v
			case "hosts":
				Hosts = SplitAndTrim(v)
			case "port":
				Port = v
			case "max_diff":
				MaxDiff, _ = strconv.ParseFloat(v, 64)
			case "timeout":
				Timeout, _ = strconv.Atoi(v)
			default:
			}
		}
		if err == io.EOF {
			break
		}
	}
}

func loadParas(filename string) {
	filename = GlobalPath + "api.cfg"
	Paras = make(map[string][]string)
	f, err := os.Open(filename)
	if err != nil {
		panic("无法打开参数文件")
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		line = strings.Trim(line, "\r\n")
		if line != "" {
			if line[0] == '#' {
				continue
			}
			k, v := parseKeyValue(line)
			if k == "api" {
				Api = v
				continue
			}
			//将固定参数分离，提高效率
			vals := SplitAndTrim(v)
			if len(vals) > 1 {
				Paras[k] = vals
			} else if len(vals) == 1 {
				ConstParas = append(ConstParas, k+"="+v)
			}

		}
		if err == io.EOF {
			break
		}
	}
}

func SplitAndTrim(line string) []string {
	ss := strings.Split(line, ",")
	for i := range ss {
		ss[i] = strings.Trim(ss[i], " ")
	}
	return ss
}

func parseKeyValue(line string) (key, value string) {
	line = strings.Trim(line, "\r\n")
	strs := strings.Split(line, "=")

	k := strings.Trim(strs[0], " ")
	v := strings.Trim(strs[1], " ")
	return k, v
}
