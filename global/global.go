package global

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var Conf Config
var Execdir string

type Config struct {
	Port     int
	BaseURL  string
	Imagedir string
	Linksdir string
	Userdir  string
}

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	Execdir = dir + "/"
	err = Conf.load()
	if err != nil {
		log.Fatal("Could not load config: ", err.Error())
	}
	if Conf.BaseURL == "" {
		Conf.BaseURL = "http://127.0.0.1"
	}
	if Conf.Imagedir == "" {
		Conf.Imagedir = Execdir + "images/"
	}
	if Conf.Linksdir == "" {
		Conf.Linksdir = Execdir + "links/"
	}
	if Conf.Userdir == "" {
		Conf.Userdir = Execdir + "user/"
	}
}

func (c *Config) load() error {
	filename := Execdir + "config.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		return err
	}
	return nil
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
