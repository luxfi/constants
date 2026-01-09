// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import "time"

// CLI and operational constants (extends paths.go)
const (
	// Devnet endpoint
	DevnetAPIEndpoint = "https://api.lux-dev.network"
	DevnetWSEndpoint  = "wss://api.lux-dev.network/ext/bc/C/ws"

	// Local network configuration
	LocalNetworkID            = CustomID // 1337 - same as CustomID
	NetrunnerLocalNetworkID   = CustomID // 1337 - same as CustomID for netrunner
	AvalancheLocalNetworkID   = 31337    // Avalanche's local network ID
	LocalNetworkNumNodes      = 5

	// Staking constants
	MinStakeDuration     = 24 * 14 * time.Hour  // 2 weeks
	MaxStakeDuration     = 24 * 365 * time.Hour // 1 year
	MinStakeWeight       = uint64(1)
	StakingStartLeadTime = 1 * time.Minute
	TimeParseLayout      = "2006-01-02 15:04:05"

	// Version management
	LuxCompatibilityURL = "https://raw.githubusercontent.com/luxfi/cli/main/lux-compatibility.json"
	DefaultLuxdVersion  = "v1.21.0"
)
