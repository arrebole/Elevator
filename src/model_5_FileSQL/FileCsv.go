package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func fileCsv(data []myPost) {
	// 创建一个文件 并返回其文件句柄
	csvFile, err := os.Create("./data/posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// 写入数据
	writer := csv.NewWriter(csvFile)
	for _, post := range data {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() //清理缓存

	/*---------------------------------------*/

	// 打开文件
	file, err := os.Open("./data/posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取数据
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 处理后输出
	var posts []myPost
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := myPost{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println("Csv: >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	for _, item := range posts {
		fmt.Printf("Id: %d  Content: %s  Author: %s\n", item.ID, item.Content, item.Author)
	}
}
