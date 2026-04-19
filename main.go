package main

import (
	"fmt"
	"isthereanydeal/config"
	isthereanydeal "isthereanydeal/itad"
	mytypes "isthereanydeal/my-types"
	"isthereanydeal/ntfy"
	"isthereanydeal/steam"
	"sync"
)

func main() {
	cfg := config.LoadTOML()

	steamWishlist := steam.GetWishlist(cfg)

	dealChan := make(chan mytypes.Deals, len(steamWishlist))
	var gameWG sync.WaitGroup

	for _, id := range steamWishlist {
		gameWG.Go(func() {
			gameInfo, err := isthereanydeal.GetDealFromSteamId(cfg, id)
			if err == nil {
				isthereanydeal.GetDealFromGameInfo(cfg, gameInfo, dealChan)
			} else {
				fmt.Println("unable to find a game")
			}
		})
	}

	go func() {
		gameWG.Wait()
		close(dealChan)
	}()

	ntfy.SendDeals(cfg, dealChan)
}
