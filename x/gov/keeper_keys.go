package gov

import (
	"bytes"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO remove some of these prefixes once have working multistore

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter = []byte("/")

	KeyNextProposalID           = []byte("newProposalID")
	PrefixActiveProposalQueue   = []byte("activeProposalQueue")
	PrefixInactiveProposalQueue = []byte("inactiveProposalQueue")
)

// Key for getting a specific proposal from the store
func KeyProposal(proposalID int64) []byte {
	return []byte(fmt.Sprintf("proposals:%d", proposalID))
}

// Key for getting a specific deposit from the store
func KeyDeposit(proposalID int64, depositerAddr sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("deposits:%d:%d", proposalID, depositerAddr))
}

// Key for getting a specific vote from the store
func KeyVote(proposalID int64, voterAddr sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("votes:%d:%d", proposalID, voterAddr))
}

// Key for getting all deposits on a proposal from the store
func KeyDepositsSubspace(proposalID int64) []byte {
	return []byte(fmt.Sprintf("deposits:%d:", proposalID))
}

// Key for getting all votes on a proposal from the store
func KeyVotesSubspace(proposalID int64) []byte {
	return []byte(fmt.Sprintf("votes:%d:", proposalID))
}

// Returns the key for a proposalID in the activeProposalQueue
func ActiveProposalQueueTimePrefix(endTime time.Time) []byte {
	return bytes.Join([][]byte{
		PrefixActiveProposalQueue,
		sdk.FormatTimeBytes(endTime),
	}, KeyDelimiter)
}

// Returns the key for a proposalID in the activeProposalQueue
func ActiveProposalQueueProposalKey(endTime time.Time, proposalID int64) []byte {
	return bytes.Join([][]byte{
		PrefixActiveProposalQueue,
		sdk.FormatTimeBytes(endTime),
		sdk.Uint64ToBigEndian(uint64(proposalID)),
	}, KeyDelimiter)
}

// Returns the key for a proposalID in the activeProposalQueue
func InactiveProposalQueueTimePrefix(endTime time.Time) []byte {
	return bytes.Join([][]byte{
		PrefixInactiveProposalQueue,
		sdk.FormatTimeBytes(endTime),
	}, KeyDelimiter)
}

// Returns the key for a proposalID in the activeProposalQueue
func InactiveProposalQueueProposalKey(endTime time.Time, proposalID int64) []byte {
	return bytes.Join([][]byte{
		PrefixInactiveProposalQueue,
		sdk.FormatTimeBytes(endTime),
		sdk.Uint64ToBigEndian(uint64(proposalID)),
	}, KeyDelimiter)
}
