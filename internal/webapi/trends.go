package webapi

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/IamVladlen/trend-bot/pkg/redisdb"
)

const (
	_trendsUrl = "https://trends.google.com/trends/trendingsearches/daily/rss?geo="

	_trendKeyPrefix = "trend:"
	_trendCacheTTL  = 1 * time.Hour
)

type WebAPI struct {
	cache *redisdb.DB
	log   *logger.Logger
}

// GetTrends returns trends from web api or cache as entity.Trends struct.
func (a *WebAPI) GetTrends(country string) (entity.Trends, error) {
	var trends entity.Trends

	if err := a.get(&trends, country); err == nil {
		return trends, nil
	}

	if err := a.fetch(&trends, country); err != nil {
		return entity.Trends{}, fmt.Errorf("webapi - GetTrends: %w", err)
	}

	a.set(country, trends)

	return trends, nil
}

// fetch gets trends from web api.
func (a *WebAPI) fetch(trends *entity.Trends, country string) error {
	res, err := http.Get(_trendsUrl + country)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(data, &trends); err != nil {
		return err
	}

	return nil
}

// get gets trends from cache.
func (a *WebAPI) get(trends *entity.Trends, country string) error {
	key := _trendKeyPrefix + country

	data, err := a.cache.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), &trends); err != nil {
		a.log.Error().Msg("webapi - GetTrends - get: " + err.Error())
	}

	return nil
}

// set saves fetched trends to cache.
func (a *WebAPI) set(country string, trend entity.Trends) {
	key := _trendKeyPrefix + country

	json, err := json.Marshal(trend)
	if err != nil {
		a.log.Error().Msg("webapi - GetTrends - set: " + err.Error())
	}

	a.cache.Set(context.Background(), key, json, _trendCacheTTL)
}

func New(cache *redisdb.DB, log *logger.Logger) *WebAPI {
	return &WebAPI{
		cache,
		log,
	}
}
