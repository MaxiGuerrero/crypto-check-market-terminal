package main

import (
	"crypto-market-terminal/src/backend"
	"crypto-market-terminal/src/backend/client"
	"crypto-market-terminal/src/components"
	"time"

	"github.com/rivo/tview"
)

func main() {
	config := backend.NewConfig()
	table := components.NewTable()
	service := backend.NewMarketService(
		client.NewClientHTTP(),
		config,
	)
	app := tview.NewApplication()
	go updateData(service, table, app)
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

// Update data table each X miliseconds. It must run in goroutine
func updateData(service *backend.MarketService, table *tview.Table, app *tview.Application) {
	defer func() {
		if r := recover(); r != nil {
			updateData(service, table, app)
		}
	}()
	for {
		data := service.RetrieveCryptoRate()
		app.QueueUpdateDraw(func() {
			table.Clear()
			components.AddDataToTable(table, *data)
		})
		time.Sleep(1 * time.Second)
	}
}
