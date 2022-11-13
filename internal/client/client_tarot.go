package client

import (
	"QQ_bot/types"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
)

const tarotURL = "https://rws-cards-api.herokuapp.com/api/v1/cards/random?n=1"

func GetTarotCard() *types.TarotCard {
	var (
		client = g.Client()
		tarot  = types.TarotCard{}
	)

	res, err := client.Get(context.Background(), tarotURL)
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &tarot)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &tarot
}
