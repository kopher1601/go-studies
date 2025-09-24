# go-tdd

- https://andmorefine.gitbook.io/learn-go-with-tests

## Formatting
```go
//  %q は値を二重引用符で囲む
t.Errorf("got %q want %q", got, want)

// v
%v the value in a default format
when printing structs, the plus flag (%+v) adds field names

// f
fは float64用で、 .2は小数点以下2桁を出力することを意味します。
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

## Array

Goでは、スライスで等号演算子を使うことはできません。got と want の各スライスを繰り返し処理して値を確認する関数を書くこともできますが、便利のために reflect.DeepEqual を使うと、2つの変数が同じであるかどうかを確認するのに便利です。

> 重要なのは、reflect.DeepEqualは「型安全」ではないことに注意することです。

```go
reflect.DeepEqual([]int{1,2}, []int{3,9})
```

