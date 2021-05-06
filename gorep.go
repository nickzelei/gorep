package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	pattern := os.Args[1]
	dirname := os.Args[2]
	fileinfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range fileinfos {
		if !file.IsDir() {
			f, err := os.Open(file.Name())
			if err != nil {
				log.Fatal(err)
				return
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			line := 1
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), pattern) {
					fmt.Printf("%s:%d: %s\n", file.Name(), line, scanner.Text())
				}
				line++
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
