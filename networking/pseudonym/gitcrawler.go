package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	request, err := http.NewRequest("GET", "https://api.github.com/users", nil)

	if err != nil {
		log.Fatalln(err)
	}

	bearerToken := os.Args[1]
	request.Header.Set("Authorization", "Bearer "+bearerToken)
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}

	HandleResponse(response)
	fmt.Println("All done~~")
}

/*HandleResponse does...
 */
func HandleResponse(response *http.Response) {
	type GitUser struct {
		Login             string
		ID                int
		NodeID            string
		AvatarURL         string
		GravatarID        string
		URL               string
		HTMLURL           string
		FollowersURL      string
		FollowingURL      string
		GistsURL          string
		StarredURL        string
		SubscriptionsURL  string
		OrganizationsURL  string
		ReposURL          string
		EventsURL         string
		ReceivedEventsURL string
		Type              string
		SiteAdmin         bool
	}

	fmt.Println("IN HERE??")

	var gitUsers []GitUser

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("Failed in HandleResponse's ReadAll! ", err)

	}

	err = json.Unmarshal(body, &gitUsers)

	if err != nil {
		fmt.Println("Failed in HandleResponses's Unmarshal!", gitUsers)
	} else {
		fmt.Println(len(gitUsers))
		fmt.Println(gitUsers)
	}
}
