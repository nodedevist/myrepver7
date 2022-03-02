package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

type field struct {
	Indexed      bool   `json:"indexed"`
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type abiItem struct {
	Inputs          []field `json:"inputs"`
	Name            string  `json:"name"`
	Outputs         []field `json:"outputs"`
	StateMutability string  `json:"stateMutability"`
	Type            string  `json:"type"`
}

func (a *abiItem) Signature() string {
	var inputs []string
	for _, i := range a.Inputs {
		inputs = append(inputs, i.Type)
	}
	return fmt.Sprintf("%s(%s)", a.Name, strings.Join(inputs, ","))
}

func fatal(msg string) {
	log.Error(msg)
	os.Exit(1)
}

type TopicDefinition struct {
	Name      string
	Signature string
	Topic     string
}

func main() {

	if len(os.Args) < 2 {
		log.Info(os.Args)
		fatal("abi path argument is required")
	}
	abiFileName := os.Args[1]

	b, err := os.ReadFile(abiFileName)
	if err != nil {
		fatal(fmt.Sprintf("error getting abi file: %v", err.Error()))
	}

	var abi []abiItem
	err = json.Unmarshal(b, &abi)
	if err != nil {
		fatal(err.Error())
	}

	var defs []TopicDefinition
	for _, item := range abi {
		if item.Type == "event" {
			defs = append(defs, TopicDefinition{
				Name:      item.Name,
				Signature: item.Signature(),
				Topic:     crypto.Keccak256Hash([]byte(item.Signature())).Hex(),
			})
		}
	}

	dir := filepath.Dir(abiFileName)
	dirTokens := strings.Split(dir, "/")
	dirName := dirTokens[len(dirTokens)-1]
	topicFile := fmt.Sprintf("%s/%s", dir, "topics.go")

	f, err := os.Create(topicFile)
	if err != nil {
		fatal(fmt.Sprintf("error creating topic file: %v", err.Error()))
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("package %s", dirName))
	lines = append(lines, "")
	for _, def := range defs {
		lines = append(lines, fmt.Sprintf("var %sSignature = \"%s\"", def.Name, def.Signature))
		lines = append(lines, fmt.Sprintf("var %sTopic = \"%s\"", def.Name, def.Topic))
	}

	_, err = f.WriteString(strings.Join(lines, "\n"))
	if err != nil {
		fatal(err.Error())
	}
}
