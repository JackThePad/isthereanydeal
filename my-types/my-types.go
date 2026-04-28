package mytypes

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

type Deals struct {
	GameName  string
	GameId    string
	Shop      shop   `json:"shop"`
	Price     price  `json:"price"`
	Regular   price  `json:"regular"`
	Cut       int    `json:"cut"`
	Expiry    string `json:"expiry"`
	Url       string `json:"url"`
	BannerArt string
}

type shop struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type price struct {
	Amount    float64 `json:"amount"`
	AmountInt int     `json:"amountInt"`
}

type TOMLConfig struct {
	Config struct {
		ITADClientId     string `toml:"itad_client_id"`
		ITADClientSecret string `toml:"itad_client_secret"`
		ITADAPIKey       string `toml:"itad_api_key"`
		SteamApiKey      string `toml:"steam_api_key"`
		SteamId          string `toml:"steam_account_id"`
		NTFYUrl          string `toml:"ntfy_url"`
		JsonName         string `toml:"json_name"`
	} `toml:"config"`
}

type SteamWishlist struct {
	Response struct {
		Items []struct {
			AppId int `json:"appid"`
		} `json:"items"`
	} `json:"response"`
}

// GameId string `json:"id"`
// append jsonData[gameid][shopId] = jsonData
type JsonData struct {
	Expiry string `json:"expiry"`
	Url    string `json:"url"`
}
