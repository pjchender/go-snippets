package enum

import "fmt"

type ActivationStatus string

const (
	Enabled   ActivationStatus = "enabled"
	Disabled  ActivationStatus = "disabled"
)

var activationStatus = map[ActivationStatus]string{
	Enabled:  "enabled",
	Disabled: "disabled",
}

func IsValidActivationStatus(status string) error {
	_, ok := activationStatus[ActivationStatus(status)]
	if !ok {
		return fmt.Errorf("%v is not valid enum of activationStatus", status)
	}

	return nil
}

func (a ActivationStatus) String() (string, error) {
	text, ok := activationStatus[a]
	if !ok {
		return "", fmt.Errorf("'%v' is not an valid enum of activationStatus", a)
	}

	return text, nil
}

func (a ActivationStatus) MustString() string {
	text, ok := activationStatus[a]
	if !ok {
		panic(fmt.Errorf("'%v' is not an valid enum of activationStatus", a))
	}

	return text
}
