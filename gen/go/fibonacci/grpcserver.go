package fibonacci

import (
	"context"
	"time"

	pb "github.com/ineverbee/fibonacci/gen/go/pb"
	"github.com/ineverbee/fibonacci/internal/app/cache/rediscache"
	"github.com/ineverbee/fibonacci/internal/app/fibo"
)

//GRPCServer ...
type GRPCServer struct {
	pb.UnimplementedFibonacciServer
	Cache *rediscache.Cache
}

func (s *GRPCServer) Fibo(ctx context.Context, req *pb.IntRequest) (*pb.IntSliceResponse, error) {
	res := []uint32{}
	x, y := req.X, req.Y
	if x > y {
		x, y = y, x
	}
	for i := x; i <= y; i++ {
		if val, err := s.Cache.Get(i); err == nil {
			res = append(res, uint32(val))
			continue
		}
		arr := fibo.Fibo(y)
		for j, v := range arr {
			if _, err := s.Cache.Get(uint32(j)); err != nil {
				s.Cache.Set(uint32(j), v, time.Hour*24)
			}
		}
		res = append(res, arr[i:y+1]...)
		break
	}
	return &pb.IntSliceResponse{Result: res}, nil
}
