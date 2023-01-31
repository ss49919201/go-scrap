```sh
% go run main.go
no tags

% go run --tags tags main.go
tags

% go test -v ./...
=== RUN   TestTags
    main_test.go:6: no tags
--- PASS: TestTags (0.00s)
PASS
ok      github.com/ss49919201/go-scrap/tags     0.008s
?       github.com/ss49919201/go-scrap/tags/internal    [no test files]

% go test --tags tags -v ./...
=== RUN   TestTags
    main_test.go:6: tags
--- PASS: TestTags (0.00s)
PASS
ok      github.com/ss49919201/go-scrap/tags     0.008s
?       github.com/ss49919201/go-scrap/tags/internal    [no test files]
```
