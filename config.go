package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	Host1   string
	Host2   string
	Port    string
	MaxDiff float64
	Timeout int
)

var (
	ConstParas []string
	Api        string
	Method     string
	Paras      [][]string
	Cookies    map[string]string
	Header     map[string]string
)

func init() {
	Paras = make([][]string, 0)
	Cookies = make(map[string]string)
	Header = make(map[string]string)
}

func loadConfig(filename string) {
	var f *os.File
	var err error
	f, err = os.Open(filename)
	if err != nil {
		//在当前目录未找到环境配置文件，将搜索上级目录中配置文件
		if os.IsNotExist(err) {
			if parentDir := getParentDir(); parentDir != "" {
				f, err = os.Open(parentDir + "\\" + filename)
				if err != nil {
					FailAndExit(err)
				}
			} else {
				FailAndExit(err)
			}
		} else {
			FailAndExit(err)
		}
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

func loadConfigFile(filename string) {
	tag := "api"
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		} else {
			FailAndExit(err)
		}
	}
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)
	ls := strings.Split(string(buf), "\n")
	for _, v := range ls {
		w := strings.Trim(v, "\r\n")
		w = strings.Trim(w, " ")
		if len(w) > 0 && w[0] != '#' {
			if strings.HasPrefix(w, "[") && strings.HasSuffix(w, "]") {
				tag = strings.ToLower(w[1 : len(w)-1])
			} else {
				setParameters(tag, w)
			}
		}
	}
}

func setParameters(tag, line string) {
	key, value := parseKeyValue(line)
	switch tag {
	case "api":
		if key == "api" {
			Api = value
			return
		} else if key == "method" {
			Method = strings.ToUpper(value)
			return
		}
		vls := SplitAndTrim(value)
		if len(vls) > 1 {
			kvs := []string{key}
			kvs = append(kvs, vls...)
			Paras = append(Paras, kvs)
		} else {
			ConstParas = append(ConstParas, key+"="+value)
		}
	case "cookies", "cookie":
		Cookies[key] = value
	case "header", "headers":
		Header[key] = value
	}
}

func SplitAndTrim(line string) []string {
	var ss []string
	if strings.Contains(line, ";") {
		ss = strings.Split(line, ";")
		for i := range ss {
			slice := strings.Split(ss[i], ",")
			for j := range slice {
				slice[j] = strings.Trim(slice[j], " ")
			}
			ss[i] = strings.Join(slice, ",")
		}
	} else {
		ss = strings.Split(line, ",")
		for i := range ss {
			ss[i] = strings.Trim(ss[i], " ")
		}
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

func getParentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	runes := []rune(currentDir)
	l := strings.LastIndex(currentDir, "\\")
	return string(runes[0:l])
}
