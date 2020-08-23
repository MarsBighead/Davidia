package main

import (
	"Davidia/pkg/process"
	"Davidia/pkg/util"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	currentPath, _ := util.GetCurrentPath(os.Args[0])
	//log.Fatalln(currentPath)
	cfg, err := util.Parse(currentPath + "/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cfg.Source)
	ss, err := cfg.Source.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ss.Header)
	fmt.Println(`
	Please input the path of Health QR Code Pictures:
	请将解压后的健康码文件夹目录输入该命令行(建议采用鼠标拖入):
	`)
	var qrCodePath string
	fmt.Scanln(&qrCodePath)
	fmt.Print(qrCodePath)
	r, err := process.Check(qrCodePath, ss.Student)
	if err != nil {
		log.Fatal(err)
	}
	today := strings.Split(time.Now().String(), " ")[0]
	fmt.Print(today)
	r.Date = today
	reportFilename := fmt.Sprintf("%s_%s.xlsx", today, cfg.Output.Base)
	r.Header = []string{
		"日期",
		"姓名",
		"学号",
		"Status",
		"格式化是否正确",
		"健康码文件名",
	}
	err = r.Output(reportFilename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Verify Health QR Code pictures file successfully, and more details in file %s.\n", reportFilename)
	log.Println("Program will exit in 10 seconds.")
	time.Sleep(10 * time.Second)

}
