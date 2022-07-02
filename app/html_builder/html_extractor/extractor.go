// Package html_extractor
// Coded by Faralaks https://github.com/Faralaks
// This package provides interface and implementation which allows u extract data from html message history from VK.com
package html_extractor

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	. "strings"
	"sync"
)

type dialogTreeMaker interface {
	MakeTree() (DialogTree, error)
}
type DialogMaker struct {
	Path string
}

func NewDialogMaker(p string) DialogMaker {
	return DialogMaker{Path: p}
}

type DialogTree map[string]*dialog

type dialog struct {
	Files []msgFile
	sync.Mutex
}
type msgFile struct {
	File string
	Path string
}

// getNumber returns int value from message filename
func (mf msgFile) getNumber() int {
	strNum := TrimSuffix(TrimPrefix(mf.File, "messages"), ".html")
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		log.Printf("[ERROR] Could not convert number from message filename to int")
		panic(fmt.Errorf("it seams wrong filename were given to getNumber func case number could not be converted to int"))
		return 0
	}
	return intNum
}

func (d DialogMaker) MakeTree() (DialogTree, error) {
	p := filepath.Join(d.Path, "messages")
	dialogs := make(DialogTree)

	dialogDIrs, err := ioutil.ReadDir(p)
	if err != nil {
		log.Printf("[DEBUG] Could not read dialogs directory | %v", err)
		return nil, err
	}

	for _, dir := range dialogDIrs {
		if dir.Name() == ".DS_Store" {
			continue
		}
		dialogFiles, err := ioutil.ReadDir(filepath.Join(p, dir.Name()))
		if err != nil {
			log.Printf("[DEBUG] Could not read msg files | %v", err)
			return nil, err
		}
		msgFiles := make([]msgFile, 0, len(dialogFiles))
		for _, file := range dialogFiles {
			if file.Name() == ".DS_Store" {
				continue
			}
			newMsg := msgFile{
				File: file.Name(),
				Path: filepath.Join(p, dir.Name(), file.Name()),
			}
			msgFiles = append(msgFiles, newMsg)
		}
		dialogs[dir.Name()] = &dialog{Files: msgFiles}

	}
	return dialogs, nil
}

func MakeSortedTree(treeMaker dialogTreeMaker) (DialogTree, error) {
	tree, err := treeMaker.MakeTree()
	if err != nil {
		return nil, err
	}
	for id, dialogFiles := range tree {
		sort.Slice(dialogFiles.Files, func(i, j int) bool {
			return dialogFiles.Files[i].getNumber() < dialogFiles.Files[i].getNumber()
		})
		tree[id] = dialogFiles

	}
	return tree, nil
}

func Extract(p string) {
	if _, err := os.Stat(filepath.Join(p, "messages")); err != nil {
		log.Printf("[ERROR] No such directory as " + p)
		return
	}
	d := NewDialogMaker(p)
	dialogs, err := MakeSortedTree(d)
	if err != nil {
		log.Printf("[ERROR] There is a problem while walking throw Archive directory | %v", err)
	}
	println()
	for _, v := range dialogs {
		fmt.Printf("%v\n", v.Files)
	}
}
