package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/chenjia404/guardnode/update"
)

var logger = log.New(os.Stderr, "httpsproxy:", log.Llongfile|log.LstdFlags)
var (
	version   = "0.0.3"
	gitRev    = ""
	buildTime = ""
)

var ForwardHost string

func main() {
	var listenAdress string
	var flag_update bool
	flag.StringVar(&listenAdress, "l", "0.0.0.0:18080", "listen address.eg: 127.0.0.1:18080")
	flag.StringVar(&ForwardHost, "f", "", "Forward to parent website eg:https://google.com")
	flag.BoolVar(&flag_update, "update", false, "update form github")
	flag.Parse()

	if flag_update {
		update.CheckGithubVersion(version)
		return
	}

	server := &http.Server{
		Addr: listenAdress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Serve(w, r)
		}),
	}
	server.ListenAndServe()
}

func Serve(w http.ResponseWriter, r *http.Request) {
	handleHttp(w, r)

}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	var u *url.URL
	if len(ForwardHost) == 0 {
		u, _ = url.Parse("https://" + r.Header.Get("o-host"))
	} else {
		u, _ = url.Parse(ForwardHost)
	}
	r.Host = u.Host
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true, // 忽略证书验证
	}
	var transport http.RoundTripper = &http.Transport{
		Proxy: nil, // 不使用代理，如果想使用系统代理，请使用 http.ProxyFromEnvironment
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = transport
	proxy.ServeHTTP(w, r)

}
