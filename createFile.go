package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"time"
)

func main() {
	fileNum := 10                //文件数量
	path := "/Users/LuJun/test/" //路径
	fileSize := 100              //文件大小单位k
	md5Ctx := md5.New()
	for i := 0; i < fileNum; i++ {
		md5Ctx.Write([]byte(string(i)))
		name := md5Ctx.Sum(nil)
		//路径
		go save(path+hex.EncodeToString(name)+".txt", fileSize)
		time.Sleep(1e9)
	}

}

func save(path string, fileSize int) {
	outputFile, outputError := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	c := "a"
	//i=307200   120k
	for i := 0; i < 1024*fileSize; i++ {
		x := rand.Intn(200)
		if x >= 48 && x <= 122 {
			c += string(x)
		}
	}
	outputFile.WriteString(c)
	defer outputFile.Close()
}
