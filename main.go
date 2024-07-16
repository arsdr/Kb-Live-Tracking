package main

import (
	"KBScraper/util"
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/websocket/v2"
)

type Transaction struct {
	Date string `json:"date"`
	IDTrx string `json:"id_trx"`
	Contact string `json:"contact_no"`
	Price string `json:"price"`
	Status string `json:"status"`
}

var (
	Status map[int]string = map[int]string{
		1: "Success",
		0: "Failed",
	}
)

func main() {
	app := fiber.New(
		fiber.Config{
			ServerHeader: "TuruLabs",
			AppName: "KiosBedul v2 System",
			EnablePrintRoutes: true,
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Get("/ws/transactions", websocket.New(func(c *websocket.Conn) {
		for {
			var trxstatus int
			probsNum, err := util.RandomWithProbability([]int{1, 0}, []int{85, 15})
			if err != nil {
				trxstatus = 0
			}
			trxstatus = probsNum

			lastTrx := Transaction{
				Date: time.Now().Format("2006-01-02 15:04:05"),
				IDTrx: fmt.Sprintf("KB*****%d", util.RandomWithDigitRange(3, 4)),
				Contact: fmt.Sprintf("08*****%d", util.RandomWithDigitRange(3, 3)),
				Price: util.FormatCurrency(util.RandomInRange(500, 900000)),
				Status: Status[trxstatus],
			}

			c.WriteJSON(lastTrx)
			time.Sleep(time.Duration(util.RandomInRange(1, 1000)) * time.Millisecond)
		}
	}))

	app.Get("/transactions", func(c *fiber.Ctx) error {
		scp := colly.NewCollector()

		transactions := []Transaction{}

		scp.OnHTML("table.table tbody tr", func(e *colly.HTMLElement) {
			if e.ChildText("td") == "" {
				return
			}

			trx := Transaction{
				Date: e.ChildText("td:nth-child(1)"),
				IDTrx: e.ChildText("td:nth-child(2)"),
				Contact: e.ChildText("td:nth-child(3)"),
				Price: e.ChildText("td:nth-child(4)"),
				Status: e.ChildText("td:nth-child(5)"),

			}

			transactions = append(transactions, trx)
		})

		scp.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		scp.Visit("https://kiosbedul.id/payment")

		return c.JSON(transactions)
	})

	app.Use("/monitor", monitor.New(
		monitor.Config{
			Title: "KiosBedul v2 Engine",
			Refresh: 1 * time.Second,
		},
	))

	app.Static("/", "./views")

	app.Listen(":80")
}

