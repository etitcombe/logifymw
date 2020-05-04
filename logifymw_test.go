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
	// "2020/05/04 13:34:12 GET  /                                                  3.2µs\n"
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d GET  \/.*s\n$`, buf.String())
	if err != nil {
		t.Fatalf("TestLogIt: error matching regular expression: %v", err)
	}
	if !ok {
		t.Errorf("log doesn't match %q", buf.String())
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
	// "2020/05/04 13:33:13 127.0.0.1:48994 GET  /                                                  Go-http-client/1.1 1.4µs\n"
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d 127.0.0.1:\d+ GET  \/\s+ Go-http-client\/1.1 .*s\n$`, buf.String())
	if err != nil {
		t.Fatalf("TestLogItMore: error matching regular expression: %v", err)
	}
	if !ok {
		t.Errorf("log doesn't match %q", buf.String())
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
	// "2020/05/04 13:33:13 127.0.0.1:44780 GET  /                                                  Go-http-client/1.1 200 5 6.5µs\n"
	ok, err := regexp.MatchString(`^\d\d\d\d\/\d\d\/\d\d \d\d:\d\d:\d\d 127.0.0.1:\d+ GET  \/\s+ Go-http-client\/1.1 200 5 .*s\n$`, buf.String())
	if err != nil {
		t.Fatalf("TestLogItMoreMore: error matching regular expression: %v", err)
	}
	if !ok {
		t.Errorf("log doesn't match %q", buf.String())
	}
}

func testHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "blorf")
	}
}
