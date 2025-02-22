/*
 * Copyright 2021 CloudWeGo Authors
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

package generic

import (
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

// BinaryGrpcGeneric raw grpc binary Generic
func BinaryGrpcGeneric(grpcSvcName string) Generic {
	return &binaryGrpcGeneric{grpcSvcName: grpcSvcName}
}

type binaryGrpcGeneric struct {
	grpcSvcName string
}

func GetGrpcSvcName(g Generic) string {
	binaryG, ok := g.(*binaryGrpcGeneric)
	if !ok {
		return ""
	}
	return binaryG.grpcSvcName
}

func (g *binaryGrpcGeneric) Framed() bool {
	return true
}

func (g *binaryGrpcGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	return serviceinfo.Protobuf
}

func (g *binaryGrpcGeneric) PayloadCodec() remote.PayloadCodec {
	return nil
}

func (g *binaryGrpcGeneric) GetMethod(req interface{}, method string) (*Method, error) {
	return &Method{method, false}, nil
}

func (g *binaryGrpcGeneric) Close() error {
	return nil
}
