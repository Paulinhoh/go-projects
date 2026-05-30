package main

import (
	"flag"
	"fmt"
	"log"
	"weather-api/internal/handlers"
)

func main() {
	search := flag.String("s", "", "busca a estatistica de tempo da cidade")
	flag.Parse()
	if *search == "" {
		log.Fatal("o parametro não pode ser vazio!")
	}

	weather := handlers.GetWeather(*search)
	fmt.Printf("[%s] - Address: %s | TimeZone: %s | Description: %s\n", weather.Timestamp.Format("02/01/2006-15:04:05"), weather.Address, weather.Timezone, weather.Description)

	// exemplo de uso:
	// ./api -s "rio de janeiro"
}
