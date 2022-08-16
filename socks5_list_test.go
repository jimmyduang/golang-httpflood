package main

import (
	"strings"
	"testing"
)

func TestFileSocks5List_List(t *testing.T) {
	t.Run("fileSocks5List", func(t *testing.T) {
		f := FileSocks5List{Path: "./socks5.txt"}
		list, err := f.List()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		if len(list) == 0 {
			t.Errorf("list is empty")
			return
		}
		for _, l := range list {
			p := &tcpProxySocket{}

			dialer, err := p.connProxy(strings.Trim(l, "\r"))
			if err != nil {
				t.Errorf("url2.Parse failed: %v", err)
				continue
			}
			_, err = dialer.Dial("tcp", "localhost:8001")
			if err != nil {
				t.Logf("Dial failed: %v", err)
			}
		}
	})
}

func TestApiSocks5List_List(t *testing.T) {
	t.Run("apiSocks5ListTest", func(t *testing.T) {
		apiList := ApiSocks5List{
			ApiServerAddress: "http://xxxxxxxx.com/api/",
			Id:               "xxxxxxxx",
			Size:             50,
			Schemes:          "socks5",
			SupportHTTPS:     "false",
			RestimeWithinMs:  5000,
			Format:           "txt2_1",
		}
		list, err := apiList.List()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		if len(list) == 0 {
			t.Errorf("list is empty")
			return
		}
		for _, l := range list {
			p := &tcpProxySocket{}
			dialer, err := p.connProxy(strings.Trim(l, "\r"))
			if err != nil {
				t.Errorf("url2.Parse failed: %v", err)
				continue
			}
			_, err = dialer.Dial("tcp", "localhost:8001")
			if err != nil {
				t.Logf("Dial failed: %v", err)
			} else {
				t.Logf("Connect success: %v", l)
			}
		}
	})
}
