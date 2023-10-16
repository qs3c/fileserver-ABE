package aesutil

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"os"
)

// IsExists 判断所给路径文件/文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// IfNoFileToCreate 文件不存在就创建文件
func IfNoFileToCreate(fileName string) (file *os.File) {
	var f *os.File
	var err error
	if !IsExists(fileName) {
		f, err = os.Create(fileName)
		if err != nil {
			return
		}
		log.Printf("IfNoFileToCreate 函数成功创建文件:%s", fileName)
		defer f.Close()
	}
	return f
}

func WriteStringToFile(fileName string, writeInfo string) {
	//先创建 fileName应该为：文件名+enc+上传者用户名
	_ = IfNoFileToCreate(fileName)

	f, err := os.OpenFile(fileName, os.O_APPEND, 0777) //打开文件
	defer f.Close()
	if err != nil {
		log.Printf("打开文件失败:%+v", err)
		return
	}
	// 创建 Writer 对象
	w := bufio.NewWriter(f)
	// 写入文件
	if _, err = w.WriteString(writeInfo); err != nil {
		log.Printf("WriteStringToFileMethod4 写入文件失败:%+v", err)
		return
	}
	w.Flush()
	log.Printf("WriteStringToFile 写入文件成功")

}

func ReadFileToString(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("打开文件失败:%#v", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var dataBlocks []byte
	buf := make([]byte, 102400) // 每次读取字节数
	for {
		n, err := r.Read(buf) // 读取字节数 n
		if err != nil {
			if err == io.EOF {
				// 判断文件读取结束
				break
			}
			log.Printf("打开文件失败:%#v", err)
		}
		dataBlocks = append(dataBlocks, buf[:n]...) // 注意有人这里[:n] 是读的字节数赋值，最后一次读取可能小于buf定义量
	}
	return string(dataBlocks)
}

func ReadFileToByteStream(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("打开文件失败:", fileName)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var dataBlocks []byte
	buf := make([]byte, 102400) // 每次读取字节数
	for {
		n, err := r.Read(buf) // 读取字节数 n
		if err != nil {
			if err == io.EOF {
				// 判断文件读取结束
				break
			}
			log.Printf("打开文件失败:%#v", err)
		}
		dataBlocks = append(dataBlocks, buf[:n]...) // 注意有人这里[:n] 是读的字节数赋值，最后一次读取可能小于buf定义量
	}
	return dataBlocks
}

// 填充
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	fmt.Println(string(pad))
	return append(src, pad...)
}

// 记号填充
func markpadding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum + 1)}, padNum)
	fmt.Println(string(pad))
	return append(src, pad...)
}

// 去掉填充
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func EncryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	fmt.Println(len(src))
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 分块记号加密
func MarkEncryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	//给 src+1
	src = append(src, byte('m'))
	//再进行标记填充
	src = markpadding(src, block.BlockSize())
	//fmt.Println("加密了：", len(src))
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 加密-32字节版
func EncryptAES32(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	src = padding(src, 32)
	fmt.Println(len(src))
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func DecryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func main() {
	username := "csb"
	filename := "tupian2"
	desDir := "./temp/"
	//获取文件字节流
	file_stream := ReadFileToByteStream("./temp/tupian2.png")
	//获取密钥字节流
	key := ReadFileToByteStream("./temp/key.txt")

	//aes加密
	enc_file_stream, err := EncryptAES(file_stream, key)
	if err != nil {
		log.Println("aes加密时发生错误！", err.Error())
	}
	// 写入文件
	WriteStringToFile(desDir+filename+"_enc_"+username, string(enc_file_stream))

	//读取加密文件测试解密
	enc_file_stream = ReadFileToByteStream("./temp/tupian2_enc_csb")
	file_stream, err = DecryptAES(enc_file_stream, key)

	//写入解密后的
	WriteStringToFile(desDir+"tupian2_recover.png", string(file_stream))
}
