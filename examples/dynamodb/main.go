package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Widget struct {
	UserID int       `dynamo:",hash"`
	Time   time.Time `dynamo:",range"`

	Msg       string              `dynamo:"Message"`
	Count     int                 `dynamo:",omitempty"`
	Children  []Widget            // Lists
	Friends   []string            `dynamo:",set"`
	Set       map[string]struct{} `dynamo:",set"`
	SecretKey string              `dynamo:"-"`
}

func existTable(db *dynamo.DB, name string) (bool, error) {
	iter := db.ListTables().Iter()
	var tableName string
	for iter.Next(&tableName) {
		if tableName == name {
			return true, nil
		}
	}
	if err := iter.Err(); err != nil {
		return false, err
	}
	return false, nil
}

func createTable(db *dynamo.DB, name string) error {
	exist, err := existTable(db, name)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	return db.CreateTable(name, Widget{}).Run()

}

func getWidget(db *dynamo.DB, id int, t time.Time) (*Widget, error) {
	table := db.Table("Widgets")
	var result Widget
	if err := table.Get("UserID", id).
		Range("Time", dynamo.Equal, t).
		One(&result); err != nil {
		if !errors.Is(err, dynamo.ErrNotFound) {
			return nil, err
		}
	}
	return &result, nil
}

func putWidget(db *dynamo.DB, id int, w *Widget) error {
	table := db.Table("Widgets")
	return table.Put(w).Run()
}

func countWidget(db *dynamo.DB, id int, from, to time.Time, consistent bool) (int, error) {
	table := db.Table("Widgets")
	count, err := table.Get("UserID", id).
		Range("Time", dynamo.Between, from, to).
		Consistent(consistent).
		Count()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func dateRangeTosString(from, to time.Time) (string, string) {
	l := "2006-01-02"
	return from.Format(l), to.Format(l)
}

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{
		Region:   aws.String("us-west-2"),
		LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
		Logger: aws.LoggerFunc(func(args ...interface{}) {
			fmt.Println(args...)
		}),
	})

	if err := createTable(db, "Widgets"); err != nil {
		panic(err)
	}

	if err := putWidget(db, 1, &Widget{
		UserID: 1,
		Time:   time.Now(),
		Msg:    "Hello",
		Count:  1,
		Children: []Widget{
			{Msg: "World"},
		},
		Friends: []string{"Alice", "Bob"},
		Set: map[string]struct{}{
			"foo": {},
			"bar": {},
		},
	}); err != nil {
		panic(err)
	}

	count, err := countWidget(db, 1, time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, 1, 0), true)
	if err != nil {
		panic(err)
	}

	w, err := getWidget(db, 1, time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(count, w)
}
