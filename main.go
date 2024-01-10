package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var lightUrl string

func main() {
	url := flag.String("url", "192.168.1.199", "Light URL")
	turn := flag.String("turn", "", "Turn the light on or off")
	mode := flag.String("mode", "", "Change the light mode from white to color")
	color := flag.String("color", "", "Chose a color R/G/Y")

	flag.Parse()

	fmt.Println("url:", *url)
	lightUrl = *url
	fmt.Println("turn:", *turn)
	fmt.Println("mode:", *mode)
	fmt.Println("color:", *color)
	handleParams(*turn, *mode, *color)
}

func handleParams(turn, mode, color string) {
	var params string

	if turn != "" {
		params = "?turn=" + turn
	} else if mode != "" {
		params = "?mode=" + mode
	} else if color != "" {
		const Red = "R"
		const Green = "G"
		const Yellow = "Y"
		switch color {
		case Red:
			params = "?red=255&green=0&blue=0"
		case Green:
			params = "?green=255&red=0&blue=0"
		case Yellow:
			params = "?red=255&green=255&blue=0"
		}
	} else {
		params = ""
	}

	configBulb(params)
	getBulbStatus(params)
}

type Response struct {
	IsOn       bool   `json:"ison"`
	Source     string `json:"source"`
	HasTimer   bool   `json:"has_timer"`
	Mode       string `json:"mode"`
	Red        int    `json:"red"`
	Green      int    `json:"green"`
	Blue       int    `json:"blue"`
	White      int    `json:"white"`
	Gain       int    `json:"gain"`
	Temp       int    `json:"temp"`
	Brightness int    `json:"brightness"`
	Effect     int    `json:"effect"`
	Transition int    `json:"transition"`
}

func getBulbStatus(config string) Response {
	var responseObject Response
	req, err := http.NewRequest("GET", "http://"+lightUrl+"/light/0", nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "ShellyMgr/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("Status Called")
	fmt.Println(responseObject)
	return responseObject
}

func configBulb(params string) Response {
	var responseObject Response
	fmt.Println("Params Printed")
	fmt.Println("http://" + lightUrl + "/light/0" + params)
	req, err := http.NewRequest("POST", "http://"+lightUrl+"/light/0"+params, nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "ShellyMgr/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &responseObject)
	return responseObject
}
