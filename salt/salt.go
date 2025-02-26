package salt

import "crypto/rand"

const charset = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"

// GenSalt 随机生成密码盐
func GenSalt(saltLen int) string {
	randBytes := make([]byte, saltLen)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "000000"
	}

	for i := 0; i < saltLen; i++ {
		randBytes[i] = charset[int(randBytes[i])%len(charset)]
	}

	return string(randBytes)
}
