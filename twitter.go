package twitter

import (
	"time"
)

// `Tweet` represents the data we care about from tweets
type Tweet struct {
	Id string
	UserId string
	Text string
	CreatedAt time.Time
}

// `User` contains information on the account which creates tweets.
type User struct {
	Id string
	Name string
	ScreenName string
	CreatedAt time.Time
}

// todo: Versioning on these models?
// todo: json models? Wrapper with version on each object?
