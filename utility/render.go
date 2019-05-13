package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Topic struct {
	Name string
	Path string
}

func newTopic(name, path string) Topic {
	return Topic{name, path}
}

func parseMarkdown(fileName string) []byte {
	return blackfriday.Run(parseFile(fileName))
}
func parseFile(fileName string) []byte {
	bytes, _ := ioutil.ReadFile(fileName)
	return bytes
}

func otherSearchFolder(folderName string) []string {
	result := []string{}
	filepath.Walk(folderName, func(path string, _ os.FileInfo, _ error) error {

		if strings.HasSuffix(path, ".md") {
			result = append(result, path)
		}
		return nil
	})
	return result
}

func searchFolder(folderName string) []string {
	blackList := []string{".DS_Store", ".git", "util"}
	infos, _ := ioutil.ReadDir(folderName)
	result := make([]string, 0)
	for _, ele := range infos {
		if !contains(blackList, ele.Name()) {
			result = append(result, ele.Name())
		}
	}
	return result
}

func main() {
	fmt.Println("Testing")
	//searchFolder("../")

	folders := searchFolder("../")
	for _, folder := range folders {
		fmt.Println(folder)
		files := otherSearchFolder(folder)
		for _, ele := range files {
			fmt.Println(ele)
			output := parseMarkdown(ele)
			fmt.Println(string(output))
		}
	}

}

/*
Create and render index
*/
