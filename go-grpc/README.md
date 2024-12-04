# 🥽 Protocol Buffers

Goolgeによって2008年にオープンソース化された**スキーマ言語**

拡張子は`.proto`

> **スキーマ言語**
何かしらの処理をさせるのではなく、要素や属性などの構造を定義するための言語
>

## スキーマ言語がなぜ重要か

- 多くのシステムやデータが様々な技術で複数のサービスやストレージに分割されるようになってきている（モノリスからマイクロサービス）
- クライアント側もWebだけではなくiOSやAndroidなどのモバイル対応が必須になっている
- それぞれのシステムのインターフェースをすり合わせるのが大変
- こういった背景で事前にどういったデータをやり取りするのかを宣言的に定義しておく

## Protocol Buffersの特徴

- gRPCのデータフォーマットとして使用されている
- プログラミング言語からは独立しており、様々な言語に変更可能
- バイナリ形式にシリアライズするので、サイズが小さく高速な通信が可能
- 型安全にデータのやり取りが可能
- JSONに変換することも可能

## JSONとの比較

| **JSON** | **Protocol Buffers** |
| --- | --- |
| Webの世界では完全にスタンダード | JSONと比較するとまだまだ少数派 |
| あらゆる言語で読むことが可能 | 一部の言語では未対応 |
| 配列やネストなど、複雑な形式も扱うことが可能 | 複雑すぎる構造には不向き |
| ヒューマンリーダブル | バイナリに変換された後は人間には読めない |
| データスキーマを強制できない | 型が保証される |
| データサイズが大きくなりがち | データサイズは小さい |

## Protocol Buffersを使用した開発の進め方

1. スキーマの定義
2. 開発言語のオブジェクトを自動生成
3. バイナリ形式へシリアライズ

## 基本文法

proto2とproto3の文法の違いが大きい

```protobuf
syntax = "proto3";

// comment

/*

  multiline comment 
  
*/
```
## Message

- 複数のフィールドを持つことができる型定義
    - それぞれのフィールドはスカラ型もしくはコンポジット型
- 各言語のコードとしてコンパイルした場合、構造体やクラスとして変換される
- 一つのporotoファイルに複数のmessage型を定義することも可能

```protobuf
message Person {
	string name = 1;
	int32 id = 2;
	string email = 3;
}
```

- `=` は タグ番号を割り当てるだけ
    - `[フィールド型][フィールド名] = [タグ番号]`
    - `int32 id = 2;`

## Scalar Value Types
- https://protobuf.dev/programming-guides/proto3/#scalar

## Tag

- Protocol Buffersではフィールドはフィールド名ではなく、タグ番号によって識別される
- 重複は許されず、一意である必要がある
- タグの最小値１、最大値`2^29-1(536,870,911)`
- `19000 ~ 19999`はProtocol Buffersの予約番号のため使用不可
- 1 ~ 15 番までは `1byte` で表すことが出来るため、よく使うフィールドには1 ~ 15番を割り当てる
- タグは連番にする必要はないので、あまり使わないフィールドはあえて16番以後を割り当てることも可能
- タグ番号を予約するなど、安全にProtocol Buffersを使用する方法も用意されている

## Enum (列挙型)

- キーワードは`enum`を使う
- タグ番号が１からではなく０から始まる
- 番号のスキップはできない
- 0番は `_UNKNOWN` におくのが慣例

```protobuf
syntax = "proto3";

message Employee {
  int32 id = 1;
  string name = 2;
  string email = 3;
  Occupation occupation = 4;
}

enum Occupation {
  OCCUPATION_UNKNOWN = 0;
  ENGINEER = 1;
  DESIGNER = 2;
  MANAGER = 3;
}
```
## その他フィールド
```protobuf
syntax = "proto3";

message Employee {
  int32 id = 1;
  string name = 2;
  string email = 3;
  Occupation occupation = 4;
  repeated string phone_number = 5; // 配列
  map<string, Project> project = 6; // Map, Mapにはrepeatedを付けるのはできない
  oneof profile { // 複数のどれかの一つを持つ、またrepeatedを付けるのはできない
    string text = 7;
    Video video = 8;
  }
}

enum Occupation {
  OCCUPATION_UNKNOWN = 0;
  ENGINEER = 1;
  DESIGNER = 2;
  MANAGER = 3;
}

message Project {}
message Video {}
```
## デフォルト値

定義したmessageでデータをやり取りする際に、定義したフィールドがセットされていない場合、そのフィールドのデフォルト値が設定される。

デフォルト値は型によって決められている

- string : 空の文字列
- bytes : 空のbyte
- bool : false
- 整数型・浮動小数点数 : 0
- 列挙型 : タグ番号の0の値
- repeated : 空のリスト
## messageのnest
```protobuf

message Employee {
  int32 id = 1;
  map<string, Company.Project> project = 2;
}
// message nest
message Company {
  message Project {}
}
```
## Compile

- protoファイルのimport分のパスを特定する
- 複数の箇所からprotoファイルをインポートする必要がある場合、コロン区切りで複数のパスを記述することが可能
- `-I` オプションを省略した場合は、カレントディレクトリが設定される

```protobuf
// -I{PATH}, --proto_path={PATH}
	protoc -I./test:./dev 
```

### 各言語に変換するためのオプション

- オプションによって、どの言語に変換するかを決定する
- Go言語のオプションはプラグインで追加する必要がある

```protobuf
--go_out=OUT_DIR
--cpp_out=OUT_DIR
--ruby_out=OUT_DIR
```

### コンパイルするファイルの指定

```protobuf
// 複数指定
protoc -I. --go_out=. proto/employee.proto proto/date.proto
```
## JSON Mapping

```go
m := jsonpb.Marshaler{}
out, err := m.MarshalToString(employee) // json stringに変換
if err != nil {
	log.Fatalln("Can't marshal employee:", err)
}

readEmployee := &pb.Employee{}
if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
	log.Fatalln("Can't unmarshal employee:", err)
}
```
# gRPCとは

Googleによって2015年にオープンソース化されたRPC(Remote Procedure Call)のためのプロトコル

## RPC(Remote Procedure Call)とは

- Remote = 遠隔地 (リモート) サーバーの
- Procedure = 手続き (メソッド) を
- Call = 呼び出す (実行する)
- ネットワーク上の他の端末と通信するための仕組み
- REST APIのようにパスやメソッドを指定する必要はなく、メソッド名と引数を指定する
- gRPC以外にJSON-RPCなどがあるが、今はgRPCがデファクトスタンダード

## gRPCの特徴

- データフォーマットにProtocol Bufferを使用
  - バイナリにシリアライズすることで送信データ量が減り高速な通信を実現
  - 型付けされたデータ転送が可能
- IDL (Protocol Buffers) からサーバー側・クライアント側に必要なソースコードを生成
- 通信には HTTP/2 を使用
- 特定の言語やプラットフォームに依存しない

## gRPCが適したケース

- Microservice間の通信
  - 複数の言語やプラットフォームで構成される可能性がある
  - バックエンド間であればgRPCの恩恵が多く得られる
- モバイルユーザーが利用するサービス
  - 通信量が削減できるため、通信容量制限にかかりにくい
- 速度が求められる場合

## gRPCの開発の流れ

1. protoファイルの作成
2. protoファイルをコンパイルしてサーバー・クライアントの雛形(hinagata, 모형, 양식, 형식)コードを作成
3. 雛形コードを使用してサーバー・クライアントを実装
