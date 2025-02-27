package searchmodel

import (
	"encoding/json"
	"fmt"
)

func (o Tmdb) PrimaryKey() string {
	return fmt.Sprintf("%s%v", o.MediaType, o.ID)
}

func (m *SearchResponse) Decode(results interface{}) error {
	if m.error != nil {
		return m.error
	}
	var b []byte
	var hits any
	if m.all {
		hits = m.Hits
	} else {
		if len(m.Hits) == 0 {
			return ErrNoDocuments
		}
		hits = m.Hits[0]
	}
	b, m.error = json.Marshal(hits)
	if m.error != nil {
		return m.error
	}
	m.error = json.Unmarshal(b, &results)
	if m.error != nil {
		return m.error
	}

	return nil

}

func (m *SearchResponse) Err() error {
	return m.error
}
