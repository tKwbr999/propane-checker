package conf

// Prefecture は都道府県名の日本語名と英語名を格納します。
type Prefecture struct {
	Jp string
	En string
}

// PrefectureJp は都道府県名の日本語名を返します。
func (p Prefecture) PrefectureJp() string {
	return p.Jp
}

// PrefectureEn は都道府県名の英語名を返します。
func (p Prefecture) PrefectureEn() string {
	return p.En
}

var CodePrefectures = map[string]Prefecture{
	"01": {"北海道", "Hokkaido"},
	"02": {"青森県", "Aomori"},
	"03": {"岩手県", "Iwate"},
	"04": {"宮城県", "Miyagi"},
	"05": {"秋田県", "Akita"},
	"06": {"山形県", "Yamagata"},
	"07": {"福島県", "Fukushima"},
	"08": {"茨城県", "Ibaraki"},
	"09": {"栃木県", "Tochigi"},
	"10": {"群馬県", "Gunma"},
	"11": {"埼玉県", "Saitama"},
	"12": {"千葉県", "Chiba"},
	"13": {"東京都", "Tokyo"},
	"14": {"神奈川県", "Kanagawa"},
	"15": {"新潟県", "Niigata"},
	"16": {"富山県", "Toyama"},
	"17": {"石川県", "Ishikawa"},
	"18": {"福井県", "Fukui"},
	"19": {"山梨県", "Yamanashi"},
	"20": {"長野県", "Nagano"},
	"21": {"岐阜県", "Gifu"},
	"22": {"静岡県", "Shizuoka"},
	"23": {"愛知県", "Aichi"},
	"24": {"三重県", "Mie"},
	"25": {"滋賀県", "Shiga"},
	"26": {"京都府", "Kyoto"},
	"27": {"大阪府", "Osaka"},
	"28": {"兵庫県", "Hyogo"},
	"29": {"奈良県", "Nara"},
	"30": {"和歌山県", "Wakayama"},
	"31": {"鳥取県", "Tottori"},
	"32": {"島根県", "Shimane"},
	"33": {"岡山県", "Okayama"},
	"34": {"広島県", "Hiroshima"},
	"35": {"山口県", "Yamaguchi"},
	"36": {"徳島県", "Tokushima"},
	"37": {"香川県", "Kagawa"},
	"38": {"愛媛県", "Ehime"},
	"39": {"高知県", "Kochi"},
	"40": {"福岡県", "Fukuoka"},
	"41": {"佐賀県", "Saga"},
	"42": {"長崎県", "Nagasaki"},
	"43": {"熊本県", "Kumamoto"},
	"44": {"大分県", "Oita"},
	"45": {"宮崎県", "Miyazaki"},
	"46": {"鹿児島県", "Kagoshima"},
	"47": {"沖縄県", "Okinawa"},
}
