// Copyright 2023 Redpanda Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"

	"github.com/redpanda-data/redpanda/src/go/transform-sdk"
)

func main() {
	redpanda.OnRecordWritten(identityTransform)
}

func identityTransform(e redpanda.WriteEvent) ([]redpanda.Record, error) {
	key := string(e.Record().Key)
	switch key {
	case "poison":
		return nil, errors.New("☠︎")
	case "cheese":
		return []redpanda.Record{
			{Key: []byte("yum!"), Value: []byte("cheese eaten")},
		}, nil
	}
	println("unknown record:", key)
	// omit nothing
	return nil, nil
}
