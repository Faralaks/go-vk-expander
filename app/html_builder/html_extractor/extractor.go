// Package html_extractor
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u extract data from html message history from VK.com
package html_extractor

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"io/fs"
	"io/ioutil"
)

type Files []fs.FileInfo

func GetFiles(p string) (Files, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Printf("[DEBUG] Could not read \"%s\" directory | %v", p, err)
		return files, fmt.Errorf("could not read \"%s\" directory | %v", p, err)
	}
	return files, nil
}

func isNameInList(name string, list []string) bool {
	for _, listName := range list {
		if name == listName {
			return true
		}
	}

	return false
}

func (f Files) ExcludeFilenames(blackList []string) {
	res := make(Files, 0)
	for _, file := range f {
		if !isNameInList(file.Name(), blackList) {
			res = append(res, file)
		}
	}
}

func Extract(p string) error {
	println(p)
	return nil

}
