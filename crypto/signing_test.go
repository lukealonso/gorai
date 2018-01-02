package crypto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignatureSign(t *testing.T) {
	var s [64]byte
	assert(t, hex.Decode(s, []byte("F9083B03B874D971C9949B076364CED50EECC72C41B7EDEB819BBF47A0410AE7")), 32)
	var b [32]byte
	HashBlake2b(b[:], s)
	assert.NotNil(t, h)
	assert.Nil(t, err)
	pub, prv, err := GenerateKeyPair()
	assert.Nil(t, err)

	assert.Nil(t, b.Signature().Sign(prv, &bh))
	assert.Nil(t, err)
	assert.True(t, b.Signature().Verify(pub, &bh))
}

//F9083B03B874D971C9949B076364CED50EECC72C41B7EDEB819BBF47A0410AE7
//xrb_3xiempqttgfhpfpjoknswi5zi8cqmrdejfaa5d3swfisuab9r4ggyjhkb8xn
