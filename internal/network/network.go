package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Network struct {
	serverHost string
	serverPort uint16
}

func New(serverHost string, serverPort uint16) *Network {
	return &Network{
		serverHost: serverHost,
		serverPort: serverPort,
	}
}

func (n *Network) GetGames() ([]string, error) {
	response, err := http.Get(fmt.Sprintf("%s:%d/games", n.serverHost, n.serverPort))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	getGames := GetGames{}
	if err := json.Unmarshal(body, &getGames); err != nil {
		return nil, err
	}

	return getGames.Games, nil
}
