package model

type LineMessageUser struct {
	Events     []EventModel `json:"events"`
	Destinatin string       `json:"Destinatin"`
	Channal    string       `json:"Channal"`
}

type EventModel struct {
	Type       string       `json:"type"`
	ReplyToken string       `json:"replyToken"`
	Source     SourceModel  `json:"source"`
	Timestamp  int64        `json:"timestamp"`
	Message    MessageModel `json:"message"`
}

type SourceModel struct {
	UserID string `json:"userId"`
	Type   string `json:"type"`
}

type MessageModel struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ResFromAPI struct {
	DefaultMessage    string        `json:"defaultMessage"`
	Distance          float64       `json:"distance"`
	DistanceThreshold float64       `json:"distanceThreshold"`
	ObjList           []ObjectModel `json:"objList"`
}

type ObjectModel struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Channal struct {
	ID           string `json:"id" bson:"id"`
	ChannalToken string `json:"channaltoken" bson:"channaltoken"`
	Verifytoken  string `json:"verifytoken" bson:"verifytoken"`
	Url          string `json:"url"	bson:"url"`
}
