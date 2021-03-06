package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var Games = make(map[string][]string, 0)

var gamesurl = "https://gsm.skuzzi.ro/api/games.php"
var ExecutablesURL = "https://raw.githubusercontent.com/Game-Server-Management/gsm-main/master/build/"

func LoadAPI() {
	response, err := http.Get(gamesurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(html, &Games)
}
