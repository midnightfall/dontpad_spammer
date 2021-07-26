package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/89z/mech"
)

const file = "domains.txt"
const message = `
░░░░▄▄▄▄▀▀▀▀▀▀▀▀▄▄▄▄▄▄
░░░░█░░░░▒▒▒▒▒▒▒▒▒▒▒▒░░▀▀▄
░░░█░░░▒▒▒▒▒▒░░░░░░░░▒▒▒░░█
░░█░░░░░░▄██▀▄▄░░░░░▄▄▄░░░█
░▀▒▄▄▄▒░█▀▀▀▀▄▄█░░░██▄▄█░░░█
█▒█▒▄░▀▄▄▄▀░░░░░░░░█░░░▒▒▒▒▒█
█▒█░█▀▄▄░░░░░█▀░░░░▀▄░░▄▀▀▀▄▒█
░█▀▄░█▄░█▀▄▄░▀░▀▀░▄▄▀░░░░█░░█
░░█░░▀▄▀█▄▄░█▀▀▀▄▄▄▄▀▀█▀██░█
░░░█░░██░░▀█▄▄▄█▄▄█▄████░█
░░░░█░░░▀▀▄░█░░░█░███████░█
░░░░░▀▄░░░▀▀▄▄▄█▄█▄█▄█▄▀░░█
░░░░░░░▀▄▄░▒▒▒▒░░░░░░░░░░█
░░░░░░░░░░▀▀▄▄░▒▒▒▒▒▒▒▒▒▒░█
░░░░░░░░░░░░░░▀▄▄▄▄▄░░░░░█ github.com/midnightfall`

func main() {
	sda := text(file)
	troll(sda, message)

}

func text(text_file string) []string {
	domains := []string{}
	file, _ := os.Open(text_file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	return domains
}

func troll(domains []string, word string) {
	for _, v := range domains {
		the_url := "http://dontpad.com/" + v
		fmt.Println("====================================================================================")
		fmt.Printf("Finding Subdomains of %v \n", the_url)
		fmt.Println("====================================================================================")
		domains2 := []string{}
		value := ""
		resp, _ := http.Get(the_url)
		defer resp.Body.Close()
		doc, _ := mech.Parse(resp.Body)
		input := doc.ByTag("input")
		for input.Scan() {
			seila := input.Attr("value")
			final := string(seila)
			if final != "" {
				value = final
			}
		}
		complex_url := the_url + "/.menu.json?_=" + string(value)
		resp2, _ := http.Get(complex_url)
		body, _ := ioutil.ReadAll(resp2.Body)
		fuckme := string(body)
		ress := strings.Split(fuckme, `"`)
		for _, rlly := range ress {
			if rlly != "[" {
				if rlly != "," {
					if rlly != "]" {
						domains2 = append(domains2, rlly)
					}
				}
			}
		}
		if len(domains2) == 1 {
			fmt.Println("0 Subdomains Founded!")
		} else {
			for _, subdomain := range domains2 {
				subdomain_url := the_url + "/" + subdomain
				resp, error := http.PostForm(subdomain_url, url.Values{"text": {word}})
				if resp.StatusCode == 200 {
					fmt.Printf("\"%v\" was written in %v subdomain!\n", word, subdomain_url)
				} else {
					fmt.Printf("Something wrong happened: (%v)", error)
				}
			}
		}
		fmt.Println("====================================================================================")
		fmt.Println("Subdomains are done, lets go to main domain...")
		fmt.Println("====================================================================================")
		resp, error := http.PostForm(the_url, url.Values{"text": {word}})
		if resp.StatusCode == 200 {
			fmt.Printf("\"%v\" was written in %v main domain!\n", word, the_url)
		} else {
			fmt.Printf("Something wrong happened: (%v)", error)
		}
	}
	fmt.Println("\n")
	fmt.Println("Press any key to exit.")
	fmt.Scanln()
}
