# go-tdd


goではファイルを指定するのではなくディレクトリを指定する必要がある
```shell
# ディレクトリ指定
go test ./hello-word/

# または全体テスト
go test ./...
```

## Formatting
```go
//  %q は値を二重引用符で囲む
t.Errorf("got %q want %q", got, want)
```