package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

var soft_version string
var blockedDomains []string

func main() {
	// Set initial variables
	soft_version = "0.0.1"
	blockedDomains = []string{"t1777799.com", "www.t1777799.com", "www.797ka.cn", "797ka.cn", "s1ndwdrld.logicdn.com"}

	log.Println("Starting brave proxy server on port 443")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	for _, domain := range blockedDomains {
		proxy.OnRequest(goproxy.ReqHostIs(domain)).HandleConnect(goproxy.AlwaysReject)
		proxy.OnRequest(goproxy.ReqHostIs(domain)).DoFunc(
			func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					"This domain is blocked")
			})
	}

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if r.Host == "brave-admin.cosmos-softwares.com" {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText,
					200,
					"Connected to Brave Proxy by Cosmos Softwares running on version "+soft_version,
				)
			}
			return r, nil
		})

	log.Fatal(http.ListenAndServe("0.0.0.0:443", proxy))
}
