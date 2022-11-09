package types

type PicInfo struct {
	Detail string        `json:"detail"`
	Count  int           `json:"count"`
	Tags   []interface{} `json:"tags"`
	Data   []struct {
		Artwork struct {
			Title string `json:"title"`
			Id    int    `json:"id"`
		} `json:"artwork"`
		Author struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"author"`
		SanityLevel int    `json:"sanity_level"`
		R18         bool   `json:"r18"`
		Page        int    `json:"page"`
		CreateDate  string `json:"create_date"`
		Size        struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"size"`
		Tags []string `json:"tags"`
		Urls struct {
			Original string `json:"original"`
			Large    string `json:"large"`
			Medium   string `json:"medium"`
		} `json:"urls"`
	} `json:"data"`
}
