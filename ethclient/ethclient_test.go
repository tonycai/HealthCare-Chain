// Copyright 2016 The healthcare-chain Authors
// This file is part of the healthcare-chain library.
//
// The healthcare-chain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The healthcare-chain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the healthcare-chain library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/tonycai/HealthCare-Chain"

// Verify that Client implements the ethcc interfaces.
var (
	_ = ethcc.ChainReader(&Client{})
	_ = ethcc.TransactionReader(&Client{})
	_ = ethcc.ChainStateReader(&Client{})
	_ = ethcc.ChainSyncReader(&Client{})
	_ = ethcc.ContractCaller(&Client{})
	_ = ethcc.GasEstimator(&Client{})
	_ = ethcc.GasPricer(&Client{})
	_ = ethcc.LogFilterer(&Client{})
	_ = ethcc.PendingStateReader(&Client{})
	// _ = ethcc.PendingStateEventer(&Client{})
	_ = ethcc.PendingContractCaller(&Client{})
)
