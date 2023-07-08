//simple web crawler

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	//get the url from the command line
	url := os.Args[1]
	//get the html from the url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//read the html
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//convert the html to string
	s := string(body[:])
	//get the links from the html
	links := getLinks(s)
	//print the links
	for _, link := range links {
		fmt.Println(link)
	}
}

//getLinks returns a slice of strings containing the links in the html

func getLinks(s string) []string {
	//slice to store the links
	var links []string
	//loop through the html
	for i := 0; i < len(s); i++ {
		//if the current character is a '<' then the next character is a link
		if s[i] == '<' {
			//get the link
			link := getLink(s[i+1:])
			//if the link is not empty then append it to the slice
			if link != "" {
				links = append(links, link)
			}
		}
	}
	return links
}

//getLink returns a string containing the link in the html

func getLink(s string) string {
	//slice to store the link
	var link []byte
	//loop through the html
	for i := 0; i < len(s); i++ {
		//if the current character is a '>' then the link is over
		if s[i] == '>' {
			//convert the link to string and return it
			return string(link)
		}
		//append the current character to the link
		link = append(link, s[i])
	}
	return ""
}
