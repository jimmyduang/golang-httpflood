package main

import (
	"fmt"
	"net/http"
	"testing"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	go func() { _ = http.ListenAndServe(":8001", nil) }()
}
func Test_tcpProxySocket_connProxy(t *testing.T) {
	t.Run("http test", func(t *testing.T) {
		p := &tcpProxySocket{}
		dialer, err := p.connProxy("socks5://119.23.71.164:8880")
		if err != nil {
			t.Errorf("url2.Parse failed: %v", err)
			return
		}
		_, err = dialer.Dial("tcp", "localhost:8001")
		if err != nil {
			t.Errorf("url2.Dial failed: %v", err)
		}
	})
}
