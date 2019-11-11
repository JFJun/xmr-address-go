package chainid

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/bcnutil"
	"github.com/jfjun/xmr-address-go/crypto/ed25519/base58p"
	"testing"
)

func Test_CreateAddress(t *testing.T) {
	seed:="cf80ca9a2b4ef7c4b0f6ff51e444f62c2870cc2986741a75bc9b9ea20535e202"
	privSpendKey:=NewPrivateKeyBySeed(seed)
	fmt.Println("P",hex.EncodeToString(privSpendKey.Seed()))
	fmt.Println("pub:",hex.EncodeToString(privSpendKey.Public()))
	//privViewKey:=PrivateSpendToViewKey(privSpendKey.Seed())
	view:="9749e008f545190c603c854c5c25400b116f00fe782d9221950fa630175eb204"
	privViewKey:=NewPrivateKeyBySeed(view)
	fmt.Println("ViewP",hex.EncodeToString(privViewKey.Seed()))
	fmt.Println("ViewPub",hex.EncodeToString(privViewKey.Public()))
	//network_byte:=uint64(3914525)
	//b:=make([]byte,8)
	//binary.BigEndian.PutUint64(b,network_byte)
	//fmt.Println(b)
	//version:=new(big.Int).SetUint64(572238).Bytes()
	//fmt.Println(version)
	//v:=[]byte{0xce,0xf6,0x22}
	//fmt.Println(v)
	buf:=make([]byte,4)
	binary.PutUvarint(buf,3914525)
	var data []byte
	data = append(data,buf...)
	data = append(data,privSpendKey.Public()...)
	data = append(data,privViewKey.Public()...)
	checkSum:=bcnutil.FastHash(data)[:4]

	data = append(data,checkSum...)
	fmt.Println(data)
	address:=base58p.EncodeToString(data)
	fmt.Println(address)
}
func TestPrivateSpendToViewKey(t *testing.T) {
	priv:=NewPrivateSpendOrViewKey()
	fmt.Println(hex.EncodeToString(priv.Seed()))
}