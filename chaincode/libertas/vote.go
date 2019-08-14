/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"time"
)

type Vote interface {
	getVoterID() string
	getCampaignID() string
	getVoterGroupID() string
	getCreatedAt() time.Time
	getUpdatedAt() time.Time

	setVoterID(voterID string)
	setCampaignID(campaignID string)
	setVoterGroupID(voterGroupID string)
	setUpdatedAt(updatedTime time.Time)
}

type VotePrototype struct {
	VoterID      string
	CampaignID   string
	VoterGroupID string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (t VotePrototype) getVoterID() string {
	return t.VoterID
}

func (t VotePrototype) getCampaignID() string {
	return t.CampaignID
}

func (t VotePrototype) getVoterGroupID() string {
	return t.VoterGroupID
}

func (t VotePrototype) getCreatedAt() time.Time {
	return t.CreatedAt
}

func (t VotePrototype) getUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *VotePrototype) setVoterID(voterID string) {
	t.VoterID = voterID
}

func (t *VotePrototype) setCampaignID(campaignID string) {
	t.CampaignID = campaignID
}

func (t *VotePrototype) setVoterGroupID(voterGroupID string) {
	t.VoterGroupID = voterGroupID
}

func (t *VotePrototype) setUpdatedAt(updatedTime time.Time) {
	t.UpdatedAt = updatedTime
}

//---------------------------------------Vote Implementations------------------------------------------
type FPTP struct {
	VotePrototype
	CandidateName string
}
