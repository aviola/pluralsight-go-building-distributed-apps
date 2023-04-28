package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)

	// enc will allow us to serialize the registration in JSON and set the buffer to the destination
	enc := json.NewEncoder(buf)

	if err := enc.Encode(r); err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code %v", res.StatusCode)
	}

	return nil
}
