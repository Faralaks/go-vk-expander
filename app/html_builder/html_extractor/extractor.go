// Package html_extractor
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u extract data from html message history from VK.com
package html_extractor

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var excludeDSStore = []string{".DS_Store"}

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

func GetNumFromMsgFilename(name string) (int, error) {
	num := strings.TrimPrefix(strings.TrimSuffix(name, ".html"), "messages")
	return strconv.Atoi(num)
}

func (f Files) SortByNumber() Files {
	sort.Slice(f, func(i, j int) bool {
		iNum, iErr := GetNumFromMsgFilename(f[i])
		jNum, jErr := GetNumFromMsgFilename(f[j])
		if iErr != nil || jErr != nil {
			panic(fmt.Errorf("could not take int number from one of files (\"%s\", \"%s\") with following errors iErr: %v, jErr: %v", f[i], f[j], iErr, jErr))
		}
		return iNum < jNum
	})
	return f
}

func ExcludeFilenames(f Files, blackList []string) Files {
	res := make(Files, 0)
	for _, name := range f {
		if !IsNameInList(name, blackList) {
			res = append(res, name)
		}
	}
	return res
}

func Extract(p string) error {
	dialogs := make(map[string]Files)
	dialogList, err := GetFiles(p)
	if err != nil {
		log.Printf("[ERROR] Could not get files from message folder | %v", err)
		return fmt.Errorf("could not get files from message folder | %v", err)
	}
	dialogList = ExcludeFilenames(dialogList, excludeDSStore)
	for _, dialog := range dialogList {
		msgList, err := GetFiles(filepath.Join(p, dialog))
		if err != nil {
			log.Printf("[ERROR] Could not get files from message folder | %v", err)
			return fmt.Errorf("could not get files from message folder | %v", err)
		}
		msgList = ExcludeFilenames(msgList, excludeDSStore)
		msgList = msgList.SortByNumber()
		dialogs[dialog] = msgList
	}
	for k, v := range dialogs {
		fmt.Printf("Dialog: %v, Filse: %v\n", k, v)
	}

	return nil

}
