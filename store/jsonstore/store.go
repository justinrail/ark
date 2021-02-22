package jsonstore

import (
	"ark/util/cfg"
	"ark/util/exe"
	"ark/util/exp"
	"bufio"
	"bytes"
	"container/list"
	"encoding/json"
	"os"
	"path"
	"strings"
)

var defaultDir = "."

func init() {
	defaultDir = path.Join(exe.Info().AppPath, cfg.Read().App.JSONDBPath)
}

//AppendRecord add record to table
func AppendRecord(tableName string, record interface{}) bool {
	filepath := path.Join(defaultDir, tableName+".txt")

	str, err2 := json.Marshal(record)
	if err2 == nil {
		appendLineToFile(filepath, str)
	} else {
		return false
	}

	return true
}

//AppendRecords 多条记录一次写库
func AppendRecords(tableName string, records *list.List) bool {
	filepath := path.Join(defaultDir, tableName+".txt")
	buffer := bytes.NewBufferString("")

	for e := records.Front(); e != nil; e = e.Next() {
		str, err2 := json.Marshal(e.Value)
		if err2 == nil {
			buffer.Write(str)
		}
	}

	appendLineToFile(filepath, buffer.Bytes())
	return true
}

func appendLineToFile(path string, content []byte) error {
	limitFile(path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	defer f.Close()

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	_, _ = f.WriteString("\n")

	if err != nil {
		return err
	}
	return nil
}

//limitFile 限制文件大小1M
func limitFile(path string) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	defer file.Close()
	if err == nil {
		fi, err2 := file.Stat()
		exp.CheckError(err2)
		size := fi.Size()
		if size > 1024*1024 {
			err3 := file.Truncate(0)
			exp.CheckError(err3)
		}
	}
}

//RemoveTable RemoveTable
func RemoveTable(tableName string) {
}

//ReadAllRecords ReadAllRecords from x table
func ReadAllRecords(tableName string) []string {
	items := make([]string, 0)
	filepath := path.Join(defaultDir, tableName+".txt")

	//scanner 不支持65535以上的行数，如果有这样的情况会报错
	f, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return items
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line != "\n" {
			line = strings.Replace(line, "\n", "", -1)
			items = append(items, line)
		}
	}
	if err := sc.Err(); err != nil {
		return items
	}

	return items
}
