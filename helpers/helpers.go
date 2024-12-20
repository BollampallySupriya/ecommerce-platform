package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type Envelope map[string] interface{}

type Message struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

var errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var MessageLogs = &Message{
	InfoLog: infoLog,
	ErrorLog: errorLog,
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err 
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value 
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err 
	}
	return nil 
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err 
	}
	err = dec.Decode(&struct{}{}) // TODO: eliminate if possible
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}
	return nil
}
