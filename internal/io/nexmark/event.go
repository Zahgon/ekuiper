// Copyright 2025 EMQ Technologies Co., Ltd.
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

package nexmark

import (
	"math/rand"
	"strings"

	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

var (
	states          []string
	cities          []string
	firstNames      []string
	lastNames       []string
	hotChannels     []string
	PersonIDs       []uint64
	AuctionIDs      []uint64
	categoriesCount int
	mu              syncx.RWMutex
)

func init() {
	states = strings.Split("az,ca,id,or,wa,wy", ",")
	cities = strings.Split("phoenix,los angeles,san francisco,boise,portland,bend,redmond,seattle,kent,cheyenne", ",")
	firstNames = strings.Split("peter,paul,luke,john,saul,vicky,kate,julie,sarah,deiter,walter", ",")
	lastNames = strings.Split("shultz,abrams,spencer,white,bartels,walton,smith,jones,noris", ",")
	hotChannels = strings.Split("Google,Facebook,Baidu,Apple", ",")
	PersonIDs = make([]uint64, 0)
	AuctionIDs = make([]uint64, 0)
	categoriesCount = 5
}

type Person struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	CreditCard   string `json:"creditCard"`
	City         string `json:"city"`
	State        string `json:"state"`
	Datetime     uint64 `json:"datetime"`
	Extra        string `json:"extra"`
}

func (p Person) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func genPersonID(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

func pickPersonID(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

func NewPerson(eventID int64, time uint64) Person { _ = "STUB: not implemented"; return *new(Person) }

type Auction struct {
	ID          uint64 `json:"id"`
	ItemName    string `json:"itemName"`
	Description string `json:"description"`
	InitialBid  uint64 `json:"initialBid"`
	Reserve     uint64 `json:"reserve"`
	Datetime    uint64 `json:"datetime"`
	Expires     uint64 `json:"expires"`
	Seller      uint64 `json:"seller"`
	Category    uint64 `json:"category"`
	Extra       string `json:"extra"`
}

func (a Auction) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func NewAuction(eventID int64, time uint64) Auction {
	_ = "STUB: not implemented"
	return *new(Auction)
}

func genAuctionID(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

func pickAuctionID(r *rand.Rand) uint64 { _ = "STUB: not implemented"; return 0 }

type Bid struct {
	Auction  uint64 `json:"auction"`
	Bidder   uint64 `json:"bidder"`
	Price    uint64 `json:"price"`
	Channel  string `json:"channel"`
	Url      string `json:"url"`
	Datetime uint64 `json:"datetime"`
	Extra    string `json:"extra"`
}

func (b Bid) ToMap() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func NewBid(eventID int64, time uint64) Bid { _ = "STUB: not implemented"; return *new(Bid) }

func randString(length int) string { _ = "STUB: not implemented"; return "" }

func randNumber(length int) string { _ = "STUB: not implemented"; return "" }

func randUrl() string { _ = "STUB: not implemented"; return "" }
