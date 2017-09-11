package prestoClient

type column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type queryResults struct {
	ID               string `json:"id"`
	InfoURI          string `json:"infoUri"`
	PartialCancelURI string `json:"partialCancelUri"`
	NextURI          string `json:"nextUri"`
	Stats            struct {
		Stats string `json:"state"`
	} `json:"stats"`
	QueryError struct {
		Message string `json:"message"`
	}
	Columns []column      `json:"columns"`
	Data    []interface{} `json:"data"`
}

func (qr queryResults) parse() string {
	return ""
}
