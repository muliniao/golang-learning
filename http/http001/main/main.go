package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := Validate("abcd1234", r)
		if err != nil {
			fmt.Errorf("")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("not match"))
			return
		}

		fmt.Println("aaaaaaaa")
		fmt.Fprintf(w, "hello world")

	})

	http.ListenAndServe("localhost:8080", mux)
}

func Validate(secret string, r *http.Request) error {
	// sha1=9705cb0b8a78ec21eef6614f849a2f72ff2161b9

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) == 0 {
		return errors.New("wrong payload")
	}

	if len(secret) > 0 {
		signature := r.Header.Get("X-Hub-Signature")
		if len(signature) == 0 {
			return errors.New("miss signature")
		}
		mac := hmac.New(sha1.New, []byte(secret))
		_, _ = mac.Write(payload)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
			return errors.New("not match")
		}
	}

	return nil
}
