package vender

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true, err
		}
		return false, err
	}
	return true, err
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func ReadFileByLine(file string, callback func(line string) error) {
	f, err := os.Open(file) // just pass the file name
	if err != nil {
		panic(err)
		return;
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		err := callback(line)
		if err != nil {
			fmt.Print(err)
			continue;
		}
	}
}

func ReadFileAll(file string) string {
	b, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	return str
}