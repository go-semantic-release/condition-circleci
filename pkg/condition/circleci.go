package condition

import "errors"

var CIVERSION = "dev"

type CircleCI struct {
}

func (cci *CircleCI) Name() string {
	return "CircleCI"
}

func (cci *CircleCI) Version() string {
	return CIVERSION
}

func (cci *CircleCI) GetCurrentBranch() string {
	return ""
}

func (cci *CircleCI) GetCurrentSHA() string {
	return ""
}

func (cci *CircleCI) RunCondition(config map[string]string) error {
	return errors.New("not implemented")
}
