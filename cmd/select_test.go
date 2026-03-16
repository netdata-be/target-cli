package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestAddExportOrUnset(t *testing.T) {
	exportCommands := []string{}
	addExportOrUnset(&exportCommands, "TEST_ENV", "value")
	if len(exportCommands) != 1 || exportCommands[0] != "export TEST_ENV=value" {
		t.Fatalf("expected export TEST_ENV=value, got %v", exportCommands)
	}

	exportCommands = []string{}
	addExportOrUnset(&exportCommands, "TEST_ENV", "")
	if len(exportCommands) != 1 || exportCommands[0] != "unset TEST_ENV" {
		t.Fatalf("expected unset TEST_ENV, got %v", exportCommands)
	}
}

func TestAddExportBoolOrUnset(t *testing.T) {
	exportCommands := []string{}
	addExportBoolOrUnset(&exportCommands, "TEST_BOOL", true)
	if len(exportCommands) != 1 || exportCommands[0] != "export TEST_BOOL=true" {
		t.Fatalf("expected export TEST_BOOL=true, got %v", exportCommands)
	}

	exportCommands = []string{}
	addExportBoolOrUnset(&exportCommands, "TEST_BOOL", false)
	if len(exportCommands) != 1 || exportCommands[0] != "unset TEST_BOOL" {
		t.Fatalf("expected unset TEST_BOOL, got %v", exportCommands)
	}
}

func TestSelectCmd_GeneratesUnsetOutput(t *testing.T) {
	c = &Config{
		Vault: map[string]*Vault{
			"dri": {
				Endpoint: "https://vault.example",
				NoColor:  false,
			},
		},
		Nomad: map[string]*Nomad{
			"dri": {
				NomadEndpoint: "https://nomad.example",
			},
		},
		Consul: map[string]*Consul{
			"dri": {
				ConsulEndpoint: "https://consul.example",
			},
		},
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	selectCmd.Run(selectCmd, []string{"dri"})
	w.Close()
	os.Stdout = oldStdout

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "export VAULT_ADDR=https://vault.example") {
		t.Fatalf("output missing VAULT_ADDR, got: %s", output)
	}
	if !strings.Contains(output, "unset VAULT_TOKEN") {
		t.Fatalf("output missing unset VAULT_TOKEN, got: %s", output)
	}
	if !strings.Contains(output, "export NOMAD_ADDR=https://nomad.example") {
		t.Fatalf("output missing NOMAD_ADDR, got: %s", output)
	}
	if !strings.Contains(output, "export CONSUL_HTTP_ADDR=https://consul.example") {
		t.Fatalf("output missing CONSUL_HTTP_ADDR, got: %s", output)
	}
}
