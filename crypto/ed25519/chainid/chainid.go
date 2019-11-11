package chainid

import (
	cryptorand "crypto/rand"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/jfjun/xmr-address-go/crypto/ed25519"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/internal/edwards25519"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/base58p"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/bcnutil"
	"github.com/jfjun/xmr-address-go/params"
	"io"
	"log"
	"strconv"
)

/*
门罗币，字节币，乌龟币私钥生成
*/


/*
seed与私钥相同
*/
func NewPrivateKeyBySeed(seed string)ed25519.PrivateKey{
	seedBytes,err:=hex.DecodeString(seed)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyByPrivBytes(seedBytes)
}
func NewPrivateKeyByPrivBytes(privBytes []byte)ed25519.PrivateKey{
	if len(privBytes)==0 {
		panic(fmt.Errorf("private key bytes length is equal 0"))
	}
	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], privBytes[:32])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	var publicKeyBytes [32]byte
	A.ToBytes(&publicKeyBytes)

	privateKey := make([]byte,ed25519.PrivateKeySize)
	copy(privateKey, privBytes[:32])
	copy(privateKey[32:], publicKeyBytes[:])
	return privateKey
}


func NewPrivateSpendOrViewKey()ed25519.PrivateKey{
	seed := make([]byte, ed25519.SeedSize)
	if _, err := io.ReadFull(cryptorand.Reader, seed); err != nil {
		return nil
	}
	digest := sha512.Sum512(seed)
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64
	var hb [64]byte
	copy(hb[:],digest[:])
	var out [32]byte
	//C++ 源码有一个sc_check 所以需要进行sc_reduce计算得到私钥
	edwards25519.ScReduce(&out,&hb)

	privBytes:=make([]byte,32)
	copy(privBytes[:],out[:])

	return NewPrivateKeyByPrivBytes(privBytes)
}

func PrivateSpendToViewKey(privSpendKeybytes []byte)ed25519.PrivateKey{

	//先进性sha3-Keccak256 hash
	if l := len(privSpendKeybytes); l != ed25519.SeedSize {
		panic("ed25519: bad seed length: " + strconv.Itoa(l))
	}
	digest := sha512.Sum512(privSpendKeybytes)
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64
	return NewPrivateKeyByPrivBytes(digest[:])
}


func ToAddress(pubSpendKey,pubViewKey[]byte)string{
	if len(pubSpendKey)!=32 ||len(pubViewKey)!=32 {
		log.Fatal("Pub spend or view key length is not equal 32")
		return ""
	}
	buf:=make([]byte,4)
	binary.PutUvarint(buf,params.Params.CRYPTONOTE_PUBLIC_ADDRESS_BASE58_PREFIX)
	var data []byte
	data = append(data,buf...)
	data = append(data,pubSpendKey...)
	data = append(data,pubViewKey...)
	checkSum:=bcnutil.FastHash(data)[:4]
	data = append(data,checkSum...)
	return base58p.EncodeToString(data)
}
