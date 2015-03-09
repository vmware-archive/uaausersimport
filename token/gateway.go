package token

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

var dumpHttp = parseDumpHttp()

func parseDumpHttp() bool {
	debug, err := strconv.ParseBool(os.Getenv("DEBUG_HTTP"))
	if err != nil {
		return false
	}
	return debug
}

func RequestWithToken(token, url, method, contentType string, body io.Reader) (response *http.Response, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf(" Bearer %s", token))
	req.Header.Add("Content-Type", contentType)
	if dumpHttp {
		dump, _ := httputil.DumpRequest(req, true)
		fmt.Println(string(dump))
	}
	response, err = transport.RoundTrip(req)
	if err != nil {
		return
	}
	if dumpHttp {
		dump, _ := httputil.DumpResponse(response, true)
		fmt.Println(string(dump))
	}
	return
}
