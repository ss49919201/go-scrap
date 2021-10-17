```
// 途中で失敗したら一件も書き込みできない
tx := db.WriteTx()
for i := 0; i < 25; i++ {
	if i == 0 {
		tx.Put(table.Put(&struct{}{}))
		continue
	}
	tx.Put(table.Put(&User{UserID: strconv.Itoa(i), Name: "太郎", Age: 20}))
}
if err := tx.Run(); err != nil {
	log.Println("failed write tx: ", err)
}

var users []*User
if err := table.Scan().All(users); err != nil {
	log.Fatal(err)
}
fmt.Println(len(users)) // 0
```

```
// テーブルの指定
table := db.Table("UserTable")

// Putとdelete一緒にできる
batchWrite := table.Batch("UserID", "Name").Write()
for i := 0; i < 26; i++ {
	if i == 0 {
		batchWrite.Delete(dynamo.Keys{"27", "太郎"})
		continue
	}
	batchWrite.Put(&User{UserID: strconv.Itoa(i), Name: "太郎", Age: 20})
}
wrote, err := batchWrite.Run()
if err != nil {
	log.Println("failed write tx: ", err)
}
fmt.Println("wrote: ", wrote)

var users []*User
if err := table.Scan().All(&users); err != nil {
	log.Fatal(err)
}
fmt.Println("scan: ", len(users))
```