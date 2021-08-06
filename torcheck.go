
package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"net/url"
	"crypto/tls"
)

func main() {
	torified()
	directrequest() 
}

func directrequest() {
		client := &http.Client{}
		resp, err := client.Get("https://api.myip.com") 
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Println("Direct Connection IP Details : ", bodyString)
		}
	
}

func torified() {
	proxyurl, _ := url.Parse("socks5://127.0.0.1:9050")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyurl), TLSClientConfig:&tls.Config{InsecureSkipVerify: true }}}
	resp, err := client.Get("https://api.myip.com") 
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("Torified IP Details : ", bodyString)
	}
}