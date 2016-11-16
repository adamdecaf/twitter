package twitter

import (
	"encoding/json"
	"time"
)

// todo(adam): Versioning on these models?

// `Tweet` represents the data we care about from tweets
type Tweet struct {
	Id string `json:"id"`
	UserId string `json:"userId"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

// Serialize converts a tweet to json.
func (t Tweet) Serialize() *[]byte {
	b, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return &b
}

func ReadTweet(b []byte) (Tweet, error) {
	t := Tweet{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return Tweet{}, err
	}
	return t, nil
}

// `User` contains information on the account which creates tweets.
type User struct {
	Id string
	Name string
	ScreenName string
	CreatedAt time.Time
}

// Serialize converts a user struct to json
func (u User) Serialize() *[]byte {
	b, err := json.Marshal(u)
	if err != nil {
		return nil
	}
	return &b
}

func ReadUser(b []byte) (User, error) {
	t := User{}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return User{}, err
	}
	return t, nil
}
