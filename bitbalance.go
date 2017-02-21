package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"reflect"
	"github.com/BurntSushi/toml"
)

const addressAPI = "http://btc.blockr.io/api/v1/address/info/"
const exchangeAPI = "https://api.coinbase.com/v2/exchange-rates?currency=BTC"

type AddrResponse struct {
	Status  string         `json:"status"`
	Data    []ResponseData `json:"data"`
	Code    float64        `json:"code"`
	Message string         `json:"message"`
}

type ExchangeResponse struct {
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Address         string           `json:"address"`
	IsKnown         bool             `json:"is_unknown"`
	Balance         float64          `json:"balance"`
	BalanceMultiSig float64          `json:"balance_multisig"`
	TotalRecieved   float64          `json:"totalreceived"`
	NbTxs           float64          `json:"nb_txs"`
	FirstTxs        ResponseFirstTxs `json:"first_tx"`
	LastTxs         ResponseLastTxs  `json:"last_tx"`
	IsValid         bool             `json:"is_valid"`

	Currency        string           `json:currency`
	Rates          	Rates            `json:rates`
}

type ResponseFirstTxs struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:"block_nb"`
	Value         float64 `json:"value"`
	Confirmations int64   `json:"confirmations"`
}

type ResponseLastTxs struct {
	Time          string  `json:"time_utc"`
	Tx            string  `json:"tx"`
	BlockNb       string  `json:block_nb"`
	Value         float64 `json:value"`
	Confirmations int64   `json:confirmations"`
}

type Rates struct {
	AED float64 `json:",string"`
	AFN float64 `json:",string"`
	ALL float64 `json:",string"`
	AMD float64 `json:",string"`
	ANG float64 `json:",string"`
	AOA float64 `json:",string"`
	ARS float64 `json:",string"`
	AUD float64 `json:",string"`
	AWG float64 `json:",string"`
	AZN float64 `json:",string"`
	BAM float64 `json:",string"`
	BBD float64 `json:",string"`
	BDT float64 `json:",string"`
	BGN float64 `json:",string"`
	BHD float64 `json:",string"`
	BIF float64 `json:",string"`
	BMD float64 `json:",string"`
	BND float64 `json:",string"`
	BOB float64 `json:",string"`
	BRL float64 `json:",string"`
	BSD float64 `json:",string"`
	BTC float64 `json:",string"`
	BTN float64 `json:",string"`
	BWP float64 `json:",string"`
	BYN float64 `json:",string"`
	BYR float64 `json:",string"`
	BZD float64 `json:",string"`
	CAD float64 `json:",string"`
	CDF float64 `json:",string"`
	CHF float64 `json:",string"`
	CLF float64 `json:",string"`
	CLP float64 `json:",string"`
	CNY float64 `json:",string"`
	COP float64 `json:",string"`
	CRC float64 `json:",string"`
	CUC float64 `json:",string"`
	CVE float64 `json:",string"`
	CZK float64 `json:",string"`
	DJF float64 `json:",string"`
	DKK float64 `json:",string"`
	DOP float64 `json:",string"`
	DZD float64 `json:",string"`
	EEK float64 `json:",string"`
	EGP float64 `json:",string"`
	ERN float64 `json:",string"`
	ETB float64 `json:",string"`
	ETH float64 `json:",string"`
	EUR float64 `json:",string"`
	FJD float64 `json:",string"`
	FKP float64 `json:",string"`
	GBP float64 `json:",string"`
	GEL float64 `json:",string"`
	GGP float64 `json:",string"`
	GHS float64 `json:",string"`
	GIP float64 `json:",string"`
	GMD float64 `json:",string"`
	GNF float64 `json:",string"`
	GTQ float64 `json:",string"`
	GYD float64 `json:",string"`
	HKD float64 `json:",string"`
	HNL float64 `json:",string"`
	HRK float64 `json:",string"`
	HTG float64 `json:",string"`
	HUF float64 `json:",string"`
	IDR float64 `json:",string"`
	ILS float64 `json:",string"`
	IMP float64 `json:",string"`
	INR float64 `json:",string"`
	IQD float64 `json:",string"`
	ISK float64 `json:",string"`
	JEP float64 `json:",string"`
	JMD float64 `json:",string"`
	JOD float64 `json:",string"`
	JPY float64 `json:",string"`
	KES float64 `json:",string"`
	KGS float64 `json:",string"`
	KHR float64 `json:",string"`
	KMF float64 `json:",string"`
	KRW float64 `json:",string"`
	KWD float64 `json:",string"`
	KYD float64 `json:",string"`
	KZT float64 `json:",string"`
	LAK float64 `json:",string"`
	LBP float64 `json:",string"`
	LKR float64 `json:",string"`
	LRD float64 `json:",string"`
	LSL float64 `json:",string"`
	LTL float64 `json:",string"`
	LVL float64 `json:",string"`
	LYD float64 `json:",string"`
	MAD float64 `json:",string"`
	MDL float64 `json:",string"`
	MGA float64 `json:",string"`
	MKD float64 `json:",string"`
	MMK float64 `json:",string"`
	MNT float64 `json:",string"`
	MOP float64 `json:",string"`
	MRO float64 `json:",string"`
	MTL float64 `json:",string"`
	MUR float64 `json:",string"`
	MVR float64 `json:",string"`
	MWK float64 `json:",string"`
	MXN float64 `json:",string"`
	MYR float64 `json:",string"`
	MZN float64 `json:",string"`
	NAD float64 `json:",string"`
	NGN float64 `json:",string"`
	NIO float64 `json:",string"`
	NOK float64 `json:",string"`
	NPR float64 `json:",string"`
	NZD float64 `json:",string"`
	OMR float64 `json:",string"`
	PAB float64 `json:",string"`
	PEN float64 `json:",string"`
	PGK float64 `json:",string"`
	PHP float64 `json:",string"`
	PKR float64 `json:",string"`
	PLN float64 `json:",string"`
	PYG float64 `json:",string"`
	QAR float64 `json:",string"`
	RON float64 `json:",string"`
	RSD float64 `json:",string"`
	RUB float64 `json:",string"`
	RWF float64 `json:",string"`
	SAR float64 `json:",string"`
	SBD float64 `json:",string"`
	SCR float64 `json:",string"`
	SEK float64 `json:",string"`
	SGD float64 `json:",string"`
	SHP float64 `json:",string"`
	SLL float64 `json:",string"`
	SOS float64 `json:",string"`
	SRD float64 `json:",string"`
	STD float64 `json:",string"`
	SVC float64 `json:",string"`
	SZL float64 `json:",string"`
	THB float64 `json:",string"`
	TJS float64 `json:",string"`
	TMT float64 `json:",string"`
	TND float64 `json:",string"`
	TOP float64 `json:",string"`
	TRY float64 `json:",string"`
	TTD float64 `json:",string"`
	TWD float64 `json:",string"`
	TZS float64 `json:",string"`
	UAH float64 `json:",string"`
	UGX float64 `json:",string"`
	USD float64 `json:",string"`
	UYU float64 `json:",string"`
	UZS float64 `json:",string"`
	VEF float64 `json:",string"`
	VND float64 `json:",string"`
	VUV float64 `json:",string"`
	WST float64 `json:",string"`
	XAF float64 `json:",string"`
	XAG float64 `json:",string"`
	XAU float64 `json:",string"`
	XCD float64 `json:",string"`
	XDR float64 `json:",string"`
	XOF float64 `json:",string"`
	XPF float64 `json:",string"`
	YER float64 `json:",string"`
	ZAR float64 `json:",string"`
	ZMK float64 `json:",string"`
	ZMW float64 `json:",string"`
	ZWL float64 `json:",string"`
}

