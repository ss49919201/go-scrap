```
fmt.Println("project name ?")
var projectName string
fmt.Scan(&projectName) // 入力が入る。スペースまたは改行を入れると次の引数に入る
fmt.Printf("project name is %s", projectName)
```