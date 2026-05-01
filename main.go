// Show user's name and pfp

package main

import (
	"isthereanydeal/config"
	isthereanydeal "isthereanydeal/itad"
	mytypes "isthereanydeal/my-types"
	"isthereanydeal/ntfy"
	"isthereanydeal/steam"
	"log"
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
				log.Println("Unable to find a game")
			}
		})
	}

	go func() {
		gameWG.Wait()
		close(dealChan)
	}()

	ntfy.SendDeals(cfg, dealChan)
}
