package gitcrawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GitUser struct {
	Login   string
	URL     string
	HTMLURL string `json:"html_url"`
	Bio     string `json:"bio"`
}

func GetRandomUsers() (*[]GitUser, error) {
	request, err := http.NewRequest("GET", "https://api.github.com/users", nil)

	if err != nil {
		log.Fatalln(err)
	}

	bearerToken := os.Getenv("GIT_ACCESS_TOKEN")
	request.Header.Set("Authorization", "Bearer "+bearerToken)

	sinceMin := 1
	sinceMax := 1
	since := strconv.Itoa(rand.Intn(sinceMax) + sinceMin)

	requestQuery := request.URL.Query()
	requestQuery.Add("since", since)
	request.URL.RawQuery = requestQuery.Encode()

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	fmt.Printf("Making request to Github API here: %v", request.URL.RawQuery)
	response, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(response)
	}

	gitUsers, err := handleRandomUsersResponse(response)

	if err != nil {
		fmt.Println("Error getting random users!!")
	} else {
		fmt.Printf("Got users: %v", (*gitUsers))
	}

	return gitUsers, err
}

/*HandleResponse does...
 */
func handleRandomUsersResponse(response *http.Response) (*[]GitUser, error) {
	var gitUsers []GitUser

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln("Failed in HandleResponse's ReadAll! ", err)

	}

	err = json.Unmarshal(body, &gitUsers)

	if err != nil {
		return nil, errors.New("error retrieving users")
	}
	return &gitUsers, nil
}
