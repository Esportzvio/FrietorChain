package polybftsecrets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/esportzvio/frietorchain/command"

	"github.com/esportzvio/frietorchain/command/helper"
	"github.com/esportzvio/frietorchain/types"
)

type Results []command.CommandResult

func (r Results) GetOutput() string {
	var buffer bytes.Buffer

	for _, result := range r {
		buffer.WriteString(result.GetOutput())
	}

	return buffer.String()
}

type SecretsInitResult struct {
	Address       types.Address `json:"address"`
	BLSPubkey     string        `json:"bls_pubkey"`
	NodeID        string        `json:"node_id"`
	PrivateKey    string        `json:"private_key"`
	BLSPrivateKey string        `json:"bls_private_key"`
	Insecure      bool          `json:"insecure"`
	Generated     string        `json:"generated"`
}

func secretsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists
		return true, nil
	} else if os.IsNotExist(err) {
		// File does not exist
		return false, nil
	} else {
		// An error occurred (other than "not exist")
		return false, err
	}
}

func (r *SecretsInitResult) GetOutput() string {
	var buffer bytes.Buffer

	vals := make([]string, 0, 3)

	vals = append(
		vals,
		fmt.Sprintf("Public key (address)|%s", r.Address.String()),
	)

	if r.PrivateKey != "" {
		vals = append(
			vals,
			fmt.Sprintf("Private key|%s", r.PrivateKey),
		)
	}

	if r.BLSPrivateKey != "" {
		vals = append(
			vals,
			fmt.Sprintf("BLS Private key|%s", r.BLSPrivateKey),
		)
	}

	if r.BLSPubkey != "" {
		vals = append(
			vals,
			fmt.Sprintf("BLS Public key|%s", r.BLSPubkey),
		)
	}

	vals = append(vals, fmt.Sprintf("Node ID|%s", r.NodeID))

	if r.Insecure {
		buffer.WriteString("\n[WARNING: INSECURE LOCAL SECRETS - SHOULD NOT BE RUN IN PRODUCTION]\n")
	}

	if r.Generated != "" {
		buffer.WriteString("\n[SECRETS GENERATED]\n")
		buffer.WriteString(r.Generated)
		buffer.WriteString("\n")
	}

	buffer.WriteString("\n[SECRETS INIT]\n")
	buffer.WriteString(helper.FormatKV(vals))
	buffer.WriteString("\n")

	filePath := "secrets.json"

	exists, err := secretsFileExists(filePath)

	if err != nil {
		fmt.Println(err)
	}

	data := map[string]string{
		"address":         r.Address.String(),
		"bls_pubkey":      r.BLSPubkey,
		"node_id":         r.NodeID,
		"private_key":     r.PrivateKey,
		"bls_private_key": r.BLSPrivateKey,
		"insecure":        fmt.Sprintf("%t", r.Insecure),
		"generated":       r.Generated,
	}

	// convert data map to json using json.Marshal()
	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	if exists {
		// write  data to file as json
		os.WriteFile(filePath, jsonData, 0644)
	} else {
		// create file and write buffer data to file as json
		os.Create(filePath)
		os.WriteFile(filePath, buffer.Bytes(), 0644)
	}

	return buffer.String()
}
