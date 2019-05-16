package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	flags "github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type sesame struct {
	Locked bool   `json:"locked"`
	Error  string `json:"error"`
}

var opts struct {
	Id     *string `long:"id"     short:"d" description:"Device ID"`
	ApiKey *string `long:"apikey" short:"k" description:"API key"`
}

func getSesameLockeStatus(d string, k string) (bool, string) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://api.candyhouse.co/public/sesame/"+d, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", k)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var sesame sesame
	err = json.Unmarshal(body, &sesame)
	if err != nil {
		fmt.Println(err)
	}

	return sesame.Locked, sesame.Error
}

func run(args []string) *checkers.Checker {
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	checkSt := checkers.OK
	msg := fmt.Sprintf("Locked!")

	st, e := getSesameLockeStatus(*opts.Id, *opts.ApiKey)

	if st == true {
		checkSt = checkers.OK
		msg = fmt.Sprintf("Locked!")
	}
	if st == false && e == "" {
		checkSt = checkers.CRITICAL
		msg = fmt.Sprintf("Unocked!")
	}
	if st == false && e != "" {
		checkSt = checkers.WARNING
		msg = fmt.Sprintf(e)
	}

	return checkers.NewChecker(checkSt, msg)
}

func Do() {
	ckr := run(os.Args[1:])
	ckr.Name = "Sesame"
	ckr.Exit()
}

func main() {
	Do()
}
