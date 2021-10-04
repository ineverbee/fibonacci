package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ineverbee/fibonacci/gen/go/fibonacci"
	"github.com/ineverbee/fibonacci/internal/app/cache/rediscache"
	"google.golang.org/grpc"

	pb "github.com/ineverbee/fibonacci/gen/go/pb"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func run(cache *rediscache.Cache) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	err := pb.RegisterFibonacciHandlerServer(ctx, mux, &fibonacci.GRPCServer{Cache: cache})
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("Listening and serving on :8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	pool := newPool("redis:6379")
	cache := rediscache.New(pool, "FIBONACCI")

	go run(cache)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	srv := &fibonacci.GRPCServer{Cache: cache}
	pb.RegisterFibonacciServer(grpcServer, srv)
	log.Println("Listening and serving on :8080")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
