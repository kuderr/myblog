package comments

import "time"

type Comment struct {
	postId      int
	userId      int
	body        string
	dateCreated time.Time
	dateUpdated time.Time
}
