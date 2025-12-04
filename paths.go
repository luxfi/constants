// Copyright (C) 2019-2025, Lux Industries, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import "time"

// Base directory configuration - consistent across node, netrunner, cli
const (
	// BaseDirName is the default directory name for Lux data
	// Located at ~/.lux on all platforms
	BaseDirName = ".lux"
)

// Binary configuration
const (
	// NodeBinaryName is the name of the node binary executable
	NodeBinaryName = "luxd"

	// NodeInstallDir is the subdirectory for node binary under ~/.lux/bin/
	NodeInstallDir = "node"
)

// Snapshot configuration
const (
	// SnapshotPrefix is the prefix for network snapshot directories
	// Full path: ~/.lux/snapshots/net-snapshot-{name}/
	SnapshotPrefix = "net-snapshot-"

	// DefaultSnapshotName is the default snapshot name
	DefaultSnapshotName = "default"
)

// Directory structure
const (
	// BinDir is the subdirectory for binaries
	BinDir = "bin"

	// SnapshotsDir is the subdirectory for network snapshots
	SnapshotsDir = "snapshots"

	// RunsDir is the subdirectory for runtime data
	RunsDir = "runs"

	// PluginsDir is the subdirectory for VM plugins
	PluginsDir = "plugins"

	// LogDir is the subdirectory for logs
	LogDir = "logs"
)

// File permissions
const (
	DefaultPerms755    = 0o755
	WriteReadReadPerms = 0o644
	WriteReadOnlyPerms = 0o600
)

// API endpoints
const (
	// LocalAPIEndpoint is the default local node API endpoint
	LocalAPIEndpoint = "http://127.0.0.1:9630"

	// TestnetAPIEndpoint is the Lux testnet API endpoint
	TestnetAPIEndpoint = "https://api.lux-test.network"

	// MainnetAPIEndpoint is the Lux mainnet API endpoint
	MainnetAPIEndpoint = "https://api.lux.network"
)

// WebSocket endpoints
const (
	LocalWSEndpoint   = "ws://127.0.0.1:9630/ext/bc/C/ws"
	TestnetWSEndpoint = "wss://api.lux-test.network/ext/bc/C/ws"
	MainnetWSEndpoint = "wss://api.lux.network/ext/bc/C/ws"
)

// Default ports
const (
	// DefaultHTTPPort is the default HTTP API port
	DefaultHTTPPort = 9630

	// DefaultStakingPort is the default staking/P2P port
	DefaultStakingPort = 9631
)

// Timeouts
const (
	RequestTimeout    = 3 * time.Minute
	APIRequestTimeout = 30 * time.Second
)
