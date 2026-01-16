package xor

func Crypt(data []byte, key []byte) []byte {
	if len(data) == 0 || len(key) == 0 {
		return data
	}

	out := make([]byte, len(data))
	keyLen := len(key)

	for i := 0; i < len(data); i++ {
		out[i] = data[i] ^ key[i%keyLen]
	}

	return out
}
