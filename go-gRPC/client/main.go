package main

import (
	"context"
	"io"
	"log"

	"github.com/kopher1601/go-studies/go-gRPC/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewFileServiceClient(conn)
	callDownload(c)
	//resp, err := c.ListFiles(context.Background(), &pb.ListFilesRequest{})
	//if err != nil {
	//	log.Fatalf("failed to list files: %v", err)
	//}
	//
	//log.Println(resp)
}

func callDownload(client pb.FileServiceClient) {
	req := &pb.DownloadRequest{Filename: "name.txt"}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to download: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to download: %v", err)
		}

		log.Printf("Response from Download(bytes): %v", res.GetData())
		log.Printf("Response from Download(string): %v", string(res.GetData()))
	}
}
