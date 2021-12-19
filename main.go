package main

import (
	"strconv"
	"strings"

	_ "github.com/Cadrew/go-ggpo/ggpo/ggpolog"
	"github.com/Cadrew/go-ggpo/ggpo/ggponet"
	"github.com/Cadrew/go-ggpo/netplay"
	"github.com/sirupsen/logrus"
)

func runLoop() {
	// Loop of the game
}

func main() {
	// data example
	numPlayers := 2
	playersIP := ["127.0.0.1", "0.0.0.0"]
	localPort := "5000"
	InitNetwork(numPlayers, playersIP, localPort)
	// run the game
	runLoop()
}

func InitNetwork(numPlayers int, playersIP []string, localPort string) {
	players := make([]ggponet.GGPOPlayer, ggponet.GGPO_MAX_SPECTATORS+ggponet.GGPO_MAX_PLAYERS)

	for i := 0; i < numPlayers; i++ {
		players[i].PlayerNum = int64(i + 1)
		if playersIP[i] == "local" {
			players[i].Type = ggponet.GGPO_PLAYERTYPE_LOCAL
			continue
		}
		players[i].Type = ggponet.GGPO_PLAYERTYPE_REMOTE
		players[i].IPAddress = strings.Split(playersIP[i], ":")[0]
		port, err := strconv.Atoi(strings.Split(playersIP[i], ":")[1])
		players[i].Port = uint64(port)
		if err != nil {
			logrus.Panic("Error in InitNetwork")
		}
	}
	//TODO: Spectators

	netplay.Init(int64(numPlayers), players, localPort, 0, false)
}
