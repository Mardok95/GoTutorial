package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64
	}
	/*
		type location struct {
			Lat float64 'json:"latitudine"'
			Long float64 'json:"longitudine"'
		}
		In this way I add a tag to the JSON name for interoperability issue (python or ruby)

	*/
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
