package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readf(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func savef(text, path string) {
	f, err := os.Create(path)
	checkErr(err)

	defer f.Close()

	n3, err := f.WriteString(text)
	checkErr(err)
	fmt.Printf("wrote %d bytes => %v\n", n3, path)
	f.Sync()
}
