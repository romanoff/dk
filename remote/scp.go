package remote

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Scp struct {
	Path string // username@server.com:/home/username/files_folder
}

func (self *Scp) Setup(config *Config) {
	self.Path = strings.TrimSpace(config.Path)
	if self.Path[len(self.Path)-1:] != "/" {
		self.Path = self.Path + "/"
	}
}

func (self *Scp) Push(filepath, destination string) error {
	args := []string{filepath, self.Path + destination}
	cmd := exec.Command("scp", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func (self *Scp) Pull(filepath, destination string) error {
	args := []string{self.Path + filepath, destination}
	cmd := exec.Command("scp", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

func (self *Scp) FilesList() ([]string, error) {
	paths := strings.Split(self.Path, ":")
	serverPath := paths[0]
	out, err := exec.Command("ssh", serverPath, "cd "+paths[1]+" && ls").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(out), "\n"), nil
}

func (self *Scp) Remove(filepath string) error {
	paths := strings.Split(self.Path, ":")
	serverPath := paths[0]
	cmd := exec.Command("ssh", serverPath, "cd "+paths[1]+" && rm "+filepath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
