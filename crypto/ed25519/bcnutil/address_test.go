package bcnutil

import "testing"

func TestVerifyAddress(t *testing.T) {
	err:=VerifyAddress("TRTLuxMkzNxjhoX4aVZwWdhZzB8nYpAzTT51jfiuRGicbnfi1ZwEjAPfPibhsBZ7Um65dU8QBwLzM6s7FJUeRrRmgeyDEe3k9PJ",3914525)
	if err != nil {
		panic(err)
	}
}


