package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func hando(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

files:=[]string{
	"/home/clomollo/algo1/web/ui/html/base.html",
	"/home/clomollo/algo1/web/ui/html/index.html",
	"/home/clomollo/algo1/web/ui/html/nav.html",
}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err=ts.ExecuteTemplate(w,"base",nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// w.Write([]byte("Hello From Snippet View"))
}

func viewSnip(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil || id < 0 {
		http.Error(w, "Sorry mate! Put positive number", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Let's display a specific snipo by %d", id)
}

func createSnip(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("aLLOW", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Careful buddy..Method Not Allowed Here!!"))
		http.Error(w, "Method Not Allowed Buddy", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Header().Values("Cache-Control")
	w.Write([]byte("I create things"))
}
