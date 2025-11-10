package ratings

import (
	"fmt"
	"time"
)

type Comment struct {
	comment string
	date    time.Time
}

func (c Comment) String() string {
	return fmt.Sprintf("Comment: %q, Date: %s", c.comment, c.date.Format("2006-01-02 15:04:05"))
}

type UserRating struct {
	uid      int
	rating   float64
	comments Comment
}

func (u UserRating) String() string {
	return fmt.Sprintf("User ID: %d, Rating: %.1f, %s", u.uid, u.rating, u.comments)
}

type Rating struct {
	id        int
	avgRating float64
	ratings   []UserRating
}

func NewRating(id int) *Rating {
	return &Rating{
		id:      id,
		ratings: []UserRating{},
	}
}

func (r Rating) String() string {
	return fmt.Sprintf("Product ID: %d, Avg Rating: %.1f, Ratings: %v", r.id, r.avgRating, r.ratings)
}

func (r *Rating) AddRating(uid int, rating float64, comment string) error {

	if rating < 1 || rating > 5 {
		return fmt.Errorf("Invalid rating: %f", rating)
	}

	newRating := UserRating{
		uid:    uid,
		rating: rating,
		comments: Comment{
			comment: comment,
			date:    time.Now(),
		},
	}

	r.ratings = append(r.ratings, newRating)

	var total float64
	for _, ur := range r.ratings {
		total += ur.rating
	}
	r.avgRating = total / float64(len(r.ratings))

	return nil
}
