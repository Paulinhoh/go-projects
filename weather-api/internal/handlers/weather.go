package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"weather-api/internal/cache"
	"weather-api/internal/models"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func GetWeather(city string) models.Weather {
	// carrega o arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// verificando se ja existe essa requisição no cache. Se sim a retorna
	cached, err := cache.Get(city)
	if err != redis.Nil && err != nil {
		log.Fatal(err.Error())
	}
	if err == nil {
		var weather models.Weather
		json.Unmarshal([]byte(cached), &weather)
		return weather
	}

	// URL da API que será consumida
	city = strings.ReplaceAll(strings.TrimSpace(city), " ", "+")
	key := os.Getenv("API_KEY")
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + city + "?key=" + key

	// 1. Faz a requisição GET
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// 3. Verifica se o status HTTP é sucesso (200)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API retornou status code: %d", resp.StatusCode)
	}

	// 4. Decodifica o JSON para a struct Weather
	var weather models.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %v", err)
	}
	weather.Timestamp = time.Now()

	// adicionando a requisição ao cache
	jsonBytes, _ := json.Marshal(weather)
	if err := cache.Set(city, jsonBytes); err != nil {
		log.Fatal(err.Error())
	}

	// retornando o resultado
	return weather
}
