package gorai

import (
	"encoding/hex"
	"fmt"
)

type Account struct {
	Key [32]byte
}

const accountLookup = "13456789abcdefghijkmnopqrstuwxyz"
const accountReverse = "~0~1234567~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~89:;<=>?@AB~CDEFGHIJK~LMNO~~~~~"

func DecodeAccount(text string) (*Account, error) {
	a := &Account{}
	err := a.UnmarshalText([]byte(text))
	return a, err
}

func DecodeAccountHex(text string) (*Account, error) {
	b, err := hex.DecodeString(text)
	if err != nil || len(b) > 32 {
		return nil, fmt.Errorf("unable to decode account '%s': bad hex string", string(text))
	}
	a := &Account{}
	copy(a.Key[(32-len(b)):], b[0:])
	return a, nil
}

func (a *Account) UnmarshalText(text []byte) error {
	if len(text) != 64 {
		return fmt.Errorf("unable to decode account '%s': invalid character count", string(text))
	}

	if string(text[0:3]) != "xrb" || (text[3] != '_' && text[3] != '-') {
		return fmt.Errorf("unable to decode account '%s': invalid prefix", string(text))
	}

	var v [60]uint8
	for i := 4; i < 64; i++ {
		c := text[i]
		if c < 0x30 || c >= 0x80 {
			return fmt.Errorf("unable to decode account '%s': invalid character", string(text))
		}
		r := accountReverse[c-0x30]
		if r == '~' {
			return fmt.Errorf("unable to decode account '%s': invalid character", string(text))
		}
		v[i-4] = r - 0x30
	}

	b := &a.Key
	b[0] = (v[2] >> 3) + ((v[1] & 31) << 2) + ((v[0] & 1) << 7)
	b[1] = v[3] + ((v[2] & 7) << 5)
	b[2] = (v[5] >> 2) + ((v[4] & 31) << 3)
	b[3] = (v[7] >> 4) + ((v[6] & 31) << 1) + ((v[5] & 3) << 6)
	b[4] = (v[8] >> 1) + ((v[7] & 15) << 4)
	b[5] = (v[10] >> 3) + ((v[9] & 31) << 2) + ((v[8] & 1) << 7)
	b[6] = v[11] + ((v[10] & 7) << 5)
	b[7] = (v[13] >> 2) + ((v[12] & 31) << 3)
	b[8] = (v[15] >> 4) + ((v[14] & 31) << 1) + ((v[13] & 3) << 6)
	b[9] = (v[16] >> 1) + ((v[15] & 15) << 4)
	b[10] = (v[18] >> 3) + ((v[17] & 31) << 2) + ((v[16] & 1) << 7)
	b[11] = v[19] + ((v[18] & 7) << 5)
	b[12] = (v[21] >> 2) + ((v[20] & 31) << 3)
	b[13] = (v[23] >> 4) + ((v[22] & 31) << 1) + ((v[21] & 3) << 6)
	b[14] = (v[24] >> 1) + ((v[23] & 15) << 4)
	b[15] = (v[26] >> 3) + ((v[25] & 31) << 2) + ((v[24] & 1) << 7)
	b[16] = v[27] + ((v[26] & 7) << 5)
	b[17] = (v[29] >> 2) + ((v[28] & 31) << 3)
	b[18] = (v[31] >> 4) + ((v[30] & 31) << 1) + ((v[29] & 3) << 6)
	b[19] = (v[32] >> 1) + ((v[31] & 15) << 4)
	b[20] = (v[34] >> 3) + ((v[33] & 31) << 2) + ((v[32] & 1) << 7)
	b[21] = v[35] + ((v[34] & 7) << 5)
	b[22] = (v[37] >> 2) + ((v[36] & 31) << 3)
	b[23] = (v[39] >> 4) + ((v[38] & 31) << 1) + ((v[37] & 3) << 6)
	b[24] = (v[40] >> 1) + ((v[39] & 15) << 4)
	b[25] = (v[42] >> 3) + ((v[41] & 31) << 2) + ((v[40] & 1) << 7)
	b[26] = v[43] + ((v[42] & 7) << 5)
	b[27] = (v[45] >> 2) + ((v[44] & 31) << 3)
	b[28] = (v[47] >> 4) + ((v[46] & 31) << 1) + ((v[45] & 3) << 6)
	b[29] = (v[48] >> 1) + ((v[47] & 15) << 4)
	b[30] = (v[50] >> 3) + ((v[49] & 31) << 2) + ((v[48] & 1) << 7)
	b[31] = v[51] + ((v[50] & 7) << 5)

	checksum := uint64(0)
	checksum |= uint64((v[53]>>2)+((v[52]&31)<<3)) << 32
	checksum |= uint64((v[55]>>4)+((v[54]&31)<<1)+((v[53]&3)<<6)) << 24
	checksum |= uint64((v[56]>>1)+((v[55]&15)<<4)) << 16
	checksum |= uint64((v[58]>>3)+((v[57]&31)<<2)+((v[56]&1)<<7)) << 8
	checksum |= uint64(v[59] + ((v[58] & 7) << 5))

	if c, err := a.Checksum(); err != nil || c != checksum {
		return fmt.Errorf("unable to decode account: bad checksum")
	}

	return nil
}

