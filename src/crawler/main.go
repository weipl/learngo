package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

//mian_init_page
/*package main

import (
"bufio"
"fmt"
"golang.org/x/net/html/charset"
"golang.org/x/text/encoding"
"golang.org/x/text/transform"
"io"
"io/ioutil"
"net/http"
)

func main() {

	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: response code",
			resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body,
		e.NewDecoder())

	all, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",all)

}
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	encoding, _, _ := charset.DetermineEncoding(bytes,"")
	return encoding
}*/

//技术总结
//加入了第三方库simplifiedchinese
//1、(1)http.get()	传入指定url，返回Response
//if 自定义头文件 http.NewRequest() 配合 DefaultClient.Do.
// (2)http.StatusOK 判断resp.StatusCode 的错误码
//(3)ioutil.ReadALl 从resp.boody读取数据
//2、如果网页时GBK 加上这段
// utf8reader := transform.NewReader(resp.Body,
//		simplifiedchinese.GBK.NewDecoder())

//		ioutil.ReadAll() resp.body 换成utf8reader
//3.可能有些网页不是GBK编码的，那就需要自动识别网页中编码方式
//charset.DetermineEncoding 来获取网页中编码
//-----------------------------2019/3/22
