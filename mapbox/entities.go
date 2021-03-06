package mapbox

type (
	Feature struct {
		ID          string     `json:"id"`
		Type        string     `json:"type"`
		PlaceType   []string   `json:"place_type"`
		Relevance   float64    `json:"relevance"`
		Properties  Properties `json:"properties"`
		Text        string     `json:"text"`
		PlaceName   string     `json:"place_name"`
		Center      []float64  `json:"center"`
		Geometry    Geometry   `json:"geometry"`
		Address     string     `json:"address"`
		Context     []Context  `json:"context"`
		BoundingBox []float64  `json:"bbox"`
	}

	Properties struct {
		Accuracy  string `json:"accuracy"`
		ShortCode string `json:"short_code"`
	}

	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	Context struct {
		ID        string `json:"id"`
		Text      string `json:"text"`
		Wikidata  string `json:"wikidata"`
		ShortCode string `json:"short_code"`
	}
)
