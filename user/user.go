package user

import (
	"encoding/json"
	"io/ioutil"

	"github.com/quiteawful/WebEssentials/global"
)

type User struct {
	//Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"-"`
	Email    string `json:"email"`
}

func (u *User) Save(n string) error {
	b, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(global.Conf.Userdir+n+".json", b, 0777)
	return err
}

func (u *User) Load(n string) error {
	filename := global.Conf.Userdir + n + ".json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &u)
	return err
}

type UserLinks map[string]string

func (u *UserLinks) Save(n string) error {
	b, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(global.Conf.Userdir+n+"-links.json", b, 0777)
	return err
}

func (u *UserLinks) Load(n string) error {
	filename := global.Conf.Userdir + n + "-links.json"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &u)
	return err
}
