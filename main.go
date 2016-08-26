package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Qs-F/coloring"
)

func main() {
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
