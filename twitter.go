package twitter

import (
	"encoding/json"
	"time"
)

// `Tweet` represents the data we care about from tweets
type Tweet struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

// todo
func (t Tweet) Serialize() *[]byte {
	b, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return &b
}

// `User` contains information on the account which creates tweets.
type User struct {
	Id string
	Name string
	ScreenName string
	CreatedAt time.Time
}

// todo
func (u User) Serialize() *[]byte {
	b, err := json.Marshal(u)
	if err != nil {
		return nil
	}
	return &b
}

// todo: Versioning on these models?
// todo: json models? Wrapper with version on each object?