func (a *Account) MarshalText() ([]byte, error) {
	hv := a.checksumBytes()

	unusedBits := uint8(0)

	b := &a.Key
	var v [64]uint8

	v[0] = 'x'
	v[1] = 'r'
	v[2] = 'b'
	v[3] = '_'

	// Account
	v[4] = (((b[0] >> 7) & 1) << 0) + (((unusedBits >> 0) & 15) << 1)
	v[5] = (((b[0] >> 2) & 31) << 0)
	v[6] = (((b[1] >> 5) & 7) << 0) + (((b[0] >> 0) & 3) << 3)
	v[7] = (((b[1] >> 0) & 31) << 0)
	v[8] = (((b[2] >> 3) & 31) << 0)
	v[9] = (((b[3] >> 6) & 3) << 0) + (((b[2] >> 0) & 7) << 2)
	v[10] = (((b[3] >> 1) & 31) << 0)
	v[11] = (((b[4] >> 4) & 15) << 0) + (((b[3] >> 0) & 1) << 4)
	v[12] = (((b[5] >> 7) & 1) << 0) + (((b[4] >> 0) & 15) << 1)
	v[13] = (((b[5] >> 2) & 31) << 0)
	v[14] = (((b[6] >> 5) & 7) << 0) + (((b[5] >> 0) & 3) << 3)
	v[15] = (((b[6] >> 0) & 31) << 0)
	v[16] = (((b[7] >> 3) & 31) << 0)
	v[17] = (((b[8] >> 6) & 3) << 0) + (((b[7] >> 0) & 7) << 2)
	v[18] = (((b[8] >> 1) & 31) << 0)
	v[19] = (((b[9] >> 4) & 15) << 0) + (((b[8] >> 0) & 1) << 4)
	v[20] = (((b[10] >> 7) & 1) << 0) + (((b[9] >> 0) & 15) << 1)
	v[21] = (((b[10] >> 2) & 31) << 0)
	v[22] = (((b[11] >> 5) & 7) << 0) + (((b[10] >> 0) & 3) << 3)
	v[23] = (((b[11] >> 0) & 31) << 0)
	v[24] = (((b[12] >> 3) & 31) << 0)
	v[25] = (((b[13] >> 6) & 3) << 0) + (((b[12] >> 0) & 7) << 2)
	v[26] = (((b[13] >> 1) & 31) << 0)
	v[27] = (((b[14] >> 4) & 15) << 0) + (((b[13] >> 0) & 1) << 4)
	v[28] = (((b[15] >> 7) & 1) << 0) + (((b[14] >> 0) & 15) << 1)
	v[29] = (((b[15] >> 2) & 31) << 0)
	v[30] = (((b[16] >> 5) & 7) << 0) + (((b[15] >> 0) & 3) << 3)
	v[31] = (((b[16] >> 0) & 31) << 0)
	v[32] = (((b[17] >> 3) & 31) << 0)
	v[33] = (((b[18] >> 6) & 3) << 0) + (((b[17] >> 0) & 7) << 2)
	v[34] = (((b[18] >> 1) & 31) << 0)
	v[35] = (((b[19] >> 4) & 15) << 0) + (((b[18] >> 0) & 1) << 4)
	v[36] = (((b[20] >> 7) & 1) << 0) + (((b[19] >> 0) & 15) << 1)
	v[37] = (((b[20] >> 2) & 31) << 0)
	v[38] = (((b[21] >> 5) & 7) << 0) + (((b[20] >> 0) & 3) << 3)
	v[39] = (((b[21] >> 0) & 31) << 0)
	v[40] = (((b[22] >> 3) & 31) << 0)
	v[41] = (((b[23] >> 6) & 3) << 0) + (((b[22] >> 0) & 7) << 2)
	v[42] = (((b[23] >> 1) & 31) << 0)
	v[43] = (((b[24] >> 4) & 15) << 0) + (((b[23] >> 0) & 1) << 4)
	v[44] = (((b[25] >> 7) & 1) << 0) + (((b[24] >> 0) & 15) << 1)
	v[45] = (((b[25] >> 2) & 31) << 0)
	v[46] = (((b[26] >> 5) & 7) << 0) + (((b[25] >> 0) & 3) << 3)
	v[47] = (((b[26] >> 0) & 31) << 0)
	v[48] = (((b[27] >> 3) & 31) << 0)
	v[49] = (((b[28] >> 6) & 3) << 0) + (((b[27] >> 0) & 7) << 2)
	v[50] = (((b[28] >> 1) & 31) << 0)
	v[51] = (((b[29] >> 4) & 15) << 0) + (((b[28] >> 0) & 1) << 4)
	v[52] = (((b[30] >> 7) & 1) << 0) + (((b[29] >> 0) & 15) << 1)
	v[53] = (((b[30] >> 2) & 31) << 0)
	v[54] = (((b[31] >> 5) & 7) << 0) + (((b[30] >> 0) & 3) << 3)
	v[55] = (((b[31] >> 0) & 31) << 0)

	// Checksum
	v[56] = (((hv[4] >> 3) & 31) << 0)
	v[57] = (((hv[3] >> 6) & 3) << 0) + (((hv[4] >> 0) & 7) << 2)
	v[58] = (((hv[3] >> 1) & 31) << 0)
	v[59] = (((hv[2] >> 4) & 15) << 0) + (((hv[3] >> 0) & 1) << 4)
	v[60] = (((hv[1] >> 7) & 1) << 0) + (((hv[2] >> 0) & 15) << 1)
	v[61] = (((hv[1] >> 2) & 31) << 0)
	v[62] = (((hv[0] >> 5) & 7) << 0) + (((hv[1] >> 0) & 3) << 3)
	v[63] = (((hv[0] >> 0) & 31) << 0)

	for i := 4; i < 64; i++ {
		v[i] = accountLookup[v[i]&0x1F]
	}
	return v[:], nil
}

func (a *Account) checksumBytes() []byte {
	var checksum [5]byte
	hashBlake2b(checksum[:], a.Key[:])
	return checksum[:]
}

func (a *Account) Checksum() (uint64, error) {
	hv := a.checksumBytes()
	return uint64(hv[0]) + (uint64(hv[1]) << 8) + (uint64(hv[2]) << 16) + (uint64(hv[3]) << 24) + (uint64(hv[4]) << 32), nil
}

func (a *Account) String() string {
	b, err := a.MarshalText()
	if err != nil {
		// Shouldn't happen, our MarshalText returns no errors.
		panic(err)
	}
	return string(b)
}

func (a *Account) HexString() string {
	return hex.EncodeToString(a.Key[:])
}

func (a *Account) PublicKey() []byte {
	return a.Key[:]
}

func (a *Account) Bytes() []byte {
	return a.Key[:]
}
