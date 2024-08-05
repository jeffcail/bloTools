package bloTools

import "crypto/sha256"

// Sha256Checksum 计算 SHA-256 校验和
func Sha256Checksum(data []byte) [32]byte {
	f := Sha256Hash(data)
	return Sha256Hash(f[:])
}

// Sha256Hash 计算 SHA-256 哈希
func Sha256Hash(data []byte) [32]byte {
	hash := sha256.New()
	hash.Write(data)
	var res [32]byte
	copy(res[:], hash.Sum(nil))
	return res
}
