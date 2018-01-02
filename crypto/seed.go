package crypto

func DeterministicKey(seed []byte, index uint64) {

}

// blake2b_state hash;
// blake2b_init (&hash, prv_a.bytes.size ());
// blake2b_update (&hash, seed_a.bytes.data (), seed_a.bytes.size ());
// rai::uint256_union index (index_a);
// blake2b_update (&hash, reinterpret_cast <uint8_t *> (&index.dwords [7]), sizeof (uint32_t));
// blake2b_final (&hash, prv_a.bytes.data (), prv_a.bytes.size ());
