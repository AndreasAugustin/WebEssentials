package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quiteawful/WebEssentials/global"
	"github.com/quiteawful/WebEssentials/link"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/l/{code}", urlResolveHandler)
	port := ":" + strconv.Itoa(global.Conf.Port)
	log.Fatal(http.ListenAndServe(port, r))
}

func urlResolveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	fmt.Println(code)
	re := regexp.MustCompile("^[a-zA-Z0-9]{6}$")
	if !re.MatchString(code) {
		http.Error(w, "URL Code mallformed", http.StatusBadRequest)
		return
	}
	u, err := url.Parse("https://" + r.Host + r.URL.String())
	if err != nil {
		http.Error(w, "URL mallformed", http.StatusBadRequest)
	}
	red := link.GetRealURL(u)
	http.Redirect(w, r, red, http.StatusTemporaryRedirect)
}
