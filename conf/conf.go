package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Url      string `json:"url"`
	ButtonID string `json:"button_id"`
}

func Parse() *Conf {
	var conf Conf
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return &conf
}
