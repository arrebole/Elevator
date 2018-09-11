/*--------------------------------------------------

			数据文件化储存

----------------------------------------------------*/
package main

type myPost struct {
	ID      int
	Content string
	Author  string
}

func main() {

	// 创建数据
	allPosts := []myPost{
		myPost{ID: 1, Content: "hello", Author: "jon"},
		myPost{ID: 2, Content: "world", Author: "Bob"},
		myPost{ID: 3, Content: "!    ", Author: "ailm"},
	}

	// 通过csv储存数据
	fileCsv(allPosts)
	// 通过gob 二进制流储存数据
	fileGod(allPosts)
}
