package util

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnpqrstuvwxyz0123456789"

func toBase62(num int64) string {
	var result string
	for num > 0 {
		result = string(chars[num%62]) + result
		num = num / 62
	}
	return result
}

func EncryptAndConvertToBase62(id int, salt int64, secret_key string) string {
	salted_id := int64(id) + salt
	return toBase62(salted_id)
}
