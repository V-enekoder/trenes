package station

type Station struct {
	ID          int64  `json:"station_id"`
	Name        string `json:"name"`
	Line        int64  `json:"line"`
	Typestation string `json:"type"`
	System      string `json:"system"`
}

type StationDTO struct {
	ID   int64  `json:"station_id"`
	Name string `json:"name"`
}

type OptimalPath struct {
	Path   []interface{} `json:"path"`
	Weight float64       `json:"weight"`
	Time   float64       `json:"time"`
}
