/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getAccountsList(stub shim.ChaincodeStubInterface) (AccountsList, error) {
	accountsListBytes, err := stub.GetState("Accounts List")
	if err != nil {
		return AccountsList{}, err
	}
	accountsList := AccountsList{}
	json.Unmarshal(accountsListBytes, &accountsList)

	return accountsList, nil
}

func getCampaignsList(stub shim.ChaincodeStubInterface) (CampaignsList, error) {
	campaignsListBytes, err := stub.GetState("Campaigns List")
	if err != nil {
		return CampaignsList{}, err
	}
	campaignsList := CampaignsList{}
	json.Unmarshal(campaignsListBytes, &campaignsList)

	return campaignsList, nil
}

func getVoterGroupsList(stub shim.ChaincodeStubInterface) (VoterGroupsList, error) {
	voterGroupsListBytes, err := stub.GetState("Voter Groups List")
	if err != nil {
		return VoterGroupsList{}, err
	}
	voterGroupsList := VoterGroupsList{}
	json.Unmarshal(voterGroupsListBytes, &voterGroupsList)

	return voterGroupsList, nil
}

//----------------------------------------------Query--------------------------------------------------

// queryAccountByID is a helper that queries the Accounts array for id and returns the account with id.
func queryAccountByID(id string, accounts []Account) (Account, error) {

	for _, v := range accounts {
		if v.ID == id {
			return v, nil
		}
	}

	return Account{}, errors.New("Account with id: " + id + " does not exist.")
}

// queryByAccountExistsById queries the Accounts array for id and returns whether it exists.
func queryAccountExistsByID(id string, accounts []Account) bool {

	for _, v := range accounts {
		if v.ID == id {
			return true
		}
	}

	return false
}

//----------------------------------------------Identity---------------------------------------------------
// CheckCertAttribute checks whether parameter matches with the caller's certificates attributes.
// Returns true if attribute matches.
func CheckCertAttribute(stub shim.ChaincodeStubInterface, attribute string, parameter string) (bool, error) {
	val, ok, err := cid.GetAttributeValue(stub, attribute)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("The client identity does not possess attribute: " + attribute)
	}
	if val != parameter {
		return false, errors.New("User certificate's attribute: " + attribute +
			" does not match " + parameter)
	}
	return true, nil
}

// GetCertAttribute returns the value of attribute in the caller's certificate.
func GetCertAttribute(stub shim.ChaincodeStubInterface, attribute string) (string, error) {
	val, ok, err := cid.GetAttributeValue(stub, attribute)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("The client identity does not possess attribute: " + attribute)
	}

	return val, nil
}

//------------------------------------------MISC---------------------------------------------------------
func removeAccount(s []Account, i int) []Account {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeCampaign(s []Campaign, i int) []Campaign {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeVoterGroup(s []VoterGroup, i int) []VoterGroup {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeVoter(s []Voter, i int) []Voter {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
