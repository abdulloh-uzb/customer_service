package grpcClient

import (
	"exam/customer_service/config"
	pbp "exam/customer_service/genproto/post"
	pbr "exam/customer_service/genproto/reyting"
	"fmt"

	"google.golang.org/grpc"
)

// GrpcClientI ...
type GrpcClientI interface {
	Post() pbp.PostServiceClient
	Ranking() pbr.RankingServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg            config.Config
	postService    pbp.PostServiceClient
	rankingService pbr.RankingServiceClient
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	con, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("post service dial host:%s port:%d", cfg.PostServiceHost, cfg.PostServicePort)
	}

	conRanking, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.RankingServiceHost, cfg.RankingServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("ranking service dial host:%s port:%d", cfg.RankingServiceHost, cfg.RankingServicePort)
	}

	return &GrpcClient{
		cfg:            cfg,
		postService:    pbp.NewPostServiceClient(con),
		rankingService: pbr.NewRankingServiceClient(conRanking),
	}, nil
}

func (g *GrpcClient) Post() pbp.PostServiceClient {
	return g.postService
}

func (g *GrpcClient) Ranking() pbr.RankingServiceClient {
	return g.rankingService
}
