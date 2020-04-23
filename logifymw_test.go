package logifymw_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/etitcombe/logifymw"
)

func TestLogIt(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	ts := httptest.NewServer(logifymw.LogIt(testHandler()))
	defer ts.Close()

	_, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	// We should get a string like this:
	// 2019/11/22 14:18:05 GET  /                                                  0s
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d GET  \/\s+`, buf.String())
	if !ok {
		t.Error("log doesn't match", buf.String(), err)
	}
}

func TestLogItMore(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	ts := httptest.NewServer(logifymw.LogItMore(testHandler()))
	defer ts.Close()

	//_, err := http.Get(ts.URL + "?id=22")
	_, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	// We should get a string like this:
	// 2019/11/22 14:18:05 127.0.0.1:62643 GET  /                                                  0s Go-http-client/1.1
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d 127.0.0.1:\d+ GET  \/\s+ Go-http-client\/1.1`, buf.String())
	if !ok {
		t.Error("log doesn't match", buf.String(), err)
	}
}

func TestLogItMoreMore(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	ts := httptest.NewServer(logifymw.LogItMoreMore(testHandler()))
	defer ts.Close()

	_, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	// We should get a string like this:
	// 2020/04/23 14:17:04 127.0.0.1:58790 GET  /                                                  Go-http-client/1.1 200 5 0s
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d 127.0.0.1:\d+ GET  \/\s+ Go-http-client\/1.1 200 5 0s`, buf.String())
	if !ok {
		t.Error("log doesn't match", buf.String(), err)
	}
}

func testHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "blorf")
	}
}
