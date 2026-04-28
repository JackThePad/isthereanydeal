package ntfy

import (
	"encoding/json"
	"fmt"
	mytypes "isthereanydeal/my-types"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func GetJsonContent(cfg mytypes.TOMLConfig) map[string]map[int]mytypes.JsonData {
	jsonContent, err := os.ReadFile(cfg.Config.JsonName)

	if err != nil {
		log.Fatal("Unable to load json file please make sure you have the permission to read and write")
	}

	var jsonDeals map[string]map[int]mytypes.JsonData

	_ = json.Unmarshal(jsonContent, &jsonDeals)

	return jsonDeals
}

func addDealsToJson(cfg mytypes.TOMLConfig, deal mytypes.Deals) {
	var seenDeals map[string]map[int]mytypes.JsonData

	data, err := os.ReadFile(cfg.Config.JsonName)

	if err != nil {
		log.Fatal("Unable to load json file please make sure you have the permission to read and write")
	} else if len(data) == 0 {
		seenDeals = map[string]map[int]mytypes.JsonData{}
	}

	json.Unmarshal(data, &seenDeals)

	if seenDeals[deal.GameId] == nil {
		seenDeals[deal.GameId] = map[int]mytypes.JsonData{}
	}

	seenDeals[deal.GameId][deal.Shop.Id] = mytypes.JsonData{Url: deal.Url, Expiry: deal.Expiry}

	data, err = json.Marshal(seenDeals)

	if err != nil {
		log.Fatal(err)
	}

	_ = os.WriteFile(cfg.Config.JsonName, data, 0644)

}

func SendDeals(cfg mytypes.TOMLConfig, dealChan <-chan mytypes.Deals) error {
	allNotifications := map[string][]mytypes.Deals{}

	totalNumberOfDeals := 0

	jsonData := GetJsonContent(cfg)

	for deal := range dealChan {
		if checkExperation(jsonData, deal) {
			allNotifications[deal.GameId] = append(allNotifications[deal.GameId], deal)
			addDealsToJson(cfg, deal)
			totalNumberOfDeals += 1
		}
	}

	fmt.Printf("Found %v deals\n", totalNumberOfDeals)

	// Sorts the game prices
	for key := range allNotifications {
		sort.Slice(allNotifications[key], func(i, j int) bool {
			return allNotifications[key][i].Price.AmountInt < allNotifications[key][j].Price.AmountInt
		})
	}

	for key, deals := range allNotifications {
		sendDealByGame(cfg, deals, allNotifications[key][0])
		time.Sleep(250 * time.Millisecond)
	}

	return nil
}

func sendDealByGame(cfg mytypes.TOMLConfig, gameDeals []mytypes.Deals, gameInfo mytypes.Deals) error {
	var body strings.Builder
	fmt.Fprintf(&body, "Found a deal on %v (normally $%.2f)", gameInfo.GameName, gameInfo.Regular.Amount)

	for _, deal := range gameDeals {
		fmt.Fprintf(&body, "\n$%.2f (%v%% off) at [%v](%v)", deal.Price.Amount, deal.Cut, deal.Shop.Name, deal.Url)
	}

	req, err := http.NewRequest("POST", cfg.Config.NTFYUrl, strings.NewReader(body.String()))

	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	title := fmt.Sprintf("Found a deal on %s", gameInfo.GameName)

	req.Header.Set("Title", title)
	req.Header.Set("Attach", gameInfo.BannerArt)
	req.Header.Set("Markdown", "yes")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return nil
}

func checkExperation(jsonData map[string]map[int]mytypes.JsonData, deal mytypes.Deals) bool {
	if jsonData[deal.GameId][deal.Shop.Id].Url == deal.Url {
		return false
	} else {
		return true
	}

}
