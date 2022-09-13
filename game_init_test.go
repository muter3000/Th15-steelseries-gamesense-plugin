package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGameRemove(t *testing.T) {
	client := http.Client{}
	config := ReadFile("C:\\ProgramData\\SteelSeries\\SteelSeries Engine 3\\coreProps.json")
	addr := fmt.Sprintf("http://%s",config.Address)
	gameRemove(client, gameMetadata{Game: GameName,
	}, addr)
}
