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