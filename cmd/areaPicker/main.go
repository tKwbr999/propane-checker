package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"propane-checker/conf/area"
	"propane-checker/conf/area/city"
	"text/template"
)

const cityTemplate = `package city

var PrefectureCode{{.PrefectureCode}} = Prefecture{
	PrefectureCode: "{{.PrefectureCode}}",
	Cities: []City{
		{{- range .Cities}}
		{
			Code: "{{.Code}}",
			Name: "{{.Name}}",
		},
		{{- end}}
	},
}
`

func fetchCities(prefectureCode string) (*city.Prefecture, error) {
	url := fmt.Sprintf("https://www.land.mlit.go.jp/webland/api/CitySearch?area=%s", prefectureCode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var citiesResponse struct {
		Status string `json:"status"`
		Data   []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, err
	}
	var response city.Prefecture
	response.PrefectureCode = prefectureCode
	for _, c := range citiesResponse.Data {
		response.Cities = append(response.Cities, city.City{
			Code: c.ID,
			Name: c.Name,
		})
	}

	return &response, nil
}

func writeCitiesToFile(code string, p *city.Prefecture) error {
	tmpl, err := template.New("city").Parse(cityTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("./conf/area/city/prefectureCode%s.go", code))
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, p)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	for k, _ := range area.CodePrefecture {
		cities, err := fetchCities(k)
		if err != nil {
			panic(err)
			return
		}
		//citiesの内容を持ったstructを13.goに作成
		err = writeCitiesToFile(k, cities)
		if err != nil {
			panic(err)
		}
	}
}
