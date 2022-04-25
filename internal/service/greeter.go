package service

import (
	"context"
	"fmt"
	"io"

	v1 "helloword/api/helloworld/v1"
	"helloword/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

func (s *GreeterService) UploadFile(conn v1.Greeter_UploadFileServer) error {
	defer fmt.Println("finish")
	data := make([]byte, 0)
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			fmt.Printf("file: %d, content: %s", len(data), string(data))
			return conn.SendAndClose(&v1.UploadResponse{
				Message: "uploadSuccess",
			})
		}
		if err != nil {
			return err
		}
		data = append(data, req.Content...)
	}
}