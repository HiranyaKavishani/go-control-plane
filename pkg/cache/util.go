// Copyright 2017 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

// MessageToStruct encodes a protobuf Message into a Struct. Hilariously, it
// uses JSON as the intermediary
// author:glen@turbinelabs.io
func MessageToStruct(msg proto.Message) (*structpb.Struct, error) {
	buf := &bytes.Buffer{}
	if err := (&jsonpb.Marshaler{}).Marshal(buf, msg); err != nil {
		return nil, err
	}

	pbs := &structpb.Struct{}
	if err := jsonpb.Unmarshal(buf, pbs); err != nil {
		return nil, err
	}

	return pbs, nil
}