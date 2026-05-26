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
	"time"

	"github.com/jinzhu/now"
)

//var (
//	formats = map[string]string{
//		"MMMM": "January",
//		"MMM":  "Jan",
//		"MM":   "01",
//		"M":    "1",
//		"YYYY": "2006",
//		"yyyy": "2006",
//		"YY":   "06",
//		"yy":   "06",
//		"G":    "AD",
//		"EEEE": "Monday",
//		"EEE":  "Mon",
//		"dd":   "02",
//		"d":    "2",
//		"HH":   "15",
//		"hh":   "03",
//		"h":    "3",
//		"mm":   "04",
//		"m":    "4",
//		"ss":   "05",
//		"s":    "5",
//		"a":    "PM",
//		"S":    ".0",
//		"SS":    ".00",
//		"SSS":   ".000",
//		"SSSN":  ".0000",
//		"SSSNN": ".00000",
//		"SSSNNN": ".000000",
//		"SSSNNNN": ".0000000",
//		"SSSNNNNN":".00000000",
//		"SSSNNNNNN":".000000000",
//		"z":    "MST",
//		"Z":    "-0700",
//		"X":    "-07",
//		"XX":    "-0700",
//		"XXX":  "-07:00",
//	}
//)

const (
	JSISO   = "2006-01-02T15:04:05.000Z07:00"
	ISO8601 = "2006-01-02T15:04:05"
)

func init() {
	now.TimeFormats = append(now.TimeFormats, JSISO, ISO8601)
}

func GetConfiguredTimeZone() *time.Location { _ = "STUB: not implemented"; return nil }

var localTimeZone = time.Local

func SetTimeZone(name string) error { _ = "STUB: not implemented"; return nil }

func TimeToUnixMilli(time time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func InterfaceToUnixMilli(i interface{}, format string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func InterfaceToTime(i interface{}, format string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func TimeFromUnixMilli(t int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func ParseTime(t string, f string) (_ time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ParseTimeByFormats(t string, formats []string) (_ time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func FormatTime(time time.Time, f string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//func convertFormat(f string) string {
//	re := regexp.MustCompile(`(?m)(M{4})|(M{3})|(M{2})|(M{1})|(Y{4})|(Y{2})|(y{4})|(y{2})|(G{1})|(E{4})|(E{3})|(d{2})|(d{1})|(H{2})|(h{2})|(h{1})|(m{2})|(m{1})|(s{2})|(s{1})|(a{1})|(S{3}N{6})|(S{3}N{5})|(S{3}N{4})|(S{3}N{3})|(S{3}N{2})|(S{3}N{1})|(S{3})|(S{2})|(S{1})|(z{1})|(Z{1})|(X{3})|(X{2})|(X{1})`)
//	for _, match := range re.FindAllString(f, -1) {
//		for key, val := range formats {
//			if match == key {
//				f = strings.Replace(f, match, val, -1)
//			}
//		}
//	}
//	return f
//}

func convertFormat(f string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// YYYY

// YY

// era

// M MM MMM MMMM month of year

// M

// MM

// MMM

// MMMM

// d dd day of month

// d

// dd

// M MM MMM MMMM month of year

// EEE

// EEEE

// HH

// HH

// h hh

// h

// hh

// a

// m mm minute of hour

// m

// mm

// s ss

// s

// ss

// S SS SSS....

// z

// Z

// X XX XXX

// X

// XX

// XXX

// ' (text delimiter)  or '' (real quote)

// real quote

// InterfaceToDuration converts an interface to a time.Duration.
func InterfaceToDuration(i interface{}) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func padFractionalSeconds(dateStr string) string { _ = "STUB: not implemented"; return "" }
