# go-tdd

## Formatting
```go
//  %q は値を二重引用符で囲む
t.Errorf("got %q want %q", got, want)

%v the value in a default format
when printing structs, the plus flag (%+v) adds field names
```

## Test

goではファイルを指定するのではなくディレクトリを指定する必要がある
```shell
# ディレクトリ指定
go test ./hello-word/

# または全体テスト
go test ./...

# bench
go test ./iteration/ -bench=.
```

### t.Helper()
`t.Helper()`は、このメソッドがヘルパーであることをテストスイートに伝えるために必要です。こうすることで、テストが失敗したときに報告される行番号は、テストヘルパーの中ではなく 呼び出された関数 の中を示します。

