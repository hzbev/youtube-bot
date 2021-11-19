package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	"yt-bot/helper"

	"github.com/corpix/uarand"
	"github.com/go-resty/resty/v2"
)

var wg sync.WaitGroup

var streamID string = "4YBfEVq33R4"

func main() {
	proxyList := helper.ReadtoArray("proxy.txt")

	for _, proxy := range proxyList {
		wg.Add(1)
		go startBot(proxy)
	}
	wg.Wait()
}

func startBot(proxy_url string) {
	for {
		client := resty.New()
		client.SetHeader("user-agent", uarand.GetRandom())
		client.SetProxy("http://" + proxy_url)
		client.SetCookie(&http.Cookie{
			Name:   "CONSENT",
			Value:  "YES+srp.gws-20211104-0-RC1.en+FX+751",
			Domain: "https://www.youtube.com",
		})

		resp, _ := client.R().
			Get("https://www.youtube.com/watch?v=" + streamID)
		res := resp.String()
		fmt.Println(resp.StatusCode())
		startInd := strings.Index(res, "videostatsWatchtimeUrl")
		endInd := strings.Index(res, "ptrackingUrl")
		if endInd == -1 || startInd == -1 {
			fmt.Println("uwuwu")
			continue
		}
		ytURL := res[startInd+36 : endInd-4]
		ytURL = strings.ReplaceAll(ytURL, `\u0026`, "&")
		fmt.Println(ytURL)
		query, _ := url.ParseQuery(ytURL)
		cl := query.Get("cl")
		ei := query.Get("ei")
		of := query.Get("of")
		vm := query.Get("vm")
		cpn := helper.RandString(16)
		st := helper.RandInt(1000, 10000)
		et := "0"
		lio := fmt.Sprintf("%f", float64(time.Now().UnixMilli())/1000.0)[:14]

		rt := helper.RandInt(10, 200)
		lact := helper.RandInt(1000, 8000)
		rtn := rt + 300

		apiURL, _ := url.Parse("https://s.youtube.com/api/stats/watchtime?")
		_apiURL := apiURL.Query()
		_apiURL.Add("ns", "yt")
		_apiURL.Add("el", "detailpage")

		_apiURL.Add("cpn", cpn)
		_apiURL.Add("docid", streamID)

		_apiURL.Add("ver", "2")
		_apiURL.Add("cmt", et)
		_apiURL.Add("ei", ei)
		_apiURL.Add("fmt", "243")
		_apiURL.Add("fs", "0")
		_apiURL.Add("rt", strconv.Itoa(rt))
		_apiURL.Add("of", of)
		_apiURL.Add("euri", "")
		_apiURL.Add("lact", strconv.Itoa(lact))
		_apiURL.Add("cl", cl)
		_apiURL.Add("state", "playing")
		_apiURL.Add("vm", vm)
		_apiURL.Add("volume", "100")
		_apiURL.Add("cbr", "Firefox")
		_apiURL.Add("cbrver", "83.0")
		_apiURL.Add("c", "WEB")
		_apiURL.Add("cplayer", "UNIPLAYER")
		_apiURL.Add("cver", "2.20201210.01.00")
		_apiURL.Add("cos", "Windows")
		_apiURL.Add("cosver", "10")
		_apiURL.Add("cplatform", "DESKTOP")
		_apiURL.Add("delay", "5")
		_apiURL.Add("hl", "en_US")
		_apiURL.Add("rtn", strconv.Itoa(rtn))
		_apiURL.Add("aftm", "140")
		_apiURL.Add("rti", strconv.Itoa(rt))
		_apiURL.Add("muted", "0")
		_apiURL.Add("st", strconv.Itoa(st))
		_apiURL.Add("et", et)
		_apiURL.Add("lio", lio)

		client.SetHeaders(map[string]string{
			"Accept-Encoding": "gzip, deflate",
			"Host":            "www.youtube.com",
		})

		aoao, _ := client.R().Get("https://s.youtube.com/api/stats/playback?" + _apiURL.Encode())
		fmt.Println(aoao.StatusCode())
		time.Sleep(3 * time.Second)
	}
	defer wg.Done()

}
