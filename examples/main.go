package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/zerodivisi0n/exante-api-go"
)

func main() {
	clientID := os.Getenv("EXANTE_CLIENT_ID")
	applicationID := os.Getenv("EXANTE_APPLICATION_ID")
	sharedKey := os.Getenv("EXANTE_SHARED_KEY")

	client := exante.NewClient(clientID, applicationID, sharedKey)

	writeExchanges(client)
	writeSymbols(client)
}

func writeExchanges(client *exante.Client) {
	exchanges, err := client.Exchanges()
	if err != nil {
		fmt.Printf("Failed to get exchanges: %v\n", err)
		return
	}
	f, err := os.Create("exante-exchanges.csv")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"ID", "Name", "Country"})
	for _, ex := range exchanges {
		if err := w.Write([]string{ex.ID, ex.Name, ex.Country}); err != nil {
			fmt.Printf("Write failed: %v\n", err)
			break
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func writeSymbols(client *exante.Client) {
	symbols, err := client.Symbols()
	if err != nil {
		fmt.Printf("Failed to get symbols: %v\n", err)
		return
	}
	f, err := os.Create("exante-symbols.csv")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Write([]string{"ID", "Name", "Description", "Ticker", "Type", "Exchange", "Country", "Currency", "MPI", "Group"})
	for _, s := range symbols {
		if err := w.Write([]string{
			s.ID, s.Name,
			s.Description,
			s.Ticker,
			s.Type,
			s.Exchange,
			s.Country,
			s.Currency,
			strconv.FormatFloat(s.MPI, 'g', -1, 64),
			s.Group}); err != nil {
			fmt.Printf("Write failed: %v\n", err)
			break
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
