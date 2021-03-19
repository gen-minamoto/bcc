package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type bank struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kana string `json:"kana"`
	Hira string `json:"hira"`
	Roma string `json:"roma"`
}

type banks map[string]bank

const dataURL = "https://raw.githubusercontent.com/zengin-code/source-data/master/data/banks.json"
const jsonFileName = "banks.json"

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("specify bank code")
	}

	var r io.Reader
	var err error

	r, err = os.Open(filepath.Join(os.TempDir(), jsonFileName))

	if err != nil {
		req, err := http.NewRequest(http.MethodGet, dataURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		f, err := os.Create(filepath.Join(os.TempDir(), jsonFileName))
		if err != nil {
			log.Fatal(err)
		}
		r = io.TeeReader(res.Body, f)
	}

	var bs banks
	if err := json.NewDecoder(r).Decode(&bs); err != nil {
		log.Fatal(err)
	}

	codes := strings.Split(flag.Arg(0), ",")

	for _, code := range codes {
		v := bs[code]
		fmt.Printf("%s,%s,%s,%s,%s\n", code, v.Name, v.Kana, v.Hira, v.Roma)
	}
}
