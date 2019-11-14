package cipher

// Cipher describes a type for encrypting and decrypting strings
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
