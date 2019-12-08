package infrastructure

import (
	"bytes"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"net/http"
)

func iSendAPOSTRequestToWithABody(url string, body *gherkin.DocString) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body.Content)))
	req.Header.Set("Content-Type", body.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
/*
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	resp, err := http.Post(url, "text/json", &buf)
*/
	return godog.ErrPending
}

func theResponseStatusCodeShouldBe(arg1 int) error {
	return godog.ErrPending
}

func theResponseShouldBeEmpty() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send a POST request to "([^"]*)" with a body:$`, iSendAPOSTRequestToWithABody)
	s.Step(`^the response status code should be (\d+)$`, theResponseStatusCodeShouldBe)
	s.Step(`^the response should be empty$`, theResponseShouldBeEmpty)
}
