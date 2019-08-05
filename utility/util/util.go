package util

import "io/ioutil"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SearchFolder(folderName string) []string {
	blackList := []string{".DS_Store", ".git", "utility", "build", ".gitignore", "README.md", "LICENSE"}
	infos, _ := ioutil.ReadDir(folderName)
	result := make([]string, 0)
	for _, ele := range infos {
		if !Contains(blackList, ele.Name()) {
			result = append(result, ele.Name())
		}
	}
	return result
}
