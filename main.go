package main

import (
	"encoding/hex"
	"fmt"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/chainid"
	"github.com/jfjun/xmr-address-go/params"
)

func createXMRAddress(){
	params.SelectParams("xmr")
	privSpendKey:=chainid.NewPrivateSpendOrViewKey()
	privViewKey:=chainid.NewPrivateSpendOrViewKey()
	address:=chainid.ToAddress(privSpendKey.Public(),privViewKey.Public())
	fmt.Println("Private Spend Key: ",hex.EncodeToString(privSpendKey.Seed()))
	fmt.Println("Public Spend Key: ",hex.EncodeToString(privSpendKey.Public()))
	fmt.Println("Private View Key: ",hex.EncodeToString(privViewKey.Seed()))
	fmt.Println("Public View Key: ",hex.EncodeToString(privViewKey.Public()))
	fmt.Println("Address: ",address)
}
func createBCNAddress(){
	params.SelectParams("bcn")
	privSpendKey:=chainid.NewPrivateSpendOrViewKey()
	privViewKey:=chainid.NewPrivateSpendOrViewKey()
	address:=chainid.ToAddress(privSpendKey.Public(),privViewKey.Public())
	fmt.Println("Private Spend Key: ",hex.EncodeToString(privSpendKey.Seed()))
	fmt.Println("Public Spend Key: ",hex.EncodeToString(privSpendKey.Public()))
	fmt.Println("Private View Key: ",hex.EncodeToString(privViewKey.Seed()))
	fmt.Println("Public View Key: ",hex.EncodeToString(privViewKey.Public()))
	fmt.Println("Address: ",address)
}
func createTRTLAddress(){
	params.SelectParams("trtl")
	privSpendKey:=chainid.NewPrivateSpendOrViewKey()
	privViewKey:=chainid.NewPrivateSpendOrViewKey()
	address:=chainid.ToAddress(privSpendKey.Public(),privViewKey.Public())
	fmt.Println("Private Spend Key: ",hex.EncodeToString(privSpendKey.Seed()))
	fmt.Println("Public Spend Key: ",hex.EncodeToString(privSpendKey.Public()))
	fmt.Println("Private View Key: ",hex.EncodeToString(privViewKey.Seed()))
	fmt.Println("Public View Key: ",hex.EncodeToString(privViewKey.Public()))
	fmt.Println("Address: ",address)
}
func main() {
	//createBCNAddress()
	//createTRTLAddress()
	createXMRAddress()

}
