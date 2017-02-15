package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"io"
)

func handleGitLab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGitLab")
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	// Doc: https://docs.gitlab.com/ce/user/project/integrations/webhooks.html

	token := r.Header.Get("X-Gitlab-Token")
	if len(token) == 0 {
		fmt.Println("No token!")
	}
	fmt.Println("token=", token)

	if token != *gitlabToken {
		fmt.Println("Wrong token!")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "{}")
		return
	}

	event := r.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		fmt.Println("No event!")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "{}")
		return
	}
	fmt.Println("event=", event)

	switch event {
	case "Push Hook":
		fmt.Println("Push Hook")
	case "Tag Push Hook":
		fmt.Println("Tag Push Hook")
	case "Issue Hook":
		fmt.Println("Issue Hook")
	case "Note Hook":
		fmt.Println("Note Hook")
	case "Merge Request Hook":
		fmt.Println("Merge Request Hook")
	case "Wiki Page Hook":
		fmt.Println("Wiki Page Hook")
	case "Pipeline Hook":
		fmt.Println("Pipeline Hook")
	case "Build Hook":
		fmt.Println("Build Hook")
	default:
		fmt.Println("unknown event")
	}


	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
	return
}

var address = flag.String("addr", "192.168.1.26:8080", "http service address")

var gitlabToken = flag.String("token", "test", "token to verify is the request is legitimate")

func main() {

	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime)

	http.HandleFunc("/gitlab", handleGitLab)
	log.Fatal(http.ListenAndServe(*address, nil))
//	log.Fatal(http.ListenAndServeTLS(*address, "server.crt", "server.key", nil))
}
