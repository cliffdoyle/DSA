package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	adrr:=flag.String("adrr",":4000","HTTP network address")
	flag.Parse()

	infolog:=log.New(os.Stdout,"INFO\n",log.Ldate|log.Ltime)
	errorlog:=log.New(os.Stderr,"ERROR\n",log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	files:=http.FileServer(http.Dir("/home/clomollo/algo1/web/ui/static/"))
	mux.Handle("/static/",http.StripPrefix("/static",files))
	mux.HandleFunc("/", hando)
	mux.HandleFunc("/create", createSnip)
	mux.HandleFunc("/snipview", viewSnip)
	serv := http.Server{
		Addr: *adrr,
		Handler: mux,
		ErrorLog: errorlog,
	}
	infolog.Printf("Starting server on %s",*adrr)
	err := serv.ListenAndServe()
	errorlog.Fatal(err)
}
