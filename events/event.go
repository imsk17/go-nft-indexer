package events

// Event represents a Smart Contract event in the evm blockchain.
// For more information about the event, see:
// https://consensys.net/blog/developers/guide-to-events-and-logs-in-ethereum-smart-contracts/
// It is just an interface, which contains the following methods:
// 1 - Topic() string :- which returns the event's topic (keckakk256('EventSignature')).
// 2 - String() string :- should be able to convert itself into a human readable string which can be
// logged if needed.
type Event interface {
	Topic() string
	String() string
}
