package main

import (
	"fmt"
	"strings"
)

func urlProcess(url string) string{
	result := strings.HasPrefix(url, "http://")
	if !result  {
		url = "http://" + url
	    
	}
	fmt.Println(url)	
	return url
}
func pathProcess(path string) string{
	result := strings.HasSuffix(path, "/")
	if result == false {
		path = path + "/"
	}
	return path
}

func main() {
	var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)
    fmt.Println(url,path)

}