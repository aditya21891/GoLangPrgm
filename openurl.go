package main

import (
	"fmt"
	"pkg/browser"
)

func main() {
	urlopen := "https://medium.com"
	//code to open the URL
	browser.OpenURL(urlopen)
	fmt.Println("My Favorite Togo Blog")

}
