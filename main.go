package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	confEndpoint := "http://localhost/v2/snaps/system/conf"

	body :=
		`{"system": { "kernel": { "dangerous-cmdline-append": "buz=bazz foo=bar" } } }`

	payload := bytes.NewBufferString(body)

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/run/snapd.socket")
			},
		},
	}

	req, err := http.NewRequest("PUT", confEndpoint, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Dump the response
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Error dumping response:", err)
		return
	}
	fmt.Printf("HTTP RESPONSE:\n%v", string(dump))

}
