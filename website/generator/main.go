package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	files, _ := ioutil.ReadDir("../markdown")
	for _, f := range files {
		fmt.Println(f.Name())
		name := f.Name()
		path := filepath.Dir("./home.md")
		fmt.Println(path)
		return
		htmlName := strings.ReplaceAll(name, ".md", ".html")
		exec.Command("pandoc", "-o", htmlName, name)
		// pandoc -o '${}%.md'.html {} -s -H css/stylesheet.html \;
	}
}
