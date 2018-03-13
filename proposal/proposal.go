/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample Manage Proposal for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Manage Proposals
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the ManageProposal structure
type ManageProposal struct {
}

// Define the Proposal structure, with 3 properties.  Structure tags are used by encoding/json library
type proposal struct{

								// Attributes of a Form 
	proposal_id string `json:"proposal_id"`	
	region string `json:"region"`
	country string `json:"country"`
	//proposal_type string `json:"proposal_type"`
	//initiated_on string `json:"initiated_on"`
	//euc_reviewed_on string `json:"euc_reviewed_on"`
	//shared_with_sd_and_a_on string `json:"shared_with_sd_and_a_on"`
	//approval_on string `json:"approval_on"`
	//shared_with_procurement_team_on string `json:"shared_with_procurement_team_on"`
	//number_of_tasks_covered string `json:"number_of_tasks_covered"`
	//device_qty string `json:"device_qty"`
	//accessary_periperal_qty string `json:"accessary_periperal_qty"`
	//total_qty string `json:"total_qty"`
	//status string `json:"status"`
	
	
	
	
	
}







/*
 * The Init method is called when the Manage Proposal "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *ManageProposal) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}




/*
 * The Invoke method is called as a result of an application request to run the Manage Proposal "fabcar"
 * The calling application program has also specified the particular Manage Proposal function to be called, with arguments
 */
func (s *ManageProposal) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Manage Proposal function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryProposal" {
		return s.queryProposal(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createProposal" {
		return s.createProposal(APIstub, args)
	}
/*	else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(APIstub, args)
	}
	*/

	return shim.Error("Invalid Manage Proposal function name.")
}




func (s *ManageProposal) queryProposal(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	proposalAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(proposalAsBytes)
}







func (s *ManageProposal) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	proposals := []proposal{
		Proposal{proposal_id: "BLRSJPPR001", region: "Bangalore", country: "India"},
		Proposal{proposal_id: "LONBKDPR001", region: "London", country: "UK"},		
		Proposal{proposal_id: "BLRECEPR001", region: "Bangaluru", country: "India"},
	}

	i := 0
	for i < len(proposals) {
		fmt.Println("i is ", i)
		proposalAsBytes, _ := json.Marshal(proposals[i])
		APIstub.PutState(proposals[i].proposal_id, proposalAsBytes)
		fmt.Println("Added", proposals[i])
		i = i + 1
	}

	return shim.Success(nil)
}





func (s *ManageProposal) createProposal(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var proposal = Proposal{proposal_id: args[1], region: args[2], country: args[3]}

	proposalAsBytes, _ := json.Marshal(proposal)
	APIstub.PutState(args[1], proposalAsBytes)

	return shim.Success(nil)
}

