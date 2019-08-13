/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

// BallotBox is an interface for a BallotBox-ish struct.
type BallotBox interface {
	castVote(Vote)
	getResults() []Vote
	getWinner() []string
}

type BallotBoxProtoType struct {
	VotingSystem string
}
