package ntfy

import (
	"fmt"
	mytypes "isthereanydeal/my-types"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func SendDeals(cfg mytypes.TOMLConfig, dealChan <-chan mytypes.Deals) error {
	allNotifications := map[string][]mytypes.Deals{}

	totalNumberOfDeals := 0

	for deal := range dealChan {
		allNotifications[deal.GameName] = append(allNotifications[deal.GameName], deal)
		totalNumberOfDeals += 1
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
	url := "https://ntfy.sh/" + cfg.Config.NTFYTopic

	var body strings.Builder
	fmt.Fprintf(&body, "Found a deal on %v (normally $%.2f)", gameInfo.GameName, gameInfo.Regular.Amount)

	for _, deal := range gameDeals {
		fmt.Fprintf(&body, "\n$%.2f (%v%% off) at [%v](%v)", deal.Price.Amount, deal.Cut, deal.Shop.Name, deal.Url)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(body.String()))

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
