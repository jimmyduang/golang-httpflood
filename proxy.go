package main

import (
	"golang.org/x/net/proxy"
	url2 "net/url"
)

//type tcpProxy interface {
//	connProxy(string) (proxy.Dialer, error)
//}
type tcpProxySocket struct{}

// connProxy only support socks5
func (tcpProxySocket) connProxy(url string) (proxy.Dialer, error) {
	proxyUrl, err := url2.Parse(url)
	if err != nil {
		return nil, err
	}
	return proxy.FromURL(proxyUrl, nil)
}
