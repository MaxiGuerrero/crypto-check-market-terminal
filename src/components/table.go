package components

import (
	"crypto-market-terminal/src/backend/models"
	"fmt"

	"github.com/rivo/tview"
)

const (
	RANK int = iota
	SYMBOL
	NAME
	PRICE
	CHANGE_PERCENT_24h
	VOLUME24H
)

func NewTable() *tview.Table {
	table := tview.NewTable()
	table.SetBorder(true)
	table.SetBorders(true)
	table.SetTitle("Cotización de cripto mercado")
	return table
}

func AddDataToTable(table *tview.Table, data []models.Currency) {
	setHeaderTable(table)
	for rowIdx, currency := range data {
		table.SetCell(rowIdx+1, RANK,
			tview.NewTableCell(currency.Rank).
				SetAlign(tview.AlignCenter),
		)
		table.SetCell(rowIdx+1, SYMBOL,
			tview.NewTableCell(currency.Symbol).
				SetAlign(tview.AlignCenter),
		)
		table.SetCell(rowIdx+1, NAME,
			tview.NewTableCell(currency.Name).
				SetAlign(tview.AlignCenter),
		)
		table.SetCell(rowIdx+1, PRICE,
			tview.NewTableCell(currency.Price).
				SetAlign(tview.AlignCenter),
		)
		table.SetCell(rowIdx+1, CHANGE_PERCENT_24h,
			tview.NewTableCell(fmt.Sprintf("%%%v", currency.ChangePercent24H)).
				SetAlign(tview.AlignCenter),
		)
		table.SetCell(rowIdx+1, VOLUME24H,
			tview.NewTableCell(currency.Volume24H).
				SetAlign(tview.AlignCenter),
		)
	}
}

func setHeaderTable(table *tview.Table) {
	table.SetCell(0, RANK, tview.NewTableCell("Rank").SetAlign(tview.AlignCenter))
	table.SetCell(0, SYMBOL, tview.NewTableCell("Symbol").SetAlign(tview.AlignCenter))
	table.SetCell(0, NAME, tview.NewTableCell("Name").SetAlign(tview.AlignCenter))
	table.SetCell(0, PRICE, tview.NewTableCell("Price").SetAlign(tview.AlignCenter))
	table.SetCell(0, CHANGE_PERCENT_24h, tview.NewTableCell("Percent variation price 24h").SetAlign(tview.AlignCenter))
	table.SetCell(0, VOLUME24H, tview.NewTableCell("Volume24H").SetAlign(tview.AlignCenter))
}
