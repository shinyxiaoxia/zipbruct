package main

import (
	"bufio"
	"flag"

	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/yeka/zip"
)

func zipbb(password string, file string) {
	r, err := zip.OpenReader(file)
	if nil != err {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}
		r, err := f.Open()
		if err != nil {
			// log.Fatal(err)
			break
		}
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		color.Set(color.FgGreen)
		fmt.Println("文件密码是：", password)
		fmt.Println("size of %v byte(s)\n", f.Name, len(buf))
		color.Unset()
		os.Exit(3)
	}
}

func fileReadline(passfile string) []string {
	ss := make([]string, 0)
	file1, err := os.Open(passfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		linetest := scanner.Text()
		ss = append(ss, linetest)
	}
	return ss
}

var (
	f string
	d string
	h bool
)

func init() {

	flag.BoolVar(&h, "h", false, "帮助文档")
	flag.StringVar(&f, "f", "", "-f xxx.zip [加密的zip文档]")
	flag.StringVar(&d, "d", "", "-d xxx.txt [用来解密的密码字典]")

	flag.Usage = usage
}

func usage() {
	color.Green(`			mzfuzz zip文件破解
Usage: zipbruct [hfd] [-f xxx.zip] [-d pass.txt]
Options:
`)
	color.Set(color.FgCyan)
	flag.PrintDefaults()
	color.Unset()
}
func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}
	// fmt.Println(os.Args)
	if len(os.Args) < 2 {
		flag.Usage()
	}
	aa := fileReadline(d)
	for _, value := range aa {
		zipbb(value, f)

	}

}
