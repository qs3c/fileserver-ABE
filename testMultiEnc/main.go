package main

import (
	"bufio"
	"chunkAES/aes/aesutil"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	key := []byte("96B6715EF50FA455")

	file, err := os.Open("C:\\Users\\27499\\Desktop\\ys-2022-06-12 22-25-44.mp4")
	if err != nil {
		fmt.Println("打开文件出错！", err.Error())
	}
	defer file.Close()

	encfile, err := os.Create("C:\\Users\\27499\\Desktop\\RFABEO\\multi_encfile_mp4_2")
	if err != nil {
		fmt.Println("创建文件出错！", err.Error())
	}
	defer encfile.Close()

	r := bufio.NewReader(file)
	//var dataBlocks []byte
	//dataBlocks := make([]byte, 52428800) //50MB
	//buf := make([]byte, 102400)          // 每次读取字节数 100KB

	//buf := make([]byte, 16777216) //16MB
	//slice优化:初始化16MB+16B空间
	buf := make([]byte, 16777232) //16MB+16B

	for {
		//slice优化
		//r.Read(buf[:16777217])
		//n, err := r.Read(buf) // 读取字节数 n
		//但是只让写入buf 16MB
		n, err := r.Read(buf[:16777216])
		fmt.Println("读入了：", n)

		if err != nil {
			if err == io.EOF {
				// 判断文件读取结束
				break
			}
			log.Printf("打开文件失败:%#v", err.Error())
		}
		//dataBlocks = append(dataBlocks, buf[:n]...) // 注意有人这里[:n] 是读的字节数赋值，最后一次读取可能小于buf定义量
		//把每次读到的 buf[:n] 进行标记加密，方便后续解密 unpad
		//但是 Go是值传递，这种优化不一定有用，好吧经过查证是有用的，两个slice在不同地址，但是指向的数组是同一个数组！
		encbuf, err := aesutil.MarkEncryptAES(buf[:n], key)
		if err != nil {
			fmt.Println("加密出错", err.Error())
		}
		//加密后写入文件
		_, err = encfile.Write(encbuf)
		if err != nil {
			fmt.Println("写入出现错误", err.Error())
		}
	}
}