func (r *Rates) get(field string) float64 {
	return float64(reflect.Indirect(reflect.ValueOf(r)).FieldByName(field).Float())
}

func getResponse(url string, userAgent string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Couldn't create request: ", err)
	}

	request.Header.Set("User-Agent", userAgent)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("HTTP request failed: ", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading body: ", err)
	}
	return body
}

func getAddrResponse(url string, userAgent string) *AddrResponse {
	body := getResponse(url, userAgent)
	if body == nil {
		return nil
	}
	resp := &AddrResponse{}
	if err := json.Unmarshal(body, resp); err != nil {
		log.Fatal("Unmarshal failed: ", err)
	}
	return resp
}

func getExchangeResponse(url string, userAgent string) *ExchangeResponse {
	body := getResponse(url, userAgent)
	if body == nil {
		return nil
	}
	resp := &ExchangeResponse{}
	if err := json.Unmarshal(body, resp); err != nil {
		log.Fatal("Unmarshal failed: ", err)
	}
	return resp
}

func readFile(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading the file: ", err)
	}
	return strings.Split(string(data), "\n")
}

func getOS() string {
	return runtime.GOOS
}

func userAgent() string {
	userAgent := `Mozilla/5.0 `
	switch getOS() {
		case "linux":
			userAgent += `(X11; Linux x86_64)`
		case "windows":
			userAgent += `(Windows NT 10.0; WOW64)`
		case "mac":
			userAgent += `(Macintosh; Intel Mac OS X 10_11_6)`
	}
	userAgent += ` AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36`
	return userAgent
}

// Info from config file
type Config struct {
	Currency  string
}

// Reads info from config file
func loadConfig() Config {
	var confFile = "conf.toml"
	if _, err := os.Stat(confFile); err != nil {
		log.Fatal("Config file is missing: ", err)
	}

	var conf Config
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatal("Failed to parse config: ", err)
	}
	return conf
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(`Usage: bitbalance [address|file]...`)
	}
	addresses := []string{""}
	for i, addressOrFile := range os.Args {
		if i == 0 {
			continue
		}
		if _, err := os.Stat(addressOrFile); !os.IsNotExist(err) {
			addresses = append(addresses, readFile(addressOrFile)...)
		} else {
			addresses = append(addresses, addressOrFile)
		}
	}

	conf := loadConfig()

	exchangeResp := getExchangeResponse(exchangeAPI, userAgent())
	btcToLocal := exchangeResp.Data.Rates.BTC / exchangeResp.Data.Rates.get(conf.Currency)

	sum := 0.0
	count := 0
	uri := strings.Join(addresses, ",")
	addrResp := getAddrResponse(addressAPI + uri, userAgent())
	if addrResp != nil {
		for _, data := range addrResp.Data {
			if data.Address != "" {
				fmt.Printf("\033[92m%34s\033[0m ", data.Address)
				fmt.Printf("\033[95m%15.8f BTC\033[0m ", data.Balance)
				fmt.Printf("\033[95m%12.2f %s\033[0m\n", data.Balance / btcToLocal, conf.Currency)
				sum += data.Balance
				count++
			}
		}
	}
	if count > 1 {
		fmt.Printf("%34s ", "")
		fmt.Printf("\033[95m%15.8f BTC\033[0m ", sum)
		fmt.Printf("\033[95m%12.2f %s\033[0m\n", sum / btcToLocal, conf.Currency)
	}
}
