package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fumist23/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	client := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "fumist23",
		Title:    "satofumi story",
		Content:  "forst content of blog",
	}
	res, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})

	if err != nil {
		log.Fatalf("Failed to create blog.")
	}

	fmt.Printf("Blog has been created! blog: %v", res.GetBlog())
}
