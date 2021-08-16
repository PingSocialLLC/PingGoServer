package models

type ShareGeoPing struct {
	ID      []string   `json:"ids" db:"ids" binding:"max=15,min=1"`
	Creator *UserBasic `db:"creator"`
	PingID  string     `json:"pingId" db:"ping_id" binding:"uuid4"`
}

type CreateGeoPing struct {
	Creator   *UserBasic `db:"creator"`
	SentMess  string     `json:"sentMessage" db:"sent_message" binding:"ascii,min=1.max=140"`
	Location  *Location  `json:"location" db:"location"`
	IsPrivate bool       `json:"isPrivate" db:"is_private"`
	TimeLimit int64      `json:"timeLimit" db:"time_limit" binding:"oneof=5 60 1440"`
}
