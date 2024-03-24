package city

type Prefecture struct {
	PrefectureCode string `json:"prefectureCode"`
	Cities         []City `json:"data"`
}

type City struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
