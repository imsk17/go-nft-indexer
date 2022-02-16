package events

// Event represents a Smart Contract event in the evm blockchain.
// For more information about the event, see:
// https://consensys.net/blog/developers/guide-to-events-and-logs-in-ethereum-smart-contracts/
type Event interface {
	Topic() string
}
