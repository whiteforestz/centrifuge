package centrifuge

import (
	"time"
)

// Config contains Node configuration options.
type Config struct {
	// Version of server – will be sent to client on connection establishment
	// phase in response to connect request.
	Version string
	// Name of this server node - must be unique, used as human readable
	// and meaningful node identifier.
	Name string
	// ClientPresenceUpdateInterval is an interval how often connected clients
	// must update presence info.
	ClientPresenceUpdateInterval time.Duration
	// ClientPresenceExpireInterval is an interval how long to consider
	// presence info valid after receiving presence ping.
	ClientPresenceExpireInterval time.Duration
	// ClientExpiredCloseDelay is an extra time given to client to
	// refresh its connection in the end of connection lifetime.
	ClientExpiredCloseDelay time.Duration
	// ClientExpiredSubCloseDelay is an extra time given to client to
	// refresh its expiring subscription in the end of subscription lifetime.
	ClientExpiredSubCloseDelay time.Duration
	// ClientStaleCloseDelay is a timeout after which connection will be
	// closed if still not authenticated (i.e. no valid connect command
	// received yet).
	ClientStaleCloseDelay time.Duration
	// ClientChannelPositionCheckDelay defines minimal time from previous
	// client position check in channel. If client does not pass check it will
	// be disconnected with DisconnectInsufficientState.
	ClientChannelPositionCheckDelay time.Duration
	// NodeInfoMetricsAggregateInterval sets interval for automatic metrics aggregation.
	// It's not very reasonable to have it less than one second.
	NodeInfoMetricsAggregateInterval time.Duration
	// LogLevel is a log level to use. By default nothing will be logged.
	LogLevel LogLevel
	// LogHandler is a handler func node will send logs to.
	LogHandler LogHandler
	// ClientQueueMaxSize is a maximum size of client's message queue in bytes.
	// After this queue size exceeded Centrifugo closes client's connection.
	ClientQueueMaxSize int
	// ClientChannelLimit sets upper limit of channels each client can subscribe to.
	ClientChannelLimit int
	// ClientUserConnectionLimit limits number of client connections from user with the
	// same ID. 0 - unlimited.
	ClientUserConnectionLimit int
	// ChannelMaxLength is a maximum length of channel name.
	ChannelMaxLength int
	// ChannelOptionsFunc ...
	ChannelOptionsFunc ChannelOptionsFunc
}

// Validate validates config and returns error if problems found
func (c *Config) Validate() error {
	return nil
}

const (
	// nodeInfoPublishInterval is an interval how often node must publish
	// node control message.
	nodeInfoPublishInterval = 3 * time.Second
	// nodeInfoCleanInterval is an interval in seconds, how often node must
	// clean information about other running nodes.
	nodeInfoCleanInterval = nodeInfoPublishInterval * 3
	// nodeInfoMaxDelay is an interval in seconds – how many seconds node
	// info considered actual.
	nodeInfoMaxDelay = nodeInfoPublishInterval*2 + time.Second
)

// DefaultConfig is Config initialized with default values for all fields.
var DefaultConfig = Config{
	Name: "centrifuge",

	ChannelMaxLength: 255,

	NodeInfoMetricsAggregateInterval: 60 * time.Second,

	ClientPresenceUpdateInterval:    25 * time.Second,
	ClientPresenceExpireInterval:    60 * time.Second,
	ClientExpiredCloseDelay:         25 * time.Second,
	ClientExpiredSubCloseDelay:      25 * time.Second,
	ClientStaleCloseDelay:           25 * time.Second,
	ClientChannelPositionCheckDelay: 40 * time.Second,
	ClientQueueMaxSize:              10485760, // 10MB by default
	ClientChannelLimit:              128,
}
