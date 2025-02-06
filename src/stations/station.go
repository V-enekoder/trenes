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
