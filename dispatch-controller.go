package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getBerarer() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	bearer, err := os.ReadFile(exPath + "/bearer.token")
	if err != nil {
		log.Fatal(err)
	}
	cleanup := strings.Replace(string(bearer), "\r", "", -1)
	cleanup = strings.Replace(string(bearer), "\n", "", -1)
	fmt.Println("Bearer: ", string(cleanup))
	return cleanup
}

type Dispatch struct {
	EventType string `json:"event_type"`
}

type UserInput struct {
	UserName  string
	RepoName  string
	EventType string
}

func collectFlags() UserInput {
	userName := flag.String("userName", "", "Name of user to get repo from")

	repoName := flag.String("repoName", "", "Name of repo to target for dispatch")

	eventType := flag.String("eventType", "", "Name of event type inside repo")

	flag.Parse()
	var errSlice []string
	if len(*userName) == 0 {
		errSlice = append(errSlice, "Please specify git -userName (or repo)")
	}
	if len(*repoName) == 0 {
		errSlice = append(errSlice, "Please specify -repoName")
	}
	if len(*eventType) == 0 {
		errSlice = append(errSlice, "Please specify -eventType")
	}

	if len(errSlice) > 0 {
		for i := 0; i < len(errSlice); i++ {
			log.Println(errSlice[i])
		}
		log.Fatalln("Please provide the above credentials...")
	}
	userInput := UserInput{
		UserName:  *userName,
		RepoName:  *repoName,
		EventType: *eventType,
	}
	fmt.Println("userName to use: ", *userName)
	fmt.Println("repoName to target: ", *repoName)
	fmt.Println("eventType to start: ", *eventType)

	return userInput
}

func main() {
	userInput := collectFlags()
	bearer := getBerarer()
	baseURL := "https://api.github.com/repos/"

	url := baseURL + userInput.UserName + "/" + userInput.RepoName + "/" + "dispatches"
	fmt.Println(url)

	dispatch := Dispatch{
		EventType: userInput.EventType,
	}
	dispatchJSON, _ := json.Marshal(dispatch)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(dispatchJSON))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearer)

	res, err := http.DefaultClient.Do(req)
	fmt.Println(res)
	if err != nil {
		log.Fatal("Could not make POST request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}
