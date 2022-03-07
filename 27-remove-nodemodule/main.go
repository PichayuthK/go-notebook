package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	p := "/path/right/here"

	ts, err := getDirFromPath(p)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		if len(ts) == 0 {
			break
		}
		current := ts[0]
		ts = ts[1:]

		if strings.Contains(current, "node_modules") {
			fmt.Println(current)
			err := os.RemoveAll(current)
			if err != nil {
				log.Fatalln(err)
			}
			continue
		}

		tss, err := getDirFromPath(current)
		if err != nil {
			log.Fatalln(err)
		}

		ts = append(ts, tss...)

	}

}

func getDirFromPath(p string) ([]string, error) {
	ds, err := ioutil.ReadDir(p)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var ts []string
	for _, d := range ds {
		if d.IsDir() {
			ts = append(ts, p+"/"+d.Name())
		}
	}

	return ts, nil
}
