// Copyright (c) 2016 Henry Slawniak <https://henry.computer/>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"net/http"
)

// ErrorMessages cotnains our general error messages for http error codes
var ErrorMessages map[int]string

// APIErrorMessages contains our http error messages for the API interface
var APIErrorMessages map[int]string

func init() {
	ErrorMessages = map[int]string{
		http.StatusNotFound:            "We can't find that page :(",
		http.StatusUnauthorized:        "You are not authorized to do that",
		http.StatusInternalServerError: "Something is going horribly wrong!",
		http.StatusBadRequest:          "Your request is malformed",
		http.StatusNotImplemented:      "We haven't finished this feature yet",
	}
}

func httpError(code int, w http.ResponseWriter, r *http.Request, err error) error {
	w.WriteHeader(code)
	message := ""
	if _, ok := ErrorMessages[code]; ok {
		message = ErrorMessages[code]
	}
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%d: %s (%s)", code, message, err.Error())))
	} else {
		w.Write([]byte(fmt.Sprintf("%d: %s", code, message)))
	}
	return nil
}

func notFound(w http.ResponseWriter, r *http.Request) error {
	return httpError(http.StatusNotFound, w, r, nil)
}

func notAuthed(w http.ResponseWriter, r *http.Request, err error) error {
	return httpError(http.StatusUnauthorized, w, r, err)
}

func internalError(w http.ResponseWriter, r *http.Request, err error) error {
	return httpError(http.StatusInternalServerError, w, r, err)
}

func badRequest(w http.ResponseWriter, r *http.Request, err error) error {
	return httpError(http.StatusBadRequest, w, r, err)
}

func forbidden(w http.ResponseWriter, r *http.Request, err error) error {
	return httpError(http.StatusForbidden, w, r, err)
}
