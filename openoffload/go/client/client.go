// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Intel Corporation, or its subsidiaries.

package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50151", "the address to connect to")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	do_sessionoffload(conn, ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go do_client_background(conn, &wg)
	wg.Wait()
}
