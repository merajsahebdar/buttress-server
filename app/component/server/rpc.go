/*
 * Copyright 2021 Meraj Sahebdar
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package servercomp

import (
	"context"
	"fmt"
	"net"

	"buttress.io/app/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RpcServer
type RpcServer struct {
	*grpc.Server
	Log *zap.Logger
}

// RpcComp
var RpcComp = fx.Options(
	fx.Provide(newRpc),
	fx.Invoke(registerLifecycle),
)

// newRpc
func newRpc() *RpcServer {
	// Init Logger
	log := config.Log.Named("server-rpc")

	// Init the gRPC server.
	server := grpc.NewServer()

	return &RpcServer{
		Server: server,
		Log:    log,
	}
}

// registerLifecycle
func registerLifecycle(lc fx.Lifecycle, server *RpcServer) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			host := config.Cog.App.Host
			port := config.Cog.App.Port

			lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				server.Log.Fatal("failed to serve the tcp listener", zap.Error(err))
			}

			server.Log.Info("starting the rpc connection...", zap.String("host", host), zap.Int("port", port))

			go func() {
				if err := server.Serve(lis); err != nil {
					server.Log.Fatal("failed to start the rpc server", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			server.GracefulStop()
			return nil
		},
	})
}
