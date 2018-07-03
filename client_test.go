package pingidentityapi

import (
	"crypto/tls"
	"net/http"
	"github.com/dnaeon/go-vcr/recorder"
	"testing"
)


func TestGet(t *testing.T) {
	r := createAccessRecorder(t)
	defer r.Stop()
	client := createClient(r)
	resp, err := client.Get("adminConfig")
	if err != nil {
		t.Fatalf("failed: %s", err)	
	} 
	if resp["hostPort"] != "localhost:9090" {
		t.Fatalf("Unexpected hostport: %v", resp["hostPort"])
	}
	t.Logf("output: %v", resp)
}

func TestPost(t *testing.T) {
	r, client := initTest(t)
	defer r.Stop()
	var m = map[string]interface{} {
		"host": "test",
		"port": 3000,
	}
	
	resp, err := client.Post("virtualhosts", m)
	if err != nil {
		t.Fatal(err)
	}

	if resp["host"].(string) != "test" || resp["port"].(float64) != 3000 {
		t.Fatalf("Test failed with response data: %v", resp)
	}
	
}

func TestPut(t *testing.T) {
	r, client := initTest(t)
	defer r.Stop()
	var m = map[string]interface{} {
		"host": 	"test",
		"port": 	4000,
	}
	
	resp, err := client.Put("virtualhosts/" + "3", m)
	if err != nil {
		t.Fatal(err)
	}

	if resp["host"].(string) != "test" || resp["port"].(float64) != 4000 {
		t.Fatalf("Test failed with response data: %v", resp)
	}
	
}

func TestDelete(t *testing.T) {
	r, client := initTest(t)
	defer r.Stop()
	
	resp, err := client.Delete("virtualhosts/" + "3")
	if err != nil {
		t.Fatal(err)
	}

	if resp["msg"].(string) != "Operation successful." {
		t.Fatalf("Test failed with response data: %v", resp)
	}
}

func createAccessRecorder(t *testing.T) *recorder.Recorder {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	r, err := recorder.New("fixtures/access-api-" + t.Name())
	if err != nil {
		t.Fatalf("Failed to initialize recorder. %v", err)
	}
	return r
}

func createClient(r http.RoundTripper) *Client {
	return NewClient(&Configuration{"https://192.168.33.111:9000/pa-admin-api/v3/", "Administrator", "Testpassword1", "pf", r})
}

func initTest(t *testing.T) (*recorder.Recorder, *Client) {
	r := createAccessRecorder(t)
	return r, createClient(r)
}