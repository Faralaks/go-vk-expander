// Package html_extractor
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u extract data from html message history from VK.com
package html_extractor

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"io/ioutil"
)

type Files []string

func GetFiles(p string) (Files, error) {
	res := make(Files, 0)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		log.Printf("[DEBUG] Could not read \"%s\" directory | %v", p, err)
		return res, fmt.Errorf("could not read \"%s\" directory | %v", p, err)
	}
	for _, file := range files {
		res = append(res, file.Name())
	}
	return res, nil
}

func IsNameInList(name string, list []string) bool {
	for _, listName := range list {
		if name == listName {
			return true
		}
	}

	return false
}

func (f Files) ExcludeFilenames(blackList []string) Files {
	res := make(Files, 0)
	for _, name := range f {
		if !IsNameInList(name, blackList) {
			res = append(res, name)
		}
	}
	return res
}

func Extract(p string) error {
	println(p)
	return nil

}
