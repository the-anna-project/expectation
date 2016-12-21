package expectation

import (
	"encoding/json"
)

func (e *expectation) MarshalJSON() ([]byte, error) {
	type ExpectationClone expectation

	b, err := json.Marshal(&struct {
		*ExpectationClone
	}{
		ExpectationClone: (*ExpectationClone)(e),
	})
	if err != nil {
		return nil, maskAny(err)
	}

	return b, nil
}

func (e *expectation) UnmarshalJSON(b []byte) error {
	type ExpectationClone expectation

	aux := &struct {
		*ExpectationClone
	}{
		ExpectationClone: (*ExpectationClone)(e),
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
