package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func main() {
	// CSVファイルを開く
	file, err := os.Open("./doc/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(file)

	// CSVデータを全て読み込む
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var rows strings.Builder
	// 各レコードをループし、各行の1列目をキーとし、その行の最後の列を値とする

	for _, record := range records[1:] {
		value, _ := strconv.Atoi(record[len(record)-1])

		//dataMap[record[0]] = value
		rows.WriteString(`"` + record[0] + `": ` + strconv.Itoa(value) + ",\n")
	}

	template := `package conf

var area = map[string]int{
	{{date}}
}
`

	// templateにjsonDataを埋め込む
	latest := strings.Replace(template, "{{date}}", rows.String(), 1)
	// latestで./conf/latest.goを作成
	file, err = os.Create("./conf/latest.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(latest)

	// latestの下から3行目の末尾にカンマをつける
	latest = strings.Replace(latest, "\n}", ",\n}", 1)

}
