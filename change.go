package main

import (
	"fmt"
	"github.com/IntelligenceX/fileconversion/xls"
	"io/ioutil"
	"os"
)

func main() {
	//dir := flag.String("n", "./", "dir")
	//flag.Parse()

	fmt.Println("请输入要修改的文件夹绝对地址:")
	var dir string
	fmt.Scanf("%s",&dir)

	files, _ := ioutil.ReadDir(dir)
	for _, i := range files {

		// 读取文件名
		name := i.Name()
		// 反转数组，判断是否为xls文件CGO_ENABLED=0 GOOS=windows  go build
		r := []rune(name)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		reverseName := string(r)

		if len(reverseName) > 4 && reverseName[:4] == "SLX." {
			xlFile,closer,err := xls.OpenWithCloser(dir+"/"+name,"utf-8")
			if err!=nil{
				fmt.Println(err)
			}
			sheet1 := xlFile.GetSheet(0)
			cell := sheet1.Row(5).Col(1)
			closer.Close()
			os.Rename(dir+"/"+name, dir+"/"+cell+".XLS")
			fmt.Printf("修改前文件名为:%s,修改后文件名为:%s.XLS\n", name, cell)
		}

	}
}
