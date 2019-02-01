package function

type Bddata struct {
	Status int `json:"status"`
	Data   struct {
		UserIndexes []struct {
			Word string `json:"word"`
			All  struct {
				StartDate string `json:"startDate"`
				EndDate   string `json:"endDate"`
				Data      string `json:"data"`
			} `json:"all"`
			Pc struct {
				StartDate string `json:"startDate"`
				EndDate   string `json:"endDate"`
				Data      string `json:"data"`
			} `json:"pc"`
			Wise struct {
				StartDate string `json:"startDate"`
				EndDate   string `json:"endDate"`
				Data      string `json:"data"`
			} `json:"wise"`
			Type string `json:"type"`
		} `json:"userIndexes"`
		GeneralRatio []struct {
			Word string `json:"word"`
			All  struct {
				Avg int `json:"avg"`
				Yoy int `json:"yoy"`
				Qoq int `json:"qoq"`
			} `json:"all"`
			Pc struct {
				Avg int `json:"avg"`
				Yoy int `json:"yoy"`
				Qoq int `json:"qoq"`
			} `json:"pc"`
			Wise struct {
				Avg int `json:"avg"`
				Yoy int `json:"yoy"`
				Qoq int `json:"qoq"`
			} `json:"wise"`
		} `json:"generalRatio"`
		Uniqid string `json:"uniqid"`
	} `json:"data"`
	Message int `json:"message"`
}

type BdKey struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
}

type bd_index_info struct {
	Provinces map[string]string              `json:"provinces"`
	CityShip  map[string][]map[string]string `json:"cityShip"`
}
