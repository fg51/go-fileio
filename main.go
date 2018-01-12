package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	savefile("out.txt")
}

// func savefile(filename string, texts []string) {
func savefile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("cannot open: %s\n", filename)
	}
	defer file.Close()

	texts := []string{"a", "b", "c"}

	var writer *bufio.Writer
	writer = bufio.NewWriter(file)
	for i := range texts {
		writer.WriteString(texts[i] + "\n")
	}
	writer.Flush()
}

//func main() {
//	path := os.Args[1]
//	files := list(path)
//	for idx := range files {
//		var filename = files[idx]
//    fmt.Println(filename)
//    tests := readfile(filename)
//    if texts[0] == "---" {
//      continue
//    }
//    title := "title: " + texts[0]
//    texts = append([]string{
//      "---",
//      "layout: default",
//      title,
//      "---",
//      ""
//    }, texts...)
//    savefile(filename, texts)
//	}
//}

//func list(dir string) (matches []string) {
//  files, _ := filepath.Glob(dir + "/
//  fp, err := os.Open(filename)
//  var ret []string
//  if err
//}
