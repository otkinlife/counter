package vender

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func MergeFile(rootPath string, outFileName string) {
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if openErr != nil {
		fmt.Printf("Can not open file %s", outFileName)
	}
	bWriter := bufio.NewWriter(outFile)

	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		//这里是文件过滤器，表示我仅仅处理txt文件
		if strings.HasSuffix(path, ".txt") {
			fp, fpOpenErr := os.Open(path)
			if fpOpenErr != nil {
				fmt.Printf("Can not open file %v", fpOpenErr)
				return fpOpenErr
			}
			bReader := bufio.NewReader(fp)
			for {
				buffer := make([]byte, 1024)
				readCount, readErr := bReader.Read(buffer)
				if readErr == io.EOF {
					bWriter.Write([]byte("\n"))
					break
				} else {
					bWriter.Write(buffer[:readCount])
				}
			}
		}
		return err
	})
	bWriter.Flush()
}