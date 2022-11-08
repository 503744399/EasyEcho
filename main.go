package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var filepath string
var filepath1 string

func main() {
	flag.StringVar(&filepath, "t", "target", "文件路径，默认为空")
	flag.Parse()
	content, _ := ioutil.ReadFile(filepath)
	r := strings.NewReplacer("<", "^<", ">", "^>", "|", "^|", "&", "^&", "^", "^^")
	fmt.Println(r.Replace(string(content)))
}
