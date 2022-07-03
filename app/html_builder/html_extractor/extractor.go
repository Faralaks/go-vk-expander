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
	"sync"
)

var excludeDSStore = []string{".DS_Store"}

// MsgFiles presents list of messages filenames from one dialog
type MsgFiles []string

// Dialog presents messages file list from one dialog amd following mutex
type Dialog struct {
	MsgFiles
	sync.Mutex
}

func NewDialog(msgFiles MsgFiles) Dialog {
	return Dialog{MsgFiles: msgFiles}
}

// GetFiles returns list of all files in specified folder
func GetFiles(p string) (MsgFiles, error) {
	res := make(MsgFiles, 0)
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

//IsNameInList checks name existence in list
func IsNameInList(name string, list []string) bool {
	for _, listName := range list {
		if name == listName {
			return true
		}
	}
	return false
}

// GetNumFromMsgFilename returns int value from messages file name.
// Example: input: "messages123.html" output: 123
func GetNumFromMsgFilename(name string) (int, error) {
	num := strings.TrimPrefix(strings.TrimSuffix(name, ".html"), "messages")
	return strconv.Atoi(num)
}

// SortByNumber sorts messages files by number in their name
// Example: input ["messages50.html", "messages0.html", "messages100.html"]
//  output: ["messages0.html", "messages50.html", "messages100.html"]
func SortByNumber(f MsgFiles) MsgFiles {
	// TAKE CARE: if given filename do not have like "messages<some int>.html it may cause undefined behaviour
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

// ExcludeFilenames  returns list without given elements
func ExcludeFilenames(f MsgFiles, blackList []string) MsgFiles {
	res := make(MsgFiles, 0)
	for _, name := range f {
		if !IsNameInList(name, blackList) {
			res = append(res, name)
		}
	}
	return res
}

// Extract start extraction and building process
func Extract(p string) error {
	dialogs := make(map[string]Dialog)
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
		msgList = SortByNumber(msgList)
		dialogs[dialog] = NewDialog(msgList)
	}

	return nil
}
