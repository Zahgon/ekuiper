// Copyright 2022-2024 EMQ Technologies Co., Ltd.
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

package graph

type (
	IoInputType      uint8
	IoRowType        uint8
	IoCollectionType uint8
)

const (
	IOINPUT_TYPE_SAME       IoInputType = iota
	IOINPUT_TYPE_ROW                    // 0b01
	IOINPUT_TYPE_COLLECTION             // 0b10
	IOINPUT_TYPE_ANY                    // 0b11
)

var inputTypes = map[IoInputType]string{
	IOINPUT_TYPE_ROW:        "row",
	IOINPUT_TYPE_COLLECTION: "collection",
	IOINPUT_TYPE_ANY:        "any",
	IOINPUT_TYPE_SAME:       "same",
}

const (
	IOROW_TYPE_SAME   IoRowType = iota
	IOROW_TYPE_SINGLE           // 0b01
	IOROW_TYPE_MERGED           // 0b10
	IOROW_TYPE_ANY              // 0b11
)

var rowTypes = map[IoRowType]string{
	IOROW_TYPE_SINGLE: "single emitter row",
	IOROW_TYPE_MERGED: "merged row",
	IOROW_TYPE_ANY:    "any",
	IOROW_TYPE_SAME:   "same",
}

const (
	IOCOLLECTION_TYPE_SAME IoCollectionType = iota
	IOCOLLECTION_TYPE_SINGLE
	IOCOLLECTION_TYPE_GROUPED
	IOCOLLECTION_TYPE_ANY
)

var collectionsTypes = map[IoCollectionType]string{
	IOCOLLECTION_TYPE_SINGLE:  "non-grouped collection",
	IOCOLLECTION_TYPE_GROUPED: "grouped collection",
	IOCOLLECTION_TYPE_ANY:     "any",
	IOCOLLECTION_TYPE_SAME:    "same",
}

// IOType is the type of input/output
// all fields are default to any
type IOType struct {
	Type           IoInputType      `json:"type"`
	RowType        IoRowType        `json:"rowType"`
	CollectionType IoCollectionType `json:"collectionType"`
	AllowMulti     bool             `json:"allowMulti"`
}

// NewIOType creates a new IOType
func NewIOType() *IOType { _ = "STUB: not implemented"; return nil }

func Fit(value, condition *IOType) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func MapOut(previous, origin *IOType) (result *IOType) { _ = "STUB: not implemented"; return nil }

// OpIO The io constraints for a node
var OpIO = map[string][]*IOType{
	"aggfunc": {
		{Type: IOINPUT_TYPE_COLLECTION, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_COLLECTION, CollectionType: IOCOLLECTION_TYPE_GROUPED},
	},
	"filter": {
		{Type: IOINPUT_TYPE_ANY, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
	"function": {
		{Type: IOINPUT_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_SINGLE, RowType: IOROW_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
	"groupby": {
		{Type: IOINPUT_TYPE_COLLECTION, CollectionType: IOCOLLECTION_TYPE_SINGLE, RowType: IOROW_TYPE_ANY},
		{Type: IOINPUT_TYPE_COLLECTION, CollectionType: IOCOLLECTION_TYPE_GROUPED},
	},
	"join": {
		{Type: IOINPUT_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_SINGLE, RowType: IOROW_TYPE_SINGLE, AllowMulti: true},
		{Type: IOINPUT_TYPE_COLLECTION, CollectionType: IOCOLLECTION_TYPE_SINGLE, RowType: IOROW_TYPE_MERGED},
	},
	"orderby": {
		{Type: IOINPUT_TYPE_COLLECTION, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
	"pick": {
		{Type: IOINPUT_TYPE_ANY, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
	"watermark": {
		{Type: IOINPUT_TYPE_ROW, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY, AllowMulti: true},
		{Type: IOINPUT_TYPE_SAME},
	},
	"window": {
		{Type: IOINPUT_TYPE_ROW, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY, AllowMulti: true},
		{Type: IOINPUT_TYPE_COLLECTION, CollectionType: IOCOLLECTION_TYPE_SINGLE, RowType: IOROW_TYPE_SINGLE},
	},
	"switch": {
		{Type: IOINPUT_TYPE_ANY, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
	"script": {
		{Type: IOINPUT_TYPE_ANY, RowType: IOROW_TYPE_ANY, CollectionType: IOCOLLECTION_TYPE_ANY},
		{Type: IOINPUT_TYPE_SAME},
	},
}
