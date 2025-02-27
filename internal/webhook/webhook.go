package webhook

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/adrg/xdg"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/requestworker"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Arr struct {
	Analysis *requestworker.TmdbAnalysis
	Log      *zap.SugaredLogger
}

func (w Arr) SonarrClassifier(ctx context.Context) error {
	if !w.Analysis.Config.Sonarr.GenerateFlags.Required {
		return nil
	}

	type sonarrCal struct {
		SeriesId      int `json:"seriesId"`
		SeasonNumber  int `json:"seasonNumber"`
		EpisodeNumber int `json:"episodeNumber"`
	}
	type sonarrSeries struct {
		Id     int `json:"Id"`
		TmdbId int `json:"tmdbId"`
	}
	var series []*sonarrSeries
	var cal []*sonarrCal
	resp, err := resty.New().R().
		SetHeader("X-Api-Key", w.Analysis.Config.Sonarr.ApiKey).
		SetResult(&series).
		Get(fmt.Sprintf("%s/api/v3/series", w.Analysis.Config.Sonarr.Url))
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] %s", resp.StatusCode(), resp.Body())
	}
	resp, err = resty.New().R().
		SetHeader("X-Api-Key", w.Analysis.Config.Sonarr.ApiKey).
		SetQueryParam("start", time.Now().Add(-10*time.Hour*24).Format(time.DateOnly)).
		SetQueryParam("end", time.Now().Add(10*time.Hour*24).Format(time.DateOnly)).
		SetResult(&cal).
		Get(fmt.Sprintf("%s/api/v3/calendar", w.Analysis.Config.Sonarr.Url))
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] %s", resp.StatusCode(), resp.Body())
	}

	cmp := func(a, b *sonarrSeries) int { return a.Id - b.Id }

	slices.SortFunc(series, cmp)
	tmdbId := func(seriesId int) int {
		idx, ok := slices.BinarySearchFunc(series, &sonarrSeries{Id: seriesId}, cmp)
		if ok {
			return series[idx].TmdbId
		} else {
			return -1
		}
	}

	type classifier struct {
		FlagDefinitions map[string]string `yaml:"flag_definitions"`
		Flags           map[string][]string
	}

	ww := classifier{Flags: map[string][]string{}, FlagDefinitions: map[string]string{
		w.Analysis.Config.Sonarr.GenerateFlags.All:      "string_list",
		w.Analysis.Config.Sonarr.GenerateFlags.Calendar: "string_list",
	}}
	seasons := make(map[string]bool)
	sonarr := make([]string, 0)
	calendar := make([]string, 0)

	for _, calEntry := range cal {
		_tmdbId := tmdbId(calEntry.SeriesId)
		calendar = append(calendar, fmt.Sprintf(
			"%d_S%02dE%02d",
			_tmdbId,
			calEntry.SeasonNumber,
			calEntry.EpisodeNumber))
		seasons[fmt.Sprintf("%d_S%02d", _tmdbId, calEntry.SeasonNumber)] = true
	}

	calendar = append(calendar, slices.Collect(maps.Keys(seasons))...)

	for _, s := range series {
		if s.TmdbId > 0 {
			sonarr = append(sonarr, fmt.Sprintf("%d", s.TmdbId))
		}
	}
	ww.Flags[w.Analysis.Config.Sonarr.GenerateFlags.All] = sonarr
	ww.Flags[w.Analysis.Config.Sonarr.GenerateFlags.Calendar] = calendar
	d, err := yaml.Marshal(ww)
	if err != nil {
		return err
	}
	path, err := xdg.ConfigFile("bitmagnet/classifier_sonarr.yml")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, d, 0666)
	if err != nil {
		return err
	}

	return nil

}

func (w Arr) RadarrClassifier(ctx context.Context) error {
	if !w.Analysis.Config.Radarr.GenerateFlags.Required {
		return nil
	}

	type radarrMovie struct {
		TmdbId int `json:"tmdbId"`
	}
	var movies []*radarrMovie
	resp, err := resty.New().R().
		SetHeader("X-Api-Key", w.Analysis.Config.Radarr.ApiKey).
		SetResult(&movies).
		Get(fmt.Sprintf("%s/api/v3/movie", w.Analysis.Config.Radarr.Url))
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("[%d] %s", resp.StatusCode(), resp.Body())
	}

	type classifier struct {
		FlagDefinitions map[string]string `yaml:"flag_definitions"`
		Flags           map[string][]string
	}

	ww := classifier{Flags: map[string][]string{}, FlagDefinitions: map[string]string{
		w.Analysis.Config.Radarr.GenerateFlags.All: "string_list",
	}}
	radarr := make([]string, 0)

	for _, s := range movies {
		if s.TmdbId > 0 {
			radarr = append(radarr, fmt.Sprintf("%d", s.TmdbId))
		}
	}
	ww.Flags[w.Analysis.Config.Radarr.GenerateFlags.All] = radarr
	d, err := yaml.Marshal(ww)
	if err != nil {
		return err
	}
	path, err := xdg.ConfigFile("bitmagnet/classifier_radarr.yml")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, d, 0666)
	if err != nil {
		return err
	}

	return nil

}

func (w Arr) Event(c *gin.Context) {
	var event Event
	err := c.ShouldBind(&event)
	if err != nil {
		w.Log.Warn(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	w.Log.Infof("%+v", event)
	switch event.EventType {
	case SeriesAdd:
		err = w.Analysis.SourceAndStore(c, w.Analysis.Config.Sonarr, "tv")
		if err == nil {
			err = w.SonarrClassifier(c)
		}
	case MovieAdded:
		err = w.Analysis.SourceAndStore(c, w.Analysis.Config.Radarr, "movie")
		if err == nil {
			err = w.RadarrClassifier(c)
		}
	}
	if err != nil {
		w.Log.Warn(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
