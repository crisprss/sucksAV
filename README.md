# sucksAV
This project used to learn golang and try to bypass AV

## 描述
基于Golang开发的BypassAV,采取的shellcode分离技术，将shellcode注入到图片中，通过加载器进行加载，使用Golang动态加载技术

>需要使用第三方库 github.com/Binject/universal

## Usage

**生成附加shellcode的图片**
```
go run generate.go xx.jpeg(最好使用jpeg文件) 
```

**CS上线**

将图片放在未压缩的图床中或VPS上，使用`Garble`进行混淆,garble的项目地址:https://github.com/burrowers/garble

>golang 调用cmd下程序隐藏黑窗口-方法 编译go时加入参数： go build -ldflags="-H windowsgui"


```
garble build loader.go
#也可以不使用garble 直接使用go build
go build loader.go
```

未使用Garbel混淆的免杀效果为:
![](https://github.com/crisprss/sucksAV/blob/main/images/image1.png)
使用Garbel混淆的免杀效果为::
![](https://github.com/crisprss/sucksAV/blob/main/images/image2.png)

截止到2021.9.21 两者均能Bypass 火绒+360
![](https://github.com/crisprss/sucksAV/blob/main/images/upload2.png)

