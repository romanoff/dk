package source

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Fs struct {
	Paths []string
}

func (self *Fs) Setup(config *Config) {
	self.Paths = config.Paths
}

func (self *Fs) CreateDump(path string) error {
	perm := os.FileMode(0777)
	for i, filePath := range self.Paths {
		folderPath := fmt.Sprintf("%v/%v", path, i)
		err := os.MkdirAll(folderPath, perm)
		if err != nil {
			return err
		}
		cmd := exec.Command("cp", "-r", filePath, folderPath)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *Fs) ApplyDump(path string) error {
	for i, filePath := range self.Paths {
		folderPath := fmt.Sprintf("%v/%v", path, i)
		files, _ := ioutil.ReadDir(folderPath)
		for _, file := range files {
			folderFilePath := folderPath + "/" + file.Name()
			if file.IsDir() {
				folderFilePath += "/"
			}
			cmd := exec.Command("rsync", "-ar", "--del", folderFilePath, filePath)
			err := cmd.Run()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
