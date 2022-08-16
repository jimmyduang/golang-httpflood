package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Socks5Lister interface {
	List() ([]string, error)
}

// FileSocks5List implements Socks5Lister interface
// open file get list of all socks5 addr
type FileSocks5List struct {
	Path string
}

func (f FileSocks5List) List() ([]string, error) {
	res := make([]string, 0, 100)
	if len(f.Path) < 1 {
		return res, fmt.Errorf("path is empty")
	}
	fi, err := os.Open(f.Path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		res = append(res, string(a))
	}
	return res, nil
}

// ApiSocks5List implements Socks5Lister interface
// get list of all socks5 addr by http request
type ApiSocks5List struct {
	ApiServerAddress string
	Id               string
	Size             int
	Schemes          string
	SupportHTTPS     string
	RestimeWithinMs  int
	Format           string
}

func (f ApiSocks5List) List() ([]string, error) {
	if f.ApiServerAddress == "" {
		return nil, fmt.Errorf("ApiServerAddress is empty")
	}
	req, err := http.NewRequest("GET", f.ApiServerAddress+"get_proxies", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("id", f.Id)
	q.Add("size", strconv.Itoa(f.Size))
	q.Add("schemes", f.Schemes)
	q.Add("support_https", f.SupportHTTPS)
	q.Add("restime_within_ms", strconv.Itoa(f.RestimeWithinMs))
	q.Add("format", f.Format)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(string(body), "success") {
		return nil, fmt.Errorf("req error: %v", string(body))
	}
	return strings.Split(string(body), "\r\n"), nil
}
