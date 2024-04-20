package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"propane-checker/conf"
	"propane-checker/conf/area"
	"text/template"
)

const cityTemplate = `package area

var Area{{.Code}}{{.En}} =  []Area{
	{{- range .List}}
		{
			ID: "{{.ID}}",
			Name: "{{.Name}}",
		},
	{{- end}}
}
`

func fetchCities(code string) ([]*area.Area, error) {
	url := fmt.Sprintf("https://www.reinfolib.mlit.go.jp/ex-api/external/XIT002?area=%s", code)

	// HTTPSのリクエストを発行する
	req, err := http.NewRequest("GET", url, nil)
	// APIキーをリクエストヘッダーOcp-Apim-Subscription-Keyに設定する
	req.Header.Set("Ocp-Apim-Subscription-Key", os.Getenv("REINFOLIB_API_KEY"))

	// リクエストを送信する
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var citiesResponse struct {
		Status string       `json:"status"`
		City   []*area.Area `json:"data"`
	}

	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, err
	}

	return citiesResponse.City, nil
}

func writeCitiesToFile(code, en string, list []*area.Area) error {
	tmpl, err := template.New("city").Parse(cityTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("./conf/area/area%s%s.go", code, en))
	if err != nil {
		return err
	}
	defer file.Close()

	var templateData = struct {
		Code string
		En   string
		List []*area.Area
	}{
		Code: code,
		En:   en,
		List: list,
	}

	err = tmpl.Execute(file, templateData)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	for k, v := range conf.CodePrefectures {
		cities, err := fetchCities(k)
		if err != nil {
			panic(err)
			return
		}
		err = writeCitiesToFile(k, v.PrefectureEn(), cities)
		if err != nil {
			panic(err)
		}
	}
}
