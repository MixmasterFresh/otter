package auth

import(
    "time"
    "math/rand"
)

const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
    letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+="
)

// GenerateKey generates a random key of a given length
func GenerateKey(n int) string {
    var src = rand.NewSource(time.Now().UnixNano())
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; i-- {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        idx := int(cache & letterIdxMask)
        b[i] = letterBytes[idx]
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}