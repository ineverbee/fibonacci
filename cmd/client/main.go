package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	pb "github.com/ineverbee/fibonacci/gen/go/pb"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatal("pass 2 arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	for x < 0 || y < 0 {
		log.Println("arguments should be >= 0")
		log.Println("Enter 2 non-negative arguments:")
		fmt.Scanf("%d %d", &x, &y)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewFibonacciClient(conn)
	res, err := c.Fibo(context.Background(), &pb.IntRequest{X: uint32(x), Y: uint32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
