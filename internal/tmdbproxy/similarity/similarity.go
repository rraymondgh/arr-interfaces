package similarity

import (
	"cmp"
	"slices"
	"time"

	api "github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/oapi"
	"github.com/rraymondgh/arr-interfaces/internal/tmdbproxy/searchmodel"

	"github.com/hbollon/go-edlib"
	"github.com/mozillazg/go-unidecode"
	"github.com/rraymondgh/arr-interfaces/internal/regex"
)

type SimilarityHelper struct{}

type similarityResult = int

const (
	similarityCalc similarityResult = iota
	similarityMeili
	similarityNoMatch
)

func (a SimilarityHelper) CalcSimilarity(query string, name string, score float64) searchmodel.Similarity {
	modQuery := regex.NormalizeString(unidecode.Unidecode(query))

	sim := searchmodel.Similarity{
		Score: score,
	}
	name = regex.NormalizeString(unidecode.Unidecode(name))

	sim.Distance.Levenshtein = edlib.LevenshteinDistance(name, modQuery)

	switch {
	case sim.Distance.Levenshtein == 0:
		sim.Measure.OSADamerauLevenshtein = 1
		sim.Measure.Lcs = 1
		sim.Measure.Cosine = 1
		sim.Measure.Jaccard = 1
		sim.Measure.SorensenDice = 1
		sim.Measure.Qgram = 1
	case sim.Distance.Levenshtein > 25:
		sim.Measure.OSADamerauLevenshtein = 0
		sim.Measure.Lcs = 0
		sim.Measure.Cosine = 0
		sim.Measure.Jaccard = 0
		sim.Measure.SorensenDice = 0
		sim.Measure.Qgram = 0
	default:
		res, _ := edlib.StringsSimilarity(name, modQuery, edlib.OSADamerauLevenshtein)
		sim.Measure.OSADamerauLevenshtein = res

		res, _ = edlib.StringsSimilarity(name, modQuery, edlib.Lcs)
		sim.Measure.Lcs = res

		res, _ = edlib.StringsSimilarity(name, modQuery, edlib.Cosine)
		sim.Measure.Cosine = res

		res, _ = edlib.StringsSimilarity(name, modQuery, edlib.Jaccard)
		sim.Measure.Jaccard = res

		res, _ = edlib.StringsSimilarity(name, modQuery, edlib.SorensenDice)
		sim.Measure.SorensenDice = res

		res, _ = edlib.StringsSimilarity(name, modQuery, edlib.Qgram)
		sim.Measure.Qgram = res
	}

	m := sim.Measure
	arr := []float32{
		m.Cosine,
		m.Jaccard,
		m.SorensenDice,
		m.Qgram,
		m.OSADamerauLevenshtein,
		m.Lcs,
	}
	slices.SortStableFunc(arr, func(a, b float32) int {
		return cmp.Compare(b, a)
	})
	l := len(arr)
	if l%2 == 0 {
		sim.Summary.Median = (arr[l/2] + arr[l/2+1]) / 2
	} else {
		sim.Summary.Median = arr[l/2]
	}
	sim.Summary.Min = arr[0]
	sim.Summary.Max = arr[0]
	for i := 0; i < len(arr); i++ {
		sim.Summary.Mean += arr[i]
		if arr[i] < sim.Summary.Min {
			sim.Summary.Min = arr[i]
		}
		if arr[i] > sim.Summary.Max {
			sim.Summary.Max = arr[i]
		}
	}
	sim.Summary.Mean = sim.Summary.Mean / float32(len(arr))
	return sim
}

func (a SimilarityHelper) goodMatch(sim *searchmodel.Similarity, result *api.FindBy, yearOK bool) similarityResult {
	if yearOK &&
		sim.Summary.Median > .5 &&
		sim.Summary.Min > .4 &&
		sim.Summary.Max > 0.8 {
		return similarityCalc
	}

	if yearOK &&
		result.RankingScore > 0.9 &&
		result.RankingScoreDetails.Attribute.QueryWordDistanceScore > .9 &&
		result.RankingScoreDetails.Exactness.MatchType == "noExactMatch" &&
		*result.RankingScoreDetails.Exactness.MatchingWords <= *result.RankingScoreDetails.Exactness.MaxMatchingWords &&
		*result.RankingScoreDetails.Exactness.MatchingWords >= (*result.RankingScoreDetails.Exactness.MaxMatchingWords-1) &&
		sim.Summary.Max >= .59 {
		return similarityMeili
	}

	return similarityNoMatch
}

func (a SimilarityHelper) yearOK(year int32, date string, score float64) bool {
	t, _ := time.Parse(time.DateOnly, date)
	yearOK := int32(t.Year()) == year
	if !yearOK && score > .97 {
		yearOK = int32(t.Year()) >= year-1 && int32(t.Year()) <= year+1
	}

	return yearOK
}

func (a SimilarityHelper) BestMatch(query string, mediaType string, year int32, results []*api.FindBy) int {
	type goodSim struct {
		index int
		match similarityResult
		sim   searchmodel.Similarity
	}
	var passedSimilarity []goodSim
	var sim searchmodel.Similarity
	for index, result := range results {
		var name *string
		if mediaType == "tv" {
			name = result.Name
		} else {
			name = result.Title
		}
		sim = a.CalcSimilarity(query, *name, result.RankingScore)

		yearOK := true
		if result.MediaType == "movie" && year != -1 && result.ReleaseDate != nil {
			yearOK = a.yearOK(year, *result.ReleaseDate, result.RankingScore)
		} else if result.MediaType == "tv" && year != -1 && result.FirstAirDate != nil {
			yearOK = a.yearOK(year, *result.FirstAirDate, result.RankingScore)
		}
		match := a.goodMatch(&sim, result, yearOK)
		if match <= similarityMeili {
			passedSimilarity = append(passedSimilarity, goodSim{index: index, sim: sim, match: match})
		}

	}

	if len(passedSimilarity) > 1 {
		slices.SortStableFunc(passedSimilarity, func(a, b goodSim) int {
			return cmp.Compare(b.sim.Summary.Mean, b.sim.Summary.Mean)
		})
	}

	if len(passedSimilarity) >= 1 {
		return passedSimilarity[0].index
	}

	return -1
}
