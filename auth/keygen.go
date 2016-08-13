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

var random64 chan int64
var goroutineCount int
var mutex chan bool

func startKeyGeneration() {
    random64 = make(chan int64, 1024)
    mutex = make(chan bool)
    mutex <- true
    go generateRandoms()
}

// GenerateKey generates a random key of a given length
func GenerateKey(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, getRandom63(), letterIdxMax; i >= 0; i-- {
        if remain == 0 {
            cache, remain = getRandom63(), letterIdxMax
        }
        idx := int(cache & letterIdxMask)
        b[i] = letterBytes[idx]
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}

func getRandom63() int64 {
    select {
    case random := <- random64:
        return random
    default:
        select {
            case <- mutex:
                go generateRandoms()
                time.Sleep(10 * time.Millisecond)
                mutex <- true
                return <- random64
            default:
                return <- random64
        }
    }
}

func generateRandoms() {
    var src = rand.NewSource(time.Now().UnixNano())
    for {
        random64 <- src.Int63()
    }
}

