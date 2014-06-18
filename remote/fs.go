package remote

import (
	"io/ioutil"
	"os"
	"io"
)

type Fs struct {
	Path string // root path
}

func (self *Fs) Setup(config *Config) {
	self.Path = config.Path
}

func (self *Fs) Push(filepath, destination string) error {
	sourceFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(self.Path + "/" + destination)
	if err != nil {
		return err
	}
	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		destinationFile.Close()
		return err
	}
	return destinationFile.Close()
}

func (self *Fs) Pull(filepath, destination string) error {
	sourceFile, err := os.Open(self.Path + "/" + filepath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		destinationFile.Close()
		return err
	}
	return destinationFile.Close()
	return nil
}

func (self *Fs) FilesList() ([]string, error) {
	filesList := []string{}
	files, err := ioutil.ReadDir(self.Path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		filesList = append(filesList, f.Name())
	}
	return filesList, nil
}
