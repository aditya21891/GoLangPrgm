package main
import (
	"io/ioutil"
	"fmt"
	"log"
)
func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
    }
}
