package main

import (
    "bufio"
    "fmt"
	"net/http"
	"strings"
)

func main() {

	body :=  strings.NewReader("hello")


    resp, err := http.Post("http://localhost:8080/messages/send", "application/text", body)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}