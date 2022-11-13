package types

type TarotCard struct {
	Nhits int `json:"nhits"`
	Cards []struct {
		Name       string `json:"name"`
		NameShort  string `json:"name_short"`
		Value      string `json:"value"`
		ValueInt   int    `json:"value_int"`
		Suit       string `json:"suit"`
		Type       string `json:"type"`
		MeaningUp  string `json:"meaning_up"`
		MeaningRev string `json:"meaning_rev"`
		Desc       string `json:"desc"`
	} `json:"cards"`
}

type SendTarot struct {
	Name       string `json:"name"`
	MeaningUp  string `json:"meaning_up"`
	MeaningRev string `json:"meaning_rev"`
	Desc       string `json:"desc"`
}
