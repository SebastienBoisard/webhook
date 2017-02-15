package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func handleGitLab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleGitLab")
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("request: ", string(requestDump))

	// Doc: https://docs.gitlab.com/ce/user/project/integrations/webhooks.html

	token := r.Header.Get("X-Gitlab-Token")
	if len(token) == 0 {
		fmt.Println("No token!")
	}

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

	switch event {
	case "Push Hook":
		var push PushEvent
		err = json.NewDecoder(r.Body).Decode(&push)
		if err != nil {
			fmt.Println("Error while decoding push event")
			break
		}
		//fmt.Printf("push event: %+v\n", push)
		res, _ := json.Marshal(&push)
		fmt.Println("push event:", string(res))

	case "Tag Push Hook":
		var tagPush TagPushEvent
		err = json.NewDecoder(r.Body).Decode(&tagPush)
		if err != nil {
			fmt.Println("Error while decoding tag push event")
			break
		}
		res, _ := json.Marshal(&tagPush)
		fmt.Println("tag push event:", string(res))

	case "Issue Hook":
		var issue IssueEvent
		err = json.NewDecoder(r.Body).Decode(&issue)
		if err != nil {
			fmt.Println("Error while decoding issue event")
			break
		}
		res, _ := json.Marshal(&issue)
		fmt.Println("issue event:", string(res))

	case "Note Hook":
		var note NoteEvent
		err = json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			fmt.Println("Error while decoding note event")
			break
		}
		res, _ := json.Marshal(&note)
		fmt.Println("note event:", string(res))

	case "Merge Request Hook":
		var mergeRequest MergeRequestEvent
		err = json.NewDecoder(r.Body).Decode(&mergeRequest)
		if err != nil {
			fmt.Println("Error while decoding merge request event")
			break
		}
		res, _ := json.Marshal(&mergeRequest)
		fmt.Println("merge request event:", string(res))

	case "Wiki Page Hook":
		var wikiPage WikiPageEvent
		err = json.NewDecoder(r.Body).Decode(&wikiPage)
		if err != nil {
			fmt.Println("Error while decocding wiki page event")
			break
		}
		res, _ := json.Marshal(&wikiPage)
		fmt.Println("wiki page event:", string(res))

	case "Pipeline Hook":
		var pipeline PipelineEvent
		err = json.NewDecoder(r.Body).Decode(&pipeline)
		if err != nil {
			fmt.Println("Error while decocding pipeline event")
			break
		}
		res, _ := json.Marshal(&pipeline)
		fmt.Println("pipeline event:", string(res))

	case "Build Hook":
		var build BuildEvent
		err = json.NewDecoder(r.Body).Decode(&build)
		if err != nil {
			fmt.Println("Error while decoding build event")
			break
		}
		res, _ := json.Marshal(&build)
		fmt.Println("build event:", string(res))

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
