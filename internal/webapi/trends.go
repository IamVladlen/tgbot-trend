package webapi

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/IamVladlen/trend-bot/internal/entity"
)

const (
	_trendsUrl = "https://trends.google.com/trends/trendingsearches/daily/rss?geo="
)

type WebAPI struct{}

func New() *WebAPI {
	return &WebAPI{}
}

func (w *WebAPI) GetTrends(country string) (entity.Trends, error) {
	var trends entity.Trends

	res, err := http.Get(_trendsUrl + country)
	if err != nil {
		return entity.Trends{}, fmt.Errorf("webapi - GetTrends: %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.Trends{}, fmt.Errorf("webapi - GetTrends: %w", err)
	}

	if err := xml.Unmarshal(data, &trends); err != nil {
		return entity.Trends{}, fmt.Errorf("webapi - GetTrends: %w", err)
	}

	return trends, nil
}
