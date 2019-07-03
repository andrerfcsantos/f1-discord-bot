package eargast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
