package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dawsonalex/rep"
	"github.com/dawsonalex/rep/parse"
	"os"
)

func main() {
	tracker := rep.Tracker{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name set: ")
	setName, err := reader.ReadString('\n')
	checkFatalErr("error reading set name", err)

	err = tracker.BeginSession(setName)
	checkFatalErr("error starting session", err)

	for {
		nextLine, err := reader.ReadString('\n')
		checkFatalErr("error reading set", err)

		if nextLine == "\n" {
			break
		}

		set, err := parse.Set(nextLine)
		if err != nil {
			fmt.Printf("err: %v", err)
			continue
		}

		tracker.LogSet(set.RepCount, set.Weight, set.Rpe, set.Note)
	}
	fmt.Println("finished loop")

	session := tracker.EndSession()
	printSession(session)
}

func printSession(session *rep.Session) {
	sessionJson, err := json.MarshalIndent(session, "", " ")
	checkFatalErr("error marshalling session", err)
	fmt.Printf("got session: %s", sessionJson)

	fmt.Printf("%+v", session.Sets)
}

// checkFatalErr checks if the error is nil and
func checkFatalErr(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v", msg, err)
		os.Exit(1)
	}
}
