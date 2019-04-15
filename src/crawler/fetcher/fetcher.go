package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong error status code %s",
				resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8reader := transform.NewReader(bodyReader,
		e.NewDecoder())
	http.Get("http://www.zhenai.com/zhenghun")

	return ioutil.ReadAll(utf8reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetcher error %v\n", err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}
