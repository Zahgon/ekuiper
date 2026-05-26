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

package cast

import (
	"sync"

	"github.com/mitchellh/mapstructure"
)

type Strictness int8

const (
	STRICT Strictness = iota
	CONVERT_SAMEKIND
	CONVERT_ALL
)

type ArrayNilConvert int8

const (
	IGNORE_NIL ArrayNilConvert = iota
	FORCE_CONVERT
)

/*********** Type Cast Utilities *****/

// TODO datetime type
func ToStringAlways(input interface{}) string { _ = "STUB: not implemented"; return "" }

func ToString(input interface{}, sn Strictness) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ToInt(input interface{}, sn Strictness) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt8(input interface{}, sn Strictness) (int8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToInt16(input interface{}, sn Strictness) (int16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToInt32(input interface{}, sn Strictness) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToInt64(input interface{}, sn Strictness) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToFloat64(input interface{}, sn Strictness) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToFloat32(input interface{}, sn Strictness) (float32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToUint64(i interface{}, sn Strictness) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToUint8(i interface{}, sn Strictness) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func ToUint16(i interface{}, sn Strictness) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToUint32(i interface{}, sn Strictness) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ToBool(input interface{}, sn Strictness) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isZeroEpsilon64(f float64) bool { _ = "STUB: not implemented"; return false }

func isZeroEpsilon32(f float32) bool { _ = "STUB: not implemented"; return false }

func ToBytes(input interface{}, sn Strictness) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToByteA converts to eKuiper internal byte array
func ToByteA(input interface{}, sn Strictness) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToStringMap(input interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// case string:
//	err := jsonStringToObject(v, &m)
//	return m, err

func ToTypedSlice(input interface{}, conv func(interface{}, Strictness) (interface{}, error), eleType string, sn Strictness) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToInt64Slice(input interface{}, sn Strictness) ([]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToUint32Slice(input interface{}, sn Strictness) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToUint64Slice(input interface{}, sn Strictness) ([]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToFloat64Slice(input interface{}, sn Strictness, anc ArrayNilConvert) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToFloat32Slice(input interface{}, sn Strictness) ([]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToBoolSlice(input interface{}, sn Strictness) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToStringSlice(input interface{}, sn Strictness) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToBytesSlice(input interface{}, sn Strictness) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//MapToStruct
/*
*   Convert a map into a struct. The output parameter must be a pointer to a struct
*   The struct can have the json metadata
 */
func MapToStruct(input, output interface{}) error { _ = "STUB: not implemented"; return nil }

func ToTimeDurationHookFunc() mapstructure.DecodeHookFunc {
	_ = "STUB: not implemented"
	return *new(mapstructure.DecodeHookFunc)
}

// MapToStructStrict
/*
*   Convert a map into a struct. The output parameter must be a pointer to a struct
*   If the input have key/value pair output do not defined, will report error
 */
func MapToStructStrict(input, output interface{}) error { _ = "STUB: not implemented"; return nil }

func ConvertMap(s map[interface{}]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func ConvertArray(s []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

func SyncMapToMap(sm *sync.Map) map[string]interface{} { _ = "STUB: not implemented"; return nil }

func MapToSyncMap(m map[string]interface{}) *sync.Map { _ = "STUB: not implemented"; return nil }

func isIntegral64(val float64) bool { _ = "STUB: not implemented"; return false }

func isIntegral32(val float32) bool { _ = "STUB: not implemented"; return false }

func ConvertToInterfaceArr(orig map[string]interface{}) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func ConvertSlice(v interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// ToType cast value into newType type
// newType support bigint, float, string, boolean, datetime, bytea
func ToType(value interface{}, newType interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
