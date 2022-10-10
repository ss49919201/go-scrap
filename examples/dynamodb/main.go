package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type widget struct {
	UserID int       // Hash key, a.k.a. partition key
	Time   time.Time // Range key, a.k.a. sort key

	Msg       string              `dynamo:"Message"`    // Change name in the database
	Count     int                 `dynamo:",omitempty"` // Omits if zero value
	Children  []widget            // Lists
	Friends   []string            `dynamo:",set"` // Sets
	Set       map[string]struct{} `dynamo:",set"` // Map sets
	SecretKey string              `dynamo:"-"`    // Ignored
}

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	w, err := getWidget(db, 1, time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(w)
}

func getWidget(db *dynamo.DB, id int, t time.Time) (*widget, error) {
	table := db.Table("Widgets")
	var result widget
	if err := table.Get("UserID", id).
		Range("Time", dynamo.Equal, t).
		One(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func dateRangeTosString(from, to time.Time) (string, string) {
	l := "2006-01-02"
	return from.Format(l), to.Format(l)
}
