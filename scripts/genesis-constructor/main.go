package main

import (
"encoding/json"
"fmt"
"os"
"time"

"github.com/cosmos/cosmos-sdk/types/bech32"
)

type Account struct {
Name         string `json:"name"`
Address      string `json:"address"`
Amount       int64  `json:"amount"`
Type         string `json:"type"`
Multisig     bool   `json:"multisig,omitempty"`
CliffSeconds int64  `json:"cliff_seconds,omitempty"`
VestSeconds  int64  `json:"vest_seconds,omitempty"`
}

type Config struct {
Denom    string    `json:"denom"`
Accounts []Account `json:"accounts"`
}

func main() {
if len(os.Args) < 4 {
fmt.Fprintln(os.Stderr, "usage: genesis-constructor <genesis.json> <accounts.json> <output.json>")
os.Exit(1)
}
genesisPath, accountsPath, outPath := os.Args[1], os.Args[2], os.Args[3]

accBytes, err := os.ReadFile(accountsPath)
if err != nil {
panic(err)
}
var cfg Config
if err := json.Unmarshal(accBytes, &cfg); err != nil {
panic(err)
}

genBytes, err := os.ReadFile(genesisPath)
if err != nil {
panic(err)
}
var genesis map[string]interface{}
if err := json.Unmarshal(genBytes, &genesis); err != nil {
panic(err)
}

appState, ok := genesis["app_state"].(map[string]interface{})
if !ok {
panic("genesis has no app_state")
}

genesisTimeStr, _ := genesis["genesis_time"].(string)
genesisTime, err := time.Parse(time.RFC3339, genesisTimeStr)
if err != nil {
panic(fmt.Sprintf("cannot parse genesis_time %q: %v", genesisTimeStr, err))
}

authAccounts := []interface{}{}
balances := []interface{}{}
var total int64

for _, a := range cfg.Accounts {
if total+a.Amount > 1000000000000000 {
panic(fmt.Sprintf("total would exceed 1B: %d + %d", total, a.Amount))
}
total += a.Amount

// Validate the address (bech32 decode + re-encode check)
hrp, addrBytes, err := bech32.DecodeAndConvert(a.Address)
if err != nil {
panic(fmt.Sprintf("bad address %q: %v", a.Address, err))
}
if hrp != "norynth" {
panic(fmt.Sprintf("unexpected hrp %q for %q", hrp, a.Name))
}
reencoded, err := bech32.ConvertAndEncode(hrp, addrBytes)
if err != nil || reencoded != a.Address {
panic(fmt.Sprintf("address round-trip failed for %q", a.Name))
}

switch a.Type {
case "base":
acc := map[string]interface{}{
"@type": "/cosmos.auth.v1beta1.BaseAccount",
"address": a.Address,
"pub_key": nil,
"account_number": "0",
"sequence": "0",
}
authAccounts = append(authAccounts, acc)

		case "continuous":
			vestingAmount := []map[string]string{{"denom": cfg.Denom, "amount": fmt.Sprintf("%d", a.Amount)}}
			cliffEnd := genesisTime.Add(time.Duration(a.CliffSeconds) * time.Second)
			endTime := genesisTime.Add(time.Duration(a.CliffSeconds+a.VestSeconds) * time.Second)
			acc := map[string]interface{}{
				"@type": "/cosmos.vesting.v1beta1.ContinuousVestingAccount",
				"base_vesting_account": map[string]interface{}{
					"base_account": map[string]interface{}{
						"address":        a.Address,
						"pub_key":        nil,
						"account_number": "0",
						"sequence":       "0",
					},
					"original_vesting":  vestingAmount,
					"delegated_free":    []map[string]string{},
					"delegated_vesting": []map[string]string{},
					"end_time":          fmt.Sprintf("%d", endTime.Unix()),
				},
				"start_time": fmt.Sprintf("%d", cliffEnd.Unix()),
			}
			authAccounts = append(authAccounts, acc)

default:
panic(fmt.Sprintf("unknown account type %q for %q", a.Type, a.Name))
}

balances = append(balances, map[string]interface{}{
"address": a.Address,
"coins":   []map[string]string{{"denom": cfg.Denom, "amount": fmt.Sprintf("%d", a.Amount)}},
})
}

if total != 1000000000000000 {
panic(fmt.Sprintf("total is %d, expected 1000000000000000", total))
}

appState["auth"].(map[string]interface{})["accounts"] = authAccounts
bankState := appState["bank"].(map[string]interface{})
bankState["balances"] = balances
bankState["supply"] = []map[string]string{{"denom": cfg.Denom, "amount": fmt.Sprintf("%d", total)}}

outBytes, err := json.MarshalIndent(genesis, "", "  ")
if err != nil {
panic(err)
}
if err := os.WriteFile(outPath, outBytes, 0644); err != nil {
panic(err)
}

fmt.Printf("OK: wrote %s\n", outPath)
fmt.Printf("OK: %d accounts, %d ukrn (%.0f KRN)\n", len(cfg.Accounts), total, float64(total)/1e6)
}
