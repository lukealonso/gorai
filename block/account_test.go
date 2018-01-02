package gorai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountDecode(t *testing.T) {
	a, err := DecodeAccount("xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuohr3")
	assert.Nil(t, err)
	assert.EqualValues(t, "e89208dd038fbb269987689621d52292ae9c35941a7484756ecced92a65093ba", a.HexString())

	a, err = DecodeAccount("xrb_1111111111111111111111111111111111111111111111111111hifc8npp")
	assert.Nil(t, err)
	assert.EqualValues(t, "0000000000000000000000000000000000000000000000000000000000000000", a.HexString())

	a, err = DecodeAccount("xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mni")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncgt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("zzzz")
	assert.NotNil(t, err)

	a, err = DecodeAccount("zzz_3t6k369ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xb_zZt6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xrb_Zt6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xrb_!t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)

	a, err = DecodeAccount("xrb_ t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuozr3")
	assert.NotNil(t, err)
}

func TestAccountDecodeHex(t *testing.T) {
	a, err := DecodeAccountHex("e89208dd038fbb269987689621d52292ae9c35941a7484756ecced92a65093ba")
	assert.Nil(t, err)
	assert.EqualValues(t, "e89208dd038fbb269987689621d52292ae9c35941a7484756ecced92a65093ba", a.HexString())

	a, err = DecodeAccountHex("00")
	assert.Nil(t, err)
	assert.EqualValues(t, "0000000000000000000000000000000000000000000000000000000000000000", a.HexString())

	a, err = DecodeAccountHex("01")
	assert.Nil(t, err)
	assert.EqualValues(t, "0000000000000000000000000000000000000000000000000000000000000001", a.HexString())

	a, err = DecodeAccountHex("XX")
	assert.NotNil(t, err)

}

func TestAccountEncode(t *testing.T) {
	a, err := DecodeAccountHex("e89208dd038fbb269987689621d52292ae9c35941a7484756ecced92a65093ba")
	assert.Nil(t, err)
	assert.EqualValues(t, "xrb_3t6k35gi95xu6tergt6p69ck76ogmitsa8mnijtpxm9fkcm736xtoncuohr3", a.String())
	a, err = DecodeAccountHex("00")
	assert.Nil(t, err)
	assert.EqualValues(t, "xrb_1111111111111111111111111111111111111111111111111111hifc8npp", a.String())
}
