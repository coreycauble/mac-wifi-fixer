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

	net := getWifiInterface()

	urlList := []string{
		"https://google.com",
	}

	var timeInt = 5
	var restartCount int

	// consume command line args for configuration

	flag.IntVar(&timeInt, "i", 5, "interval in minutes")

	flag.Parse()

	if len(flag.Args()) > 0 {
		urlList = flag.Args()
	}

	log.Printf("Wifi connection being monitored.....\n\t\twith parameters. Interval: %v minutes URLS: %v", timeInt, urlList)
	for {
		time.Sleep(time.Duration(timeInt) * time.Minute)
		// saySomething("Please wait: I am checking your WiFi connections now.")
		for _, url := range urlList {
			if !checkHTTPStatus(url) {
				//s := fmt.Sprintf("Oh NO! Connection to %v failed!", url)
				//saySomething(s)
				log.Println("Restarting Wifi Connection")
				wifiControl("off", net)
				duration := time.Second * 5
				time.Sleep(duration)
				wifiControl("on", net)
				saySomething("Your Wifi has been restarted.")
				restartCount++
				break
			} else {
				//s := fmt.Sprintf("Connection to %v successful", url)
				//saySomething(s)
			}
		}
		log.Printf("Wifi Connection reset [%v] times", restartCount)
	}
}

func wifiControl(state string, net string) {
	cmdName := "networksetup"
	cmdArgs := []string{"-setairportpower", net, state}
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

func checkHTTPStatus(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("URL: ", url)
	log.Println("HTTP Response Status: " + strconv.Itoa(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Println("HTTP Status OK!")

	} else {
		log.Println("Argh! Broken")
	}
	return true
}

func getWifiInterface() string {
	cmd := "networksetup -listallhardwareports | fgrep Wi-Fi -A1 | awk 'NF==2{print $2}'"
	eth, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		log.Fatal("Can't detect wifi interface")
	}

	return strings.Trim(string(eth), " \n")
}
