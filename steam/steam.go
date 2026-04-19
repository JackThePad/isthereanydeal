package steam

import (
	"encoding/json"
	"io"
	mytypes "isthereanydeal/my-types"
	"log"
	"net/http"
	"net/url"
)

func GetWishlist(cfg mytypes.TOMLConfig) []int {
	var org_url string = "https://api.steampowered.com/IWishlistService/GetWishlist/v1/"

	p_url, err := url.Parse(org_url)

	if err != nil {
		log.Fatal(err)
	}

	query := p_url.Query()

	query.Add("key", cfg.Config.SteamApiKey)
	query.Add("steamid", cfg.Config.SteamId)

	p_url.RawQuery = query.Encode()

	resp, err := http.Get(p_url.String())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var steamWishlist mytypes.SteamWishlist

	err = json.Unmarshal(data, &steamWishlist)

	if err != nil {
		log.Fatal(err)
	}

	var steamGameIds []int

	for _, steamGame := range steamWishlist.Response.Items {
		steamGameIds = append(steamGameIds, steamGame.AppId)
	}

	return steamGameIds
}
