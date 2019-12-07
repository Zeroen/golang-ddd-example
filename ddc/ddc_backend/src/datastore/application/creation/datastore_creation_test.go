package creation

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func iSendAPOSTRequestToWithABody(arg1 string, arg2 *gherkin.DocString) error {
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
