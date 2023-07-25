package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Network struct {
	serverHost string
}

func New(serverHost string) *Network {
	return &Network{
		serverHost: serverHost,
	}
}

func (n *Network) GetGames(ch chan<- GetGamesResponse) {
	response, err := http.Get(fmt.Sprintf("%s/games", n.serverHost))
	if err != nil {
		ch <- GetGamesResponse{
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- GetGamesResponse{
			Error: err,
		}
		return
	}

	getGames := GetGames{}
	if err := json.Unmarshal(body, &getGames); err != nil {
		ch <- GetGamesResponse{
			Error: err,
		}
		return
	}

	ch <- GetGamesResponse{
		Games: getGames.Games,
		Error: nil,
	}
}

func (n *Network) CreateGame(input CreateGameRequest, ch chan<- CreateGameResponse) {
	body, err := json.Marshal(input)
	if err != nil {
		ch <- CreateGameResponse{
			Error: err,
		}
		return
	}

	response, err := http.Post(
		fmt.Sprintf("%s/games", n.serverHost),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		ch <- CreateGameResponse{
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		ch <- CreateGameResponse{
			Error: err,
		}
		return
	}

	createGame := CreateGame{}
	if err := json.Unmarshal(body, &createGame); err != nil {
		ch <- CreateGameResponse{
			Error: err,
		}
		return
	}

	ch <- CreateGameResponse{
		HostUuid: createGame.HostUuid,
		Error:    nil,
	}
}

func (n *Network) JoinGame(input JoinGameRequest, ch chan<- JoinGameResponse) {
	body, err := json.Marshal(input)
	if err != nil {
		ch <- JoinGameResponse{
			Error: err,
		}
		return
	}

	response, err := http.Post(
		fmt.Sprintf("%s/games/join", n.serverHost),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		ch <- JoinGameResponse{
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		ch <- JoinGameResponse{
			Error: err,
		}
		return
	}

	joinGame := JoinGame{}
	if err := json.Unmarshal(body, &joinGame); err != nil {
		ch <- JoinGameResponse{
			Error: err,
		}
		return
	}

	ch <- JoinGameResponse{
		OpponentUuid: joinGame.OpponentUuid,
		Error:        nil,
	}
}

func (n *Network) StartGame(input StartGameRequest, ch chan<- StartGameResponse) {
	body, err := json.Marshal(input)
	if err != nil {
		ch <- StartGameResponse{
			Error: err,
		}
		return
	}

	response, err := http.Post(
		fmt.Sprintf("%s/games/start", n.serverHost),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		ch <- StartGameResponse{
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		ch <- StartGameResponse{
			Error: err,
		}
		return
	}

	startGame := StartGame{}
	if err := json.Unmarshal(body, &startGame); err != nil {
		ch <- StartGameResponse{
			Error: err,
		}
		return
	}

	ch <- StartGameResponse{
		Error: nil,
	}
}

func (n *Network) Wait(input WaitRequest, ch chan<- WaitResponse) {
	body, err := json.Marshal(input)
	if err != nil {
		ch <- WaitResponse{
			Error: err,
		}
		return
	}

	statusCode := 0
	response := &http.Response{}

	for err != nil || statusCode != 200 {
		response, err = http.Post(
			fmt.Sprintf("%s/games/wait", n.serverHost),
			"application/json",
			bytes.NewReader(body),
		)
		if response != nil {
			statusCode = response.StatusCode
		}
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		ch <- WaitResponse{
			Error: err,
		}
		return
	}

	wait := Wait{}
	if err := json.Unmarshal(body, &wait); err != nil {
		ch <- WaitResponse{
			Error: err,
		}
		return
	}

	ch <- WaitResponse{
		Status:        GameStatus(wait.Status),
		X:             wait.X,
		Y:             wait.Y,
		DestroyedShip: Ship(wait.DestroyedShip),
		DestroyedX:    wait.DestroyedX,
		DestroyedY:    wait.DestroyedY,
		Message:       wait.Message,
		Error:         nil,
	}
}

func (n *Network) Shoot(input ShootRequest, ch chan<- ShootResponse) {
	body, err := json.Marshal(input)
	if err != nil {
		ch <- ShootResponse{
			Error: err,
		}
		return
	}

	response, err := http.Post(
		fmt.Sprintf("%s/games/shoot", n.serverHost),
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		ch <- ShootResponse{
			Error: err,
		}
		return
	}
	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		ch <- ShootResponse{
			Error: err,
		}
		return
	}

	shoot := Shoot{}
	if err := json.Unmarshal(body, &shoot); err != nil {
		ch <- ShootResponse{
			Error: err,
		}
		return
	}

	ch <- ShootResponse{
		Status:        GameStatus(shoot.Status),
		DestroyedShip: Ship(shoot.DestroyedShip),
		X:             shoot.X,
		Y:             shoot.Y,
		Error:         nil,
	}
}
