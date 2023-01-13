// Copyright 2023 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// ExemptionType ExemptionType is used to indicate if the live balance for an account subject to a
// BalanceExemption could increase above, decrease below, or equal the computed balance. *
// greater_or_equal: The live balance may increase above or equal the computed balance. This
// typically   occurs with staking rewards that accrue on each block. * less_or_equal: The live
// balance may decrease below or equal the computed balance. This typically   occurs as balance
// moves from locked to spendable on a vesting account. * dynamic: The live balance may increase
// above, decrease below, or equal the computed balance. This   typically occurs with tokens that
// have a dynamic supply.
type ExemptionType string

// List of ExemptionType
const (
	BalanceGreaterOrEqual ExemptionType = "greater_or_equal"
	BalanceLessOrEqual    ExemptionType = "less_or_equal"
	BalanceDynamic        ExemptionType = "dynamic"
)
