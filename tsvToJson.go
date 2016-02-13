package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/**
 * 第1引数：入力ファイル
 * 第2引数：出力ファイル
 * 第3引数以降：Json定義（[例]0:key 5:title）
 */
func main() {

	var srcFile *os.File
	if len(os.Args) < 3 {
		panic("Args error")
	} else {
		var err error
		srcFile, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer srcFile.Close()
	}
	// reader
	reader := csv.NewReader(srcFile)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	var jsondefs []string
	for i := 3; i < len(os.Args); i++ {
		jsondefs = append(jsondefs, os.Args[i])
	}

	var content []byte
	var recordCount int32
	content = append(content, "[\n"...)
	recordCount = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if recordCount != 0 {
			content = append(content, ",\n"...)
		}
		content = append(content, "{"...)
		for i := 0; i < len(jsondefs); i++ {
			var key string
			var index int
			key = strings.Split(jsondefs[i], ":")[1]
			index, _ = strconv.Atoi(strings.Split(jsondefs[i], ":")[0])
			content = append(content, "\""+key+"\":\""+record[index]+"\""...)
			if i < len(jsondefs)-1 {
				content = append(content, ", "...)
			}
		}
		content = append(content, "}"...)
		recordCount++
	}
	content = append(content, "\n]"...)
	ioutil.WriteFile(os.Args[2], content, os.ModePerm)
}
