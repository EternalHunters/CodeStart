package main

import (
	"fmt"
	"github.com/EternalHunters/go-qrcode"
	"image/png"
	"os"
)

func checkErr(err error){
	if err!= nil{
		panic(err)
	}
}
func main(){
	// 打开文件
	f, err:= os.Open("icon.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// 解码logo图片
	logo,err:= png.Decode(f)
	checkErr(err)
	// 打开背景图
	bgF,err:= os.Open("xiong.png")
	checkErr(err)
	defer bgF.Close()
	// 解码背景
	bg,err:=png.Decode(bgF)
	//err = qr.WriteFileWithLogo("test.png", qr.Medium, "123", logo,200,200,5)
	err = qrcode.BGWriteFileWithLogo("test2.png", bg, logo,qrcode.Medium, "123", 200, 0)
	checkErr(err)
	//qr.WriteFile("123", qr.Medium, 200, "test2.png", 5)
}