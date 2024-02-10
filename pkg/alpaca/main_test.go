package alpaca

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/joho/godotenv"
)

// enum symbols
type Symbol string

func (s Symbol) ToString() string {
	return string(s)
}

const APPLE Symbol = "AAPL"
const MICROSOFT Symbol = "MSFT"
const TESLA Symbol = "TSLA"
const AMAZON Symbol = "AMZN"
const GOOGLE Symbol = "GOOGL"

type Portafolio struct {
	// DesiredState is the desired state of the portfolio
	DesiredState map[Symbol]int `json:"desiredState"`
	// CurrentState is the current state of the portfolio
	CurrentState map[Symbol]int `json:"currentState"`
}

func (p *Portafolio) IsCorrect(client *alpaca.Client, marketdata *marketdata.Client) (bool, error) {
	// validate all symbols exist by getting the current price
	for symbol, _ := range p.DesiredState {
		_, err := client.GetAsset(symbol.ToString())
		if err != nil {
			return false, err
		}
	}

	for symbol, desiredAmount := range p.DesiredState {
		currentAmount, ok := p.CurrentState[symbol]
		if !ok {
			return false, nil
		}
		if currentAmount != desiredAmount {
			return false, nil
		}
	}

	return true, nil
}

func NewPortafolio(desiredPortafolio map[Symbol]int) *Portafolio {
	return &Portafolio{
		DesiredState: desiredPortafolio,
		CurrentState: make(map[Symbol]int),
	}
}

var desiredPortafolio = map[Symbol]int{
	APPLE:     1,
	MICROSOFT: 2,
	TESLA:     3,
	AMAZON:    4,
	GOOGLE:    5,
}

var proxyClient http.Client = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   "localhost:9091",
		}),
	},
}

func TestAlpaca(t *testing.T) {
	client := alpaca.NewClient(alpaca.ClientOpts{
		HTTPClient: &proxyClient,
		// Alternatively you can set your key and secret using the
		// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
		APIKey:    os.Getenv("ALPACA_API_KEY"),
		APISecret: os.Getenv("ALPACA_API_SECRET_KEY"),
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	account, err := client.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%+v\n", *account)

	balanceChange := account.Equity.Sub(account.LastEquity)

	fmt.Println("Today's portfolio balance change:", balanceChange)

	// list poratfolio
	portfolio, err := client.GetPositions()
	if err != nil {
		panic(err)
	}
	if len(portfolio) == 0 {
		fmt.Println("No positions in portfolio")
	}

	for _, position := range portfolio {
		fmt.Printf("%+v\n", position)
	}

}

func TestMarketData(t *testing.T) {
	client := marketdata.NewClient(marketdata.ClientOpts{
		HTTPClient: &proxyClient,
		APIKey:     os.Getenv("ALPACA_API_KEY"),
		// Alternatively you can set your key using the APCA_API_KEY_ID environment variable
		APISecret: os.Getenv("ALPACA_API_SECRET_KEY"),
		BaseURL:   "https://data.alpaca.markets",
	})
	// Get the latest trade for a given stock
	trade, err := client.GetLatestTrade("AAPL", marketdata.GetLatestTradeRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Latest trade for AAPL: %+v\n", prettyPrint(trade))
}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(s)
}

func TestPortafolio(t *testing.T) {
	portafolio := NewPortafolio(desiredPortafolio)
	client := alpaca.NewClient(alpaca.ClientOpts{
		HTTPClient: &proxyClient,
		// Alternatively you can set your key and secret using the
		// APCA_API_KEY_ID and APCA_API_SECRET_KEY environment variables
		APIKey:    os.Getenv("ALPACA_API_KEY"),
		APISecret: os.Getenv("ALPACA_API_SECRET_KEY"),
		BaseURL:   "https://paper-api.alpaca.markets",
	})

	marketdataClient := marketdata.NewClient(marketdata.ClientOpts{
		HTTPClient: &proxyClient,
		APIKey:     os.Getenv("ALPACA_API_KEY"),
		// Alternatively you can set your key using the APCA_API_KEY_ID environment variable
		APISecret: os.Getenv("ALPACA_API_SECRET_KEY"),
		BaseURL:   "https://data.alpaca.markets",
	})

	portafolio = NewPortafolio(desiredPortafolio)

	isCorrect, err := portafolio.IsCorrect(client, marketdataClient)
	if err != nil {
		panic(err)
	}
	fmt.Println("Is correct:", isCorrect)

}

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
