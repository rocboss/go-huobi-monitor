package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/huobi"
	"github.com/rocboss/go-huobi-monitor/utils"
)

func main() {
	// 加载配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化Redis
	if initClient() != nil {
		log.Fatal("Error connect to redis server")
	}

	spotWs := huobi.NewSpotWs()

	spotWs.TickerCallback(func(ticker *goex.Ticker) {
		// 获取所有监听规则
		monitors, err := rdb.SMembers("HBPairs:" + ticker.Pair.String()).Result()
		if err == nil {
			if len(monitors) > 0 {
				alerts := []string{}
				for _, monitor := range monitors {
					monitorRune := []rune(monitor)

					if string(monitorRune[:1]) == "<" {
						target, err := strconv.ParseFloat(string(monitorRune[1:]), 64)
						if err != nil {
							continue
						}
						// 小于类型匹配
						if ticker.Last <= target {
							alerts = append(alerts, "[Hit] "+ticker.Pair.String()+" : "+monitor, fmt.Sprintf("Latest Price: %.4f", ticker.Last), "")
						}
					}

					if string(monitorRune[:1]) == ">" {
						target, err := strconv.ParseFloat(string(monitorRune[1:]), 64)
						if err != nil {
							continue
						}
						// 小于类型匹配
						if ticker.Last >= target {
							alerts = append(alerts, "[Hit] "+ticker.Pair.String()+" : "+monitor, fmt.Sprintf("Latest Price: %.4f", ticker.Last), "")
						}
					}
				}
				if len(alerts) > 0 {
					log.Println(alerts)

					_, err = rdb.Get("MSG_LOCK:" + ticker.Pair.String()).Result()
					if err != nil {
						utils.PushMessage(append(alerts, "(该交易对提醒静默1分钟)"))
						// 防止重复发送（1分钟再次发送）
						rdb.Set("MSG_LOCK:"+ticker.Pair.String(), 1, time.Minute)
					}
				}
			}
		}
	})

	// 初始化监听
	pairs, err := rdb.Keys("HBPairs:*").Result()
	if err == nil {
		for _, pair := range pairs {
			spotWs.SubscribeTicker(goex.NewCurrencyPair2(strings.Split(pair, ":")[1]))
		}
	}

	e := echo.New()

	e.POST("/dingtalk_webhook", func(c echo.Context) error {
		req := new(ReqData)
		if err = c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// TODO 校验签名

		content := req.Text.Content

		// eg. add zksusdt 1.2
		commands := strings.Split(strings.Trim(content, " "), " ")
		if len(commands) == 3 {
			switch commands[0] {
			case "add":
				// 添加监听
				rdb.SAdd("HBPairs:"+commands[1], commands[2])
				spotWs.SubscribeTicker(goex.NewCurrencyPair2(commands[1]))
				utils.PushMessage([]string{"命令执行成功"})

			case "del":
				rdb.SRem("HBPairs:"+commands[1], commands[2])
				utils.PushMessage([]string{"命令执行成功"})

			case "list":
				monitors := []string{}

				pairs, err := rdb.Keys("HBPairs:*").Result()
				if err == nil {
					for _, pair := range pairs {
						ms, err := rdb.SMembers(pair).Result()
						if err == nil {
							for _, monitor := range ms {
								monitors = append(monitors, strings.Split(pair, ":")[1]+": "+monitor)
							}
						}
					}
				}
				utils.PushMessage(monitors)

			default:
				utils.PushMessage([]string{
					"可用命令",
					"",
					"添加监控",
					"add [COIN] [PRICE]",
					"eg. add BTC_USDT <43000",
					"",
					"删除监控",
					"del [COIN] [PRICE]",
					"eg. del BTC_USDT <43000",
					"",
					"查看监控",
					"list all coins",
				})
			}
		} else {
			utils.PushMessage([]string{
				"可用命令",
				"",
				"添加监控",
				"add [COIN] [PRICE]",
				"eg. add BTC_USDT <43000",
				"",
				"删除监控",
				"del [COIN] [PRICE]",
				"eg. del BTC_USDT <43000",
				"",
				"查看监控",
				"list all coins",
			})
		}

		// utils.PushMessage()
		return c.String(http.StatusOK, "success")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
