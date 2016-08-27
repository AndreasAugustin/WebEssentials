package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quiteawful/WebEssentials/global"
	"github.com/quiteawful/WebEssentials/link"
)

type FrontPage struct {
	Message string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/l/{code}", urlResolveHandler)
	r.HandleFunc("/shorten", newShortURLHandler)
	r.HandleFunc("/", frontpageHandler)
	port := ":" + strconv.Itoa(global.Conf.Port)
	log.Fatal(http.ListenAndServe(port, r))
}

func showLinks(w http.ResponseWriter, r *http.Request, fp FrontPage) {
	t, err := template.ParseFiles(global.Execdir + "/templates/frontpage.html") //Lade Template
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	err = t.Execute(w, &fp) //Zeige template an
	if err != nil {
		fmt.Println("There was an error:", err)
	}
}

func frontpageHandler(w http.ResponseWriter, r *http.Request) {
	var fp FrontPage
	showLinks(w, r, fp)
}

func urlResolveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	re := regexp.MustCompile("^[a-zA-Z0-9]{6}$")
	if !re.MatchString(code) {
		http.Error(w, "URL Code mallformed", http.StatusBadRequest)
		return
	}
	red := link.GetRealURL(code)
	http.Redirect(w, r, red, http.StatusTemporaryRedirect)
}

func newShortURLHandler(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("url")
	ur, err := url.Parse(u)
	if err != nil {
		http.Error(w, "URL mallformed", http.StatusBadRequest)
		return
	}
	nu, err := link.GenerateNewShortURL(ur)
	if err != nil {
		http.Error(w, "Error saving link list", http.StatusInternalServerError)
		return
	}
	var fp FrontPage
	fp.Message = "Shortend URL: " + nu
	showLinks(w, r, fp)
}
