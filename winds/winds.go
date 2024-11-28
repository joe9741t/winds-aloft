package winds

import (
	"fmt"
	"io"
	"net/http"
)

type Wind struct {
	Airport string
	Ft0     string
	Ft3000  string
	Ft6000  string
	Ft9000  string
	Ft12000 string
}

type WindsAloft struct {
	Winds map[string]Wind
}

func (wa *WindsAloft) AddWind(wind Wind) {
	if wa.Winds == nil {
		wa.Winds = make(map[string]Wind)
	}
	wa.Winds[wind.Airport] = wind
}

func (wa *WindsAloft) FetchWindsData() (string, error) {
	resp, err := http.Get("https://aviationweather.gov/api/data/windtemp?region=all")
	if err != nil {
		fmt.Println("Error fetching data", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading data", err)
		return "", err
	}
	return string(body), nil
}
