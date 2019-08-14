/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func CastVote(stub shim.ChaincodeStubInterface, args []string) {

	// votingSystem :=           get voting system based on campaign id

	// todo a bunch of switch statements follow

}

// BallotBox is an interface for a BallotBox-ish struct.
type BallotBox interface {
	castVote(Vote)
	getBox() []Vote
	getWinner() []string
}

// BallotBoxPrototype is a prototype struct to be embedded in specific BallotBox structs implemented for different
// types of voting systems. Its methods are inserted below.
type BallotBoxPrototype struct {
	box []Vote
}

// castVoteChecks does preliminary checks to ensure that the vote is allowed to be casted by caller.
func (t BallotBoxPrototype) castVoteChecks(stub shim.ChaincodeStubInterface, voterID string, campaignID string, voterGroupID string) error {

	var err error

	// check for correct account, only personal accounts can vote
	accountTypeOK, err := CheckCertAttribute(stub, "accountType", "Personal")
	if !accountTypeOK {
		return errors.New(err.Error())
	}

	err = t._checkVoterHasVoted(stub, voterID, campaignID)
	if err != nil {
		return err
	}

	err = t._checkValidVoter(stub, voterID, voterGroupID)
	if err != nil {
		return err
	}

	return nil
}

// _checkVoterHasVoted checks that voterID has not already casted a vote in campaign campaignID.
func (t BallotBoxPrototype) _checkVoterHasVoted(stub shim.ChaincodeStubInterface, voterID string, campaignID string) error {

	campaignsListBytes, err := stub.GetState("Campaigns List")
	if err != nil {
		return err
	}
	campaignsList := CampaignsList{}
	json.Unmarshal(campaignsListBytes, &campaignsList)
	campaign, err := queryCampaignByID(campaignID, campaignsList.Campaigns) //todo put queryCampaignByID in helpers.go
	if err != nil {
		return err
	}

	isVoterIDExists := t.getVoterIDExists(voterID, campaign) //todo getVoterIDExists should be in helpers.go
	if isVoterIDExists {
		return errors.New("Voter with ID: " + voterID + " has already voted in campaign with ID: " + campaignID)
	}

	return nil
}

func (t BallotBoxPrototype) getVoterIDExists(voterID string, campaign Campaign) bool { //todo put in helpers.go
	for _, vote := range campaign.CampaignBallotBox.getBox() {
		if vote.getVoterID() == voterID {
			return true
		}
	}

	return false
}

// checkValidVoter checks that voterGroupID is valid and voter voterID is part of voter group voterGroupID
func (t BallotBoxPrototype) _checkValidVoter(stub shim.ChaincodeStubInterface, voterID string, voterGroupID string) error {

	// TODO: add check that voter group is part of campaign
	// campaignsList, err := _getCampaignsList(stub)
	// if err != nil {
	// 	return err
	// }
	// campaign, err := queryCampaignByID(campaignID, campaignsList.Campaigns)
	// if err != nil {
	// 	return nil
	// }

	voterGroupsList, err := getVoterGroupsList(stub)
	if err != nil {
		return err
	}

	// voter is valid when voter belongs to valid voter group
	voterGroup, err := queryVoterGroupsByID(voterGroupID, voterGroupsList.VoterGroups)
	if err != nil {
		return err
	}

	err = t._isVoterBelongVoterGroup(voterID, voterGroup)
	if err != nil {
		return err
	}

	return nil
}

// _isVoterBelongVoterGroup returns whether voterID belongs in voterGroup
func (t BallotBoxPrototype) _isVoterBelongVoterGroup(voterID string, voterGroup VoterGroup) error {
	for _, voter := range voterGroup.Voters {
		if voter.ID == voterID {
			return nil
		}
	}

	return errors.New("Voter with ID: " + voterID + " does not belong to voter group with ID: " + voterGroup.ID)
}

// getBox returns the ballotbox
func (t BallotBoxPrototype) getBox() []Vote {
	return t.box
}
