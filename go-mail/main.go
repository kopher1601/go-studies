package main

import (
	"context"
	"go-mail/internal/awsconfig"
	"go-mail/mail"
	"log"
)

func main() {
	ctx := context.Background()

	// 基本 AWSConfig を生成する
	cfg, err := awsconfig.NewDefaultAWSConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	// メールインスタンスを生成する
	m := mail.NewMail(cfg)

	// メールを送信する
	err = m.SendEmail(ctx, "Test mail subject", "Test mail body", "kopher@kopher.com", "kopher@kopher.co.jp")
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
	}
}
