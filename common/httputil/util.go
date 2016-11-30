package httputil

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type SimpleResponse struct {
	Response string
}

func WriteSimpleResponse(w http.ResponseWriter, text string) error {
	return WriteJson(w, SimpleResponse{Response: text})
}

func WriteJson(w http.ResponseWriter, response interface{}) error {
	data, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func ReadRequest(r *http.Request, into interface{}) error {
	return ReadReader(&r.Body, into)
}

func ReadReader(r *io.ReadCloser, into interface{}) error {
	body, err := ioutil.ReadAll(*r)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, into); err != nil {
		return err
	}
	return nil
}
