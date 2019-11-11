package params

import "strings"

type MainParams struct {
	COIN_TYPE	string
	CRYPTONOTE_PUBLIC_ADDRESS_BASE58_PREFIX uint64
}
const (
	XMR = "xmr"
	BCN = "bcn"
	TRTL = "trtl"
)

var(
	Params *MainParams

	xmrParams = &MainParams{
		COIN_TYPE:XMR,
		CRYPTONOTE_PUBLIC_ADDRESS_BASE58_PREFIX:18,

	}
	bcnParams = &MainParams{
		COIN_TYPE:BCN,
		CRYPTONOTE_PUBLIC_ADDRESS_BASE58_PREFIX:572238,
	}
	trtlParams = &MainParams{
		COIN_TYPE:TRTL,
		CRYPTONOTE_PUBLIC_ADDRESS_BASE58_PREFIX:3914525,
	}
)

func SelectParams(coinType string)*MainParams{
	switch strings.ToLower(coinType) {
	case XMR:
		Params = xmrParams
	case BCN:
		Params = bcnParams
	case TRTL:
		Params = trtlParams
	}
	return nil
}