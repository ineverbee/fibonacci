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
	x, y := req.X, req.Y
	if x > y {
		x, y = y, x
	}
	if y > 47 {
		y = 47
	}
	res := make([]uint32, y+1)
	if y < 2 {
		res[1] = 1
		for j, v := range res {
			if _, err := s.Cache.Get(uint32(j)); err != nil && (j == 0 || v != 0) {
				s.Cache.Set(uint32(j), v, time.Hour*24)
			}
		}
		return &pb.IntSliceResponse{Result: res[x : y+1]}, nil
	}
	for i := x; i <= y; i++ {
		if val, err := s.Cache.Get(i); err == nil {
			res[i] = uint32(val)
			continue
		}
		if (i - x) >= 2 {
			fibo.Fibo(i, y, &res)
		} else if val, err := s.Cache.Get(i - 2); err == nil && (i-x) == 1 {
			res[i-2] = uint32(val)
			fibo.Fibo(i, y, &res)
		} else {
			res[1] = 1
			fibo.Fibo(2, y, &res)
		}
		break
	}
	for j, v := range res {
		if _, err := s.Cache.Get(uint32(j)); err != nil && (j == 0 || v != 0) {
			s.Cache.Set(uint32(j), v, time.Hour*24)
		}
	}
	return &pb.IntSliceResponse{Result: res[x : y+1]}, nil
}
