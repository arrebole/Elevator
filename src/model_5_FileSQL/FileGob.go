/*
	将数据二进制形式保存成文件
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

func fileGod(data []myPost) {
	store(data, "./data/PostGob")
	var postRead []myPost
	load(&postRead, "./data/PostGob")

	fmt.Println("FileGob: >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	for _, item := range postRead {
		fmt.Printf("Id: %d  Content: %s  Author: %s\n", item.ID, item.Content, item.Author)
	}
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
