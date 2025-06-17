# Go で AWS SES を利用してメールを送信してみる

### 使用技術
- `go`
- `aws-ses-v2-local`
- `aws-sdk-go-v2`

### 動作確認
1. Dockerの起動
```shell
docker compose up -d
```
2. メール送信プログラム実行
```shell
go run main.go
```
3. ブラウザで確認
- `http://localhost:8005/`　で確認できる

