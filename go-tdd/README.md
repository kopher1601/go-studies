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

## Pointer

Goでは、関数またはメソッドを呼び出すと、引数は _ コピーされます。

`func (w Wallet) Deposit(amount int)`を呼び出すとき、 `w`はメソッドの呼び出し元のコピーです。

## Maps

キーのタイプは特別です。 2つのキーが等しいかどうかを判別できないと、正しい値が取得されていることを確認する方法がないため、比較可能な型にしかできません。

- https://go.dev/ref/spec#Comparison_operators

マップの興味深い特性は、マップをポインタとして渡さなくても変更できることです。これは、 mapが参照型であるためです。つまり、ポインタのように、基礎となるデータ構造への参照を保持します。 基本的なデータ構造はhash tablesまたはhash mapである

参照型がもたらす落とし穴は、マップがnil値になる可能性があることです。 nilマップは読み取り時に空のマップのように動作しますが、nilマップに書き込もうとすると、ランタイムパニックが発生します。
したがって、空のマップ変数を初期化しないでください。

```go
// 🙅 , m は nil マップ
var m map[string]string


// 🙆
var dictionary = map[string]string{}

// 🙆
var dictionary = make(map[string]string)
```

値がすでに存在する場合、マップはエラーをスローしません。 代わりに、先に進み、新しく提供された値で値を上書きします。