package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	urlList := []string{
		"https://google.com",
	}
	var timeInt int = 5

	// consume command line args for configuration

	flag.IntVar(&timeInt, "i", 5, "interval in minutes")

	flag.Parse()

	if len(flag.Args()) > 0 {
		urlList = flag.Args()
	}

	log.Printf("Wifi connection being monitored.....\n\t\twith parameters. Interval: %v URLS: %v", timeInt, urlList)
	for {
		time.Sleep(time.Duration(timeInt) * time.Minute)
		// saySomething("Please wait: I am checking your WiFi connections now.")
		for _, url := range urlList {
			if !checkHttpStatus(url) {
				//s := fmt.Sprintf("Oh NO! Connection to %v failed!", url)
				//saySomething(s)
				log.Println("Restarting Wifi Connection")
				wifiControl("off")
				duration := time.Second * 5
				time.Sleep(duration)
				wifiControl("on")
				saySomething("Your Wifi has been restarted.")
				break
			} else {
				//s := fmt.Sprintf("Connection to %v successful", url)
				//saySomething(s)
			}
		}
	}
}

func wifiControl(state string) {
	cmdName := "networksetup"
	cmdArgs := []string{"-setairportpower", getWifiInterface(), state}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running : ", err)
		os.Exit(1)
	}
}

func saySomething(speach string) {
	cmdName := "say"
	cmdArgs := []string{speach}
	if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running : ", err)
		os.Exit(1)
	}
}

func checkHttpStatus(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("URL: ", url)
	log.Println("HTTP Response Status: " + strconv.Itoa(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Println("HTTP Status OK!")

		return true
	} else {
		log.Println("Argh! Broken")
		return true
	}
}

func getWifiInterface() string {
	cmd := "networksetup -listallhardwareports | fgrep Wi-Fi -A1 | awk 'NF==2{print $2}'"
	eth, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		log.Fatal("Can't detect wifi interface")
	}

	return strings.Trim(string(eth), " \n")
}
