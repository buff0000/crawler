package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"time"
)

func httpclientDef(url string) (resp *http.Response, err error){
	return http.Get(url)
}

func httpclient(url string) (resp *http.Response, err error){
	client := &http.Client{}
	req, err := http.NewRequest("GET",url, nil)
	if err != nil{
		log.Fatal(err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	return client.Do(req)
}

func Fetcher(url string) ([]byte, error){
	time.Sleep(time.Second)
	resp, err := httpclient(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		fmt.Println("Error: staut code", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}