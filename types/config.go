package types

import (
	"errors"
	"fmt"
)

var networkNames = map[string]string{
	"1":  "DiscoveryNet",
	"2":  "IdentityForge",
	"3":  "QuestChain",
	"4":  "AuthenticityNet",
	"5":  "PersonaPulse",
	"6":  "VeritySphere",
	"7":  "InsightNet",
	"8":  "EmergeNet",
	"9":  "EgoNet",
	"10": "IdentityNEXA",
}

const AccAddprifix = "ssi"

func GetDefaultDidPrefix(networkID string) string {
	networkName, exists := networkNames[networkID]
	if !exists {
		networkName = "UnknownNetwork" // or handle the error as needed
	}
	return "did:sovid:" + networkName + ":"
}

// findDividerInDid finds the divider "1" in the string and return all characters before "1"
func findPrefixAndDataInDid(s string) (string, string, error) {
	for i, ch := range s {
		if ch == '1' {
			if len(s[i:]) == 0 {
				return s[:i], "", nil
			}
			return s[:i], s[i:], nil
		}
	}
	return "", "", errors.New("divider not found in the string")
}

func GetPrefixFromDidString(did string) (string, error) {
	if len(did) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}

	prifix, _, err := findPrefixAndDataInDid(did)

	if err != nil {
		return "", err
	}
	return prifix, nil
}

func GetDataAfterDivider(did string) (string, error) {

	if len(did) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}
	_, data, err := findPrefixAndDataInDid(did)

	if err != nil {
		return "", err
	}
	// for i, ch := range did {
	// 	if ch == '1' {
	// 		if len(did[i:]) == 0 {
	// 			return "", errors.New("data does not exist")
	// 		}
	// 		return did[i:], nil
	// 	}
	// }
	return data, nil
}
