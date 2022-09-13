package main

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func gameInit(client http.Client, gameMetadata gameMetadata, addr string)  {
	postBody, _ := json.Marshal(gameMetadata)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := client.Post(addr+"/game_metadata", "application/json",responseBody)

	if err != nil {
		fmt.Printf("init err: %s\n", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func gameRemove(client http.Client, gameMetadata gameMetadata, addr string) {
	postBody, _ := json.Marshal(gameMetadata)

	responseBody := bytes.NewBuffer(postBody)
	resp,err := client.Post(addr+"/remove_game", "application/json",responseBody)

	if err != nil {
		fmt.Printf("init err: %s\n", err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func registerGameEvent(client http.Client, register GameEventRegister, addr string)  {
	postBody, _ := json.Marshal(register)

	responseBody := bytes.NewBuffer(postBody)

	_, err := client.Post(addr+"/register_game_event", "application/json",responseBody)

	if err != nil {
		fmt.Printf("register game event err: %s\n", err)
		return
	}
}

func bindGameEvent(client http.Client, registrable Registrable, addr string)  {
	postBody, _ := json.Marshal(registrable)
	responseBody := bytes.NewBuffer(postBody)

	_, err := client.Post(addr+"/bind_game_event", "application/json",responseBody)

	if err != nil {
		fmt.Printf("bind err: %s\n", err)
		return
	}
}

func heartbeat(client http.Client, addr string, debug bool) {
	for  {
		postBody, _ := json.Marshal(map[string]string{
			"game": GameName,
		})
		responseBody := bytes.NewBuffer(postBody)

		resp, err := client.Post(addr+"/game_heartbeat", "application/json",responseBody)
		if err != nil {
			fmt.Printf("Err: %s\n", err)
			return
		}

		time.Sleep(time.Second*4)

		if !debug {
			resp.Body.Close()
			continue
		}

		//Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)

		resp.Body.Close()
	}
}

func registerPeriodicEvent(client http.Client, addr string, f func() int, eventName string, duration time.Duration){
	for {
		event := BuildEvent().BuildEventName(eventName).BuildGameName(GameName).BuildValue(f()).Get()

		postBody, _ := json.Marshal(event)
		responseBody := bytes.NewBuffer(postBody)

		_, err := client.Post(addr+"/game_event", "application/json",responseBody)
		if err != nil {
			fmt.Printf("Register periodic %s err: %s\n", eventName, err)
			return
		}
		time.Sleep(duration)
	}
}

func postEvent(client http.Client, addr string, event GameEvent)  {
	postBody, _ := json.Marshal(event)
	responseBody := bytes.NewBuffer(postBody)

	_, err := client.Post(addr+"/game_event", "application/json",responseBody)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
}

func ReadFile(path string) Config {

	b, err := os.ReadFile(path)
	if err != nil {
		println(err.Error())
		return Config{}
	}
	c := Config{}

	err = json.Unmarshal(b, &c)
	if err != nil {
		println(err.Error())
		return Config{}
	}
	return c
}