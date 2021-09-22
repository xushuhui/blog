package main

import (
	"bufio"

	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(url string) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Referer", "http://www.imooc.com/")

	resp, err := (&http.Client{}).Do(req)

	s := strings.Split(url, "/")
	name := s[len(s)-1]

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./img/"+name, data, 0644)
	if err != nil {
		panic(err)
	}
	//img := uploadSinaImg(name)
	return
}
func main() {
	dir := "./java/"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.OpenFile(dir+"/"+file.Name(), os.O_RDWR|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}
		output, needHandle, err := readImg(f)
		if needHandle {
			err = writeToFile(dir+"/"+file.Name(), output)
			if err != nil {
				panic(err)
			}
		}
	}

}
func readImg(f *os.File) ([]byte, bool, error) {
	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)

	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				//str1 := []byte("### 微信公众号\n")
				//str2 := []byte("\n![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)")
				//output = append(output, str1...)
				//output = append(output, str2...)

				return output, needHandle, nil
			}
			return nil, needHandle, err
		}
		if strings.Contains(string(line), "img.mukewang.com") {
			s := string(line)

			var index int
			for k, v := range s {
				if v == 40 {
					index = k
				}
			}
			old := s[index+1 : len(s)-1]

			Download(old)
			newByte := []byte(s)
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)

		}
	}
}
func readFile(f *os.File, fname string) ([]byte, bool, error) {

	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)
	var n int
	for {
		n++

		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				//str1 := []byte("### 微信公众号\n")
				//str2 := []byte("\n![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)")
				//output = append(output, str1...)
				//output = append(output, str2...)

				return output, needHandle, nil
			}
			return nil, needHandle, err
		}
		if n == 1 {
			newByte := []byte(
				`---
title: ` + fname + `
zhihu-title-image: https://pica.zhimg.com/v2-e1cf667c04b0f63c15003183ddd03e79_1440w.jpg?source=172ae18b				
---`)
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)

			output = append(output, line...)
			output = append(output, []byte("\n")...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}

	}
}
func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
