/**
 * Tencent is pleased to support the open source community by making CL5 available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	polaris "github.com/polarismesh/grpc-go-polaris"
	"github.com/polarismesh/grpc-go-polaris/examples/common/pb"
)

var (
	listenPort int
)

// EchoRateLimitService gRPC echo service struct
type EchoRateLimitService struct {
}

// Echo gRPC testing method
func (h *EchoRateLimitService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Value: fmt.Sprintf("echo: %s", req.Value)}, nil
}

func main() {
	address := fmt.Sprintf("0.0.0.0:%d", listenPort)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to addr %s: %v", address, err)
	}

	srv, err := polaris.NewServer(polaris.WithServiceName("RateLimitEchoServerGRPC"),
		polaris.WithServerHost("127.0.0.1"),
		// 开启限流能力
		polaris.WithPolarisRateLimit(),
	)
	if err != nil {
		log.Fatalf("Failed to addr %s: %v", address, err)
	}

	pb.RegisterEchoServerServer(srv.Server, &EchoRateLimitService{})

	// 启动服务
	if err := srv.Serve(listen); nil != err {
		log.Printf("listen err: %v", err)
	}
}
