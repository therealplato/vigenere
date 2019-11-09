package vigenere

func Encode(plaintext []byte, key []byte) []byte {
	var (
		m_i    int
		k_i    int
		m      int
		k      int
		c      int
		m_byte byte
	)

	msg := strip(plaintext)
	ciphertext := make([]byte, len(msg))

	for m_i, m_byte = range msg {
		k_i = m_i % len(key)
		m = atoi[m_byte]
		k = atoi[key[k_i]]
		c = (m + k) % 26
		ciphertext[m_i] = itoa[c]
	}
	return ciphertext
}

func Decode(ciphertext []byte, key []byte) []byte {
	var (
		c_i    int
		k_i    int
		m      int
		k      int
		c      int
		c_byte byte
	)

	msg := strip(ciphertext)
	plaintext := make([]byte, len(msg))

	for c_i, c_byte = range msg {
		k_i = c_i % len(key)
		c = atoi[c_byte]
		k = atoi[key[k_i]]
		m = (26 + c - k) % 26
		plaintext[c_i] = itoa[m]
	}
	return plaintext
}
