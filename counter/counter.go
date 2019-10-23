package counter

import (
	"../vender"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sum = 0.0
var report map[string]float64

func DoSingle(path string) {
	_init()
	content := vender.ReadFileAll(path);
	lines := strings.Split(content, "\n")

	for _, line := range (lines) {
		if line == "" {
			continue
		}
		items := strings.Split(line, " ")
		if items[1] == "" || items[0] == "" {
			continue;
		}
		tag := matchTag(items[0])
		cost, _ := strconv.ParseFloat(items[1], 64)
		report[tag] += cost
		sum += cost;
	}
	output()
}

func DoDir(path string) {
	outFileName := "./.merge_result.txt"
	vender.MergeFile(path, outFileName)
	DoSingle(outFileName)
	os.Remove(outFileName)
}

func matchTag(name string) string {
	for regStr, tagValue := range (TagConfig) {
		reg := regexp.MustCompile(regStr)
		res := reg.MatchString(name)
		if res {
			return TagList[tagValue]
		}
	}
	return "其他"
}

func _init() {
	report = make(map[string]float64)
	for _, v := range (TagList) {
		report[v] = 0.0
	}
}

func output() {
	fmt.Println("账单总计：")
	fmt.Println("总花费：", sum)
	fmt.Println("分类汇总：")
	for k, v := range (report) {
		fmt.Println(k, ": ", v)
	}
}
