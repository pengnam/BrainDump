package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"./util"

	"gopkg.in/russross/blackfriday.v2"
)

const (
	thresholdRecentFiles = 20
)

type Topic struct {
	Name  string
	Files []File
}

func parseMarkdown(filePath string, fileName string, topic string) File {
	title := strings.TrimSuffix(fileName, ".md")

	return File{topic, title, fileName, parseFileContent(filePath + fileName), getLastModified(filePath + fileName).String(), title + ".html"}
}

// parseFile parses markdown into HTML and strips the first h1 tag of the html
func parseFileContent(file string) string {
	bytes, _ := ioutil.ReadFile(file)
	re := regexp.MustCompile("<h1>(.*)</h1>")
	result := blackfriday.Run(bytes)
	header := re.Find(result)
	return string(result[len(header):])
}

// getLastModified returns the date that file was last modified
func getLastModified(file string) time.Time {
	f, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	return f.ModTime()
}

func retrieveMarkdowns(folderName string) []string {
	infos, _ := ioutil.ReadDir(folderName)
	result := make([]string, 0)
	for _, ele := range infos {
		if strings.HasSuffix(ele.Name(), ".md") {
			result = append(result, ele.Name())
		}
	}
	return result
}

type File struct {
	Topic   string
	Title   string
	Name    string
	Content string
	Time    string
	Path    string
}

func renderFile(content File) error {
	fileTmpl := template.New("file")
	tmpl, err := ioutil.ReadFile("templates/file.tmpl")
	if err != nil {
		return err
	}
	_, err = fileTmpl.Parse(string(tmpl))
	if err != nil {
		return err
	}

	writer, err := os.Create("../docs/" + content.Title + ".html")
	if err != nil {
		return err
	}
	fileTmpl.Execute(writer, content)
	return nil
}

func renderIndex(files []File) error {
	indexTmpl := template.New("index")
	tmpl, err := ioutil.ReadFile("templates/index.tmpl")
	if err != nil {
		return err
	}
	_, err = indexTmpl.Parse(string(tmpl))
	if err != nil {
		return err
	}

	writer, err := os.Create("../docs/index.html")
	if err != nil {
		return err
	}
	indexTmpl.Execute(writer, files)
	return nil
}

func getTopFiles(topics []Topic) []File {
	files := []File{}
	for _, t := range topics {
		files = append(files, t.Files...)
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Time > files[j].Time
	})
	return files[:thresholdRecentFiles]
}

func parseFilesInTopic(topic string) Topic {
	fileNames := retrieveMarkdowns("../" + topic)
	files := []File{}
	for _, ele := range fileNames {
		fmt.Println("../" + topic + "/" + ele)
		f := parseMarkdown("../"+topic+"/", ele, topic)
		err := renderFile(f)
		if err != nil {
			panic(err)
		}
		files = append(files, f)
	}
	return Topic{topic, files}
}

func renderListing(topics []Topic) error {
	listingsTmpl := template.New("listing")
	tmpl, err := ioutil.ReadFile("templates/listing.tmpl")
	if err != nil {
		return err
	}
	_, err = listingsTmpl.Parse(string(tmpl))
	if err != nil {
		return err
	}

	writer, err := os.Create("../docs/listing.html")
	if err != nil {
		return err
	}
	listingsTmpl.Execute(writer, topics)
	return nil
}

func main() {
	fmt.Println("Testing")

	folders := util.SearchFolder("../")
	topics := []Topic{}
	for _, folder := range folders {
		fmt.Println(folder)
		topic := parseFilesInTopic(folder)
		topics = append(topics, topic)
	}
	err := renderIndex(getTopFiles(topics))
	if err != nil {
		panic(err)
	}
	err = renderListing(topics)
	if err != nil {
		panic(err)
	}
}

/*
Create and render index
*/
