package link

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/quiteawful/WebEssentials/global"
)

var linkMap map[string]string
var shortenerBase string

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	shortenerBase = global.Conf.BaseURL + ":" + strconv.Itoa(global.Conf.Port) + "/l/"
	os.MkdirAll(global.Conf.Linksdir, 0755)
	linkMap = make(map[string]string)
	filename := global.Conf.Linksdir + "linklist.json"
	cont, err := global.Exists(filename)
	if err != nil {
		log.Fatal(err)
	}
	if !cont {
		err := SaveLinkList()
		if err != nil {
			log.Fatal(err)
		}
	}
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &linkMap)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveLinkList() error {
	b, err := json.MarshalIndent(linkMap, "", "	")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(global.Conf.Linksdir+"linklist.json", b, 0777)
	return err
}

func randomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func GenerateNewShortURL(in *url.URL) (string, error) {
	res := "random"
	var u string
	for res != "" {
		u = randomString(6)
		res = linkMap[u]
	}
	linkMap[u] = in.String()
	err := SaveLinkList()
	if err != nil {
		return "", err
	}
	return shortenerBase + u, nil
}

func GetRealURL(c string) string {
	return linkMap[c]
}
