package isthereanydeal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	mytypes "isthereanydeal/my-types"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetDealFromSteamId(cfg mytypes.TOMLConfig, steamId int) (mytypes.GameInfo, error) {
	orgUrl := "https://api.isthereanydeal.com/games/lookup/v1"

	p_url, err := url.Parse(orgUrl)
	if err != nil {
		log.Fatal(err)
	}

	query := p_url.Query()

	query.Add("appid", strconv.Itoa(steamId))
	query.Add("key", cfg.Config.ITADAPIKey)

	p_url.RawQuery = query.Encode()

	resp, err := http.Get(p_url.String())

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var gameInfo mytypes.GameInfo

	err = json.Unmarshal(data, &gameInfo)

	if err != nil {
		log.Fatal(err)
	}

	if !gameInfo.Found {
		return mytypes.GameInfo{}, fmt.Errorf("Unable to find game")
	}

	return gameInfo, nil
}

func GetDealFromGameInfo(cfg mytypes.TOMLConfig, gameInfo mytypes.GameInfo, dealChan chan<- mytypes.Deals) []mytypes.Deals {
	gamePrices, err := getGameDeals(cfg, gameInfo)

	if err != nil {
		return []mytypes.Deals{}
	}

	gamePrices.BannerArt = gameInfo.Game.Assets.BannerArt

	validDeals := getValidDeal(gamePrices, dealChan)

	return validDeals
}

func getGameDeals(cfg mytypes.TOMLConfig, gameInfo mytypes.GameInfo) (mytypes.GamePrices, error) {
	p_url, err := url.Parse("https://api.isthereanydeal.com/games/prices/v3")

	if err != nil {
		log.Fatal(err)
	}

	query := p_url.Query()

	query.Add("key", cfg.Config.ITADAPIKey)
	query.Add("deals", "true")

	body := []string{
		gameInfo.Game.Id,
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		log.Fatal(err)
	}

	p_url.RawQuery = query.Encode()

	resp, err := http.Post(p_url.String(), "application/json", bytes.NewBuffer(jsonBody))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if string(data) == "[]" {
		return mytypes.GamePrices{}, fmt.Errorf("%v has no sales", gameInfo.Game.Title)
	}

	var deals []mytypes.GamePrices

	err = json.Unmarshal(data, &deals)

	if err != nil {
		log.Fatal(err)
	}

	deal := deals[0]

	deal.Name = gameInfo.Game.Title

	return deal, nil
}

func getValidDeal(gamePrice mytypes.GamePrices, dealChan chan<- mytypes.Deals) []mytypes.Deals {
	var validDeals []mytypes.Deals

	for _, deal := range gamePrice.Deals {
		deal.GameName = gamePrice.Name
		deal.BannerArt = gamePrice.BannerArt
		deal.GameId = gamePrice.Id

		if deal.Price.Amount <= 5 {
			validDeals = append(validDeals, deal)
			dealChan <- deal
		} else if deal.Cut >= 75 {
			validDeals = append(validDeals, deal)
			dealChan <- deal
		}

	}

	return validDeals
}
