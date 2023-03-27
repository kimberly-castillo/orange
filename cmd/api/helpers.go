// Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) readIDParams(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// convert the data to a JSON format
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	// Add any headers that were sent
	for key, value := range headers {
		w.Header()[key] = value
	}
	// add header information
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(js))
	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, destination interface{}) error { //interface{} meaning of any type
	//specify the max size of our JSON request body
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	// we're trying to decode the json request
	// read json into destination
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() //if we put in something that doesnt exist in our thing, then an error should be generated
	// bc destination is a pointer, we need to make sure to pass it as a pointer

	//start the decoding
	err := dec.Decode(destination)
	if err != nil {
		// something went wrong
		// deciding the type of error - can test for different classes of errors using json
		var syntaxError *json.SyntaxError
		var unMarshalTypeError *json.UnmarshalTypeError       // something went wrong when they were converting json to text
		var invalidUnmarshalError *json.InvalidUnmarshalError // when you didnt pass a pointer which is invalid
		var maxBytesError *http.MaxBytesError

		//let's check for the type of decode error
		switch {
		case errors.As(err, &syntaxError): //AS is for looking for type of error
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset) //offset tell you where
		case errors.Is(err, io.ErrUnexpectedEOF): // IS for specific error
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unMarshalTypeError):
			if unMarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q)", unMarshalTypeError.Field)
			}
			return fmt.Errorf("body contains badly-formed JSON (at character%d)", unMarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
			//unmappable field/non-existent fieldl
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)
		case errors.As(err, &invalidUnmarshalError):
			panic(err) //kill the program
		default:
			return err
		}
	}
	//Lets call the decoder again to check if there are any trailing json objects
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain a single JSON value")
	}
	return nil

}
