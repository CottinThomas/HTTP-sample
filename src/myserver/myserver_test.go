package myserver

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
	"bytes"
)

func TestMyHandler(t *testing.T) {
	handler := &MyHandler{}
	server  := httptest.NewServer(handler)
	
	defer server.Close()

	for _, i := range []int{1, 2} {
		resp, err := http.Get(server.URL)
		
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}

		var expected bytes.Buffer
		actual, err := ioutil.ReadAll(resp.Body)
		
		if err != nil {
			t.Fatal(err)
		}
				
		temp, err := template.ParseFiles("./template.html")
		if err != nil {
			panic(err)
		}
		err = temp.ExecuteTemplate(&expected, "template.html", i)
		if err != nil {
			panic(err)
		}


		if expected.String() != string(actual) {
			t.Errorf("Expected the message '%s'\nGot '%s'", expected.String(),actual)
		}
	}
}