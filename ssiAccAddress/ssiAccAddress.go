package ssibech32address

import (
	"fmt"
	"strings"

	// "github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/harrybite/ssiBech32Address/types"
)

func VerifyPrefixFormat(did string) (string, error) {

	didprifix := types.GetDefaultDidPrefix("3")

	if len(strings.TrimSpace(did)) == 0 {
		return "", fmt.Errorf("empty did string is not allowed")
	}

	fountprefix, err := types.GetPrefixFromDidString(did)

	if err != nil {
		return "", err
	}

	if fountprefix != didprifix {
		return "", fmt.Errorf("invalid did prefix; expected %s, got %s", didprifix, fountprefix)
	}

	return fountprefix, nil
}

func VerifyDidData(did string) (string, error) {

	data, detaErr := types.GetDataAfterDivider(did)

	if detaErr != nil {
		return "", detaErr
	}
	return data, nil
}

func VerifyDidFormat(did string) (bool, error) {

	_, Prefixerr := VerifyPrefixFormat(did)
	if Prefixerr != nil {
		return false, Prefixerr
	}

	data, Dataerr := VerifyDidData(did)
	if Dataerr != nil {
		return false, Dataerr
	}

	_, _, err := bech32.DecodeAndConvert(types.AccAddprifix + data)
	if err != nil {
		return false, err
	}

	// _, _, err := bech32.Decode(types.AccAddprifix + data)
	// if err != nil {
	// 	return false, err
	// }

	return true, nil
}
