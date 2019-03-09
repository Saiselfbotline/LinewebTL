package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"os"
	"log"
	"fmt"
)
// Config 
const (
	host        = "https://access.line.me"
	qrlogin_url = host + "/qrlogin"
	qrwait_url  = qrlogin_url + "/v1/qr/wait"
	pinwait_url = qrlogin_url + "/v1/pin/wait" 
)

func main() {
	// unixtime
	t := time.Now()
	unixtime := t.Unix()
	// Generate qrCode
	req, err := http.NewRequest("GET", qrlogin_url + fmt.Sprintf("/v1/session?_=%v&channelId=1341209950", unixtime), nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) like Gecko")
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	fmt.Printf("%s %s\n", resp.Proto, resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}