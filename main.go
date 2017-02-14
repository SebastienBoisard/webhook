package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handleGitLab(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

var address = flag.String("addr", "localhost:8080", "http service address")

//var token = flag.String("token", "1234", "token to identify this client")

func main() {

	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime)

	http.HandleFunc("/gitlab", handleGitLab)
	log.Fatal(http.ListenAndServe(*address, nil))
}
