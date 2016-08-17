package main

import (
    "crypto/md5"
    "fmt"
)

func main() {
    data := []byte("Hello World!")
    fmt.Println(md5.Sum(data))
    fmt.Printf("%x\n", md5.Sum(data))
}
