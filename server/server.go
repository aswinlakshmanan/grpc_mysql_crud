package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

var repo repository.ArticleRepository

func init() {
	repo = repository.NewMySQL()
}

func (*server) CreateArticle(ctx context.Context, req *blogpb.CreateArticleRequest) (*blogpb.CreateArticleResponse, error) {
	article := req.GetArticle()

	newArticle := model.Article{
		Title: article.GetTitle(),
		Body:  article.GetBody(),
	}

	res, err := repo.Create(ctx, newArticle)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error : %v", err))

	}

	return &blogpb.CreateArticleResponse{
		Article: &blogpb.Article{
			Id:    res,
			Title: article.GetTitle(),
			Body:  article.GetBody(),
		},
	}, nil

}

func (*server) ListArticle(req *blogpb.ListArticleRequest, stream blogpb.ServiceName_ListArticleServer) error {
	var perPage int64 = 10
	var start, end int64

	if req.GetPage() < 1 {
		start = 0
	} else {
		start = (req.GetPage() - 1) * perPage
	}

	end = perPage

	la, err := repo.List(context.Background(), start, end)

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("internal error : %v", err))

	}

	for _, key := range la {
		stream.Send(&blogpb.ListArticleResponse{Article: &blogpb.Article{
			Id:    key.ID,
			Title: key.Title,
			Body:  key.Body,
		}})
	}
	return nil
}

func main() {
	list, err := net.Listen("tcp", "0.0.0.0:5000")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	blogpb.RegisterServiceNameServer(s, &server{})
	go func() {
		fmt.Println("Starting server..")
		if err := s.Serve(list); err != nil {
			log.Fatal(err)
		}
	}()

}
