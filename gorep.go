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
	dirname := os.Args[1]
	pattern := os.Args[2]
	fileinfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range fileinfos {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			//
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
					fmt.Printf("L%d: %s\n", line, scanner.Text())
				}
				line++
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
