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

func (n *Network) GetGames(ch chan<- GetGamesResponse) {
	response, err := http.Get(fmt.Sprintf("%s:%d/games", n.serverHost, n.serverPort))
	if err != nil {
		ch <- GetGamesResponse{
			Games: nil,
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- GetGamesResponse{
			Games: nil,
			Error: err,
		}
		return
	}

	getGames := GetGames{}
	if err := json.Unmarshal(body, &getGames); err != nil {
		ch <- GetGamesResponse{
			Games: nil,
			Error: err,
		}
		return
	}

	ch <- GetGamesResponse{
		Games: getGames.Games,
		Error: nil,
	}
}
