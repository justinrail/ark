package service

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

//GetAllLog get all logs from ark.log
func GetAllLog() []string {

	content := make([]string, 0)

	list, err := filepath.Glob("logs/ark.log.*")
	if err == nil {
		for _, v := range list {
			//b, err2 := ioutil.ReadFile(v) // just pass the file name
			f, err2 := os.OpenFile(v, os.O_RDONLY, os.ModePerm)
			if err2 == nil {
				br := bufio.NewReader(f)
				for {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						break
					}
					content = append(content, string(a))
				}
			}
		}
	}

	return content
}
