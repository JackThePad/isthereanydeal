package mytypes

// Used for holding the resonse from isthereanydeal useful as it holds all the important details except for the game's price
type GameInfo struct {
	Found bool `json:"found"`
	Game  struct {
		Id     string `json:"id"`
		Slug   string `json:"slug"`
		Title  string `json:"title"`
		Assets struct {
			BannerArt string `json:"banner600"`
		} `json:"assets"`
	} `json:"game"`
}

type GamePrices struct {
	Id        string `json:"id"`
	Name      string
	Deals     []Deals `json:"deals"`
	BannerArt string
}

// In isthereanydeal will return a list of these if there is an active deal available
type Deals struct {
	GameName  string
	GameId    string // This is isthereanydeal's game id and not steam's
	Shop      shop   `json:"shop"`
	Price     price  `json:"price"`   // The current price on sale
	Regular   price  `json:"regular"` // The price the game is when not on sale
	Cut       int    `json:"cut"`     // The percent off
	Expiry    string `json:"expiry"`  // The time in unix when the deal is over, some stores do not provide this so it is untrustworthy
	Url       string `json:"url"`     // Where the deal is available and will only change when a new deal is generated
	BannerArt string
}

type shop struct {
	Id   int    `json:"id"` // Each store has it's own id
	Name string `json:"name"`
}

type price struct {
	Amount    float64 `json:"amount"`    // 19.99
	AmountInt int     `json:"amountInt"` // 1999
}

type TOMLConfig struct {
	Config struct {
		ITADAPIKey  string `toml:"itad_api_key"`
		SteamApiKey string `toml:"steam_api_key"`
		SteamId     string `toml:"steam_account_id"`
		NTFYUrl     string `toml:"ntfy_url"`
		JsonName    string `toml:"json_name"`
	} `toml:"config"`
}

type SteamWishlist struct {
	Response struct {
		Items []struct {
			AppId int `json:"appid"`
		} `json:"items"`
	} `json:"response"`
}

type JsonData struct {
	Expiry string `json:"expiry"`
	Url    string `json:"url"`
}
