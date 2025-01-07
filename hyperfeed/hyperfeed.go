package hyperfeed

import "encoding/json"

type (
	response struct {
		Articles []article `json:"articles"`
	}

	article struct {
		Name    string `json:"name"`
		Link    string `json:"link"`
		Summary string `json:"summary"`
	}

	HyperResponse interface {
		GetArticles() []article
	}
)

func (r response) GetArticles() []article {
	return r.Articles
}

func ParseJSON(data []byte) (HyperResponse, error) {
	var resp response
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
