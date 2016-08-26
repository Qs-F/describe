package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Qs-F/coloring"
)

func main() {
	_, err := os.Open(".describe")
	m := flag.Bool("init", false, "init option touch .describe file in current directory")
	flag.Parse()
	if *m == true && err != nil { // if err is nil, there is a file.
		_, err := os.Create(".describe")
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println("success!")
			return
		}
		return
	} else if *m == true && err == nil {
		fmt.Println("There is a .describe file. Please remove it.")
		return
	}
	wd, err := os.Getwd()
	os.Chdir(wd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dir := ""
	if len(os.Args) > 1 {
		dir = os.Args[1] + "/"
	} else {
		dir = "./"
	}
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, f := range fs {
		if f.IsDir() {
			b, err := ioutil.ReadFile(dir + f.Name() + "/.describe")
			if err != nil {
				continue
			} else {
				fmt.Printf(`%s
  %s`, wd+"/"+f.Name(), coloring.Yellow(fmt.Sprintf("%s", b)))
			}
		}
	}
}
