package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type logWriter struct {}

func main() {
	resp, err := http.Get("https://olx.pt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//One way
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	//One way
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//AnotherWay
	//io.Copy(os.Stdout, resp.Body)
}

//Custom implementation of the writer function
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Num of bytes: ", len(bs))
	return 1, nil
}