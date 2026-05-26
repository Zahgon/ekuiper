// Copyright 2021-2024 EMQ Technologies Co., Ltd.
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
	"github.com/lf-edge/ekuiper/contract/v2/api"
	"github.com/mmcloughlin/geohash"
)

type (
	geohashEncode         struct{}
	geohashEncodeInt      struct{}
	geohashDecode         struct{}
	geohashDecodeInt      struct{}
	geohashBoundingBox    struct{}
	geohashBoundingBoxInt struct{}
	geohashNeighbor       struct{}
	geohashNeighborInt    struct{}
	geohashNeighbors      struct{}
	geohashNeighborsInt   struct{}
	position              struct {
		Longitude float64
		Latitude  float64
	}
)

var (
	GeohashEncode         geohashEncode
	GeohashEncodeInt      geohashEncodeInt
	GeohashDecode         geohashDecode
	GeohashDecodeInt      geohashDecodeInt
	GeohashBoundingBox    geohashBoundingBox
	GeohashBoundingBoxInt geohashBoundingBoxInt
	GeohashNeighbor       geohashNeighbor
	GeohashNeighborInt    geohashNeighborInt
	GeohashNeighbors      geohashNeighbors
	GeohashNeighborsInt   geohashNeighborsInt
	g_direction           = map[string]geohash.Direction{
		"North":     geohash.North,
		"NorthEast": geohash.NorthEast,
		"East":      geohash.East,
		"SouthEast": geohash.SouthEast,
		"South":     geohash.South,
		"SouthWest": geohash.SouthWest,
		"West":      geohash.West,
		"NorthWest": geohash.NorthWest,
	}
)

func (r *geohashEncode) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashEncodeInt) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashDecode) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashDecodeInt) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashBoundingBox) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashBoundingBoxInt) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashNeighbor) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashNeighborInt) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashNeighbors) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashNeighborsInt) IsAggregate() bool { _ = "STUB: not implemented"; return false }

func (r *geohashEncode) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashEncodeInt) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashDecode) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashDecodeInt) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashBoundingBox) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashBoundingBoxInt) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashNeighbor) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashNeighborInt) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashNeighbors) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashNeighborsInt) Validate(args []any) error { _ = "STUB: not implemented"; return nil }

func (r *geohashEncode) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashEncodeInt) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashDecode) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashDecodeInt) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashBoundingBox) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashBoundingBoxInt) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashNeighbor) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashNeighborInt) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashNeighbors) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (r *geohashNeighborsInt) Exec(args []any, _ api.FunctionContext) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}
