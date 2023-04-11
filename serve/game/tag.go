package game

type Tag struct {
	RowIndex int    `json:"RowIndex" from:"RowIndex"`
	ColIndex int    `json:"ColIndex" from:"ColIndex"`
	Name     string `json:"Name" from:"Name"`
	Uid      string `json:"Uid" from:"Uid"`
	Style    Style  `json:"Style" from:"Style"`
	Status   int    `json:"Status" from:"Status"`
}
type Style struct {
	Top     int     `json:"top" from:"top"`
	Left    int     `json:"left" from:"left"`
	Opacity float64 `json:"opacity" from:"opacity"`
}
