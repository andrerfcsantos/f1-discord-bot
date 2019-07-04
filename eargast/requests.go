package eargast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// RequestNextRace uses the eargast api to request information the next race
func RequestNextRace() (MRReply, error) {

	// Make the request
	reply, err := Client.Get("https://ergast.com/api/f1/current/next.json")
	if err != nil {
		return MRReply{}, fmt.Errorf("GET Request: %v", err)
	}

	// Read the reply bytes
	replyBytes, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return MRReply{}, fmt.Errorf("reading reply bytes: %v", err)
	}

	// Unmarshal the bytes into a MRReply
	var res MRReply
	err = json.Unmarshal(replyBytes, &res)
	if err != nil {
		return MRReply{}, fmt.Errorf("unmarshaling reply bytes into json: %v", err)
	}

	return res, nil

}

// RequestLastRace requests information about the last race
func RequestLastRace() (MRReply, error) {

	// Make the request
	reply, err := Client.Get("https://ergast.com/api/f1/current/last/results.json")
	if err != nil {
		return MRReply{}, fmt.Errorf("GET Request: %v", err)
	}

	// Read the reply bytes
	replyBytes, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return MRReply{}, fmt.Errorf("reading reply bytes: %v", err)
	}

	// Unmarshal the bytes into a MRReply
	var res MRReply
	err = json.Unmarshal(replyBytes, &res)
	if err != nil {
		return MRReply{}, fmt.Errorf("unmarshaling reply bytes into json: %v", err)
	}

	return res, nil
}

// RequestCircuitResults requests information about results on a given circuit in the last years
func RequestCircuitResults(circuitID string) (MRReply, error) {
	circuitID = strings.ToLower(circuitID)
	// Make the request
	reply, err := Client.Get(fmt.Sprintf("http://ergast.com/api/f1/circuits/%s/results/1.json?limit=1000", circuitID))
	if err != nil {
		return MRReply{}, fmt.Errorf("GET Request: %v", err)
	}

	// Read the reply bytes
	replyBytes, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return MRReply{}, fmt.Errorf("reading reply bytes: %v", err)
	}

	// Unmarshal the bytes into a MRReply
	var res MRReply
	err = json.Unmarshal(replyBytes, &res)
	if err != nil {
		return MRReply{}, fmt.Errorf("unmarshaling reply bytes into json: %v", err)
	}

	return res, nil
}

// Circuits requests information about all the circuits
func Circuits() (MRReply, error) {
	// Make the request
	reply, err := Client.Get("https://ergast.com/api/f1/circuits.json?limit=1000")
	if err != nil {
		return MRReply{}, fmt.Errorf("GET Request: %v", err)
	}

	// Read the reply bytes
	replyBytes, err := ioutil.ReadAll(reply.Body)
	if err != nil {
		return MRReply{}, fmt.Errorf("reading reply bytes: %v", err)
	}

	// Unmarshal the bytes into a MRReply
	var res MRReply
	err = json.Unmarshal(replyBytes, &res)
	if err != nil {
		return MRReply{}, fmt.Errorf("unmarshaling reply bytes into json: %v", err)
	}

	return res, nil
}
