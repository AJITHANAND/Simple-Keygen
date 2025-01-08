package utils

import (
    "crypto/rand"
    "encoding/hex"
    "sync"
    "time"
)

var (
    currentKey  string
    lastRotated time.Time
    mutex       sync.Mutex
)

func init() {
    RotateKey()
    go func() {
        for {
            time.Sleep(24 * time.Hour)
            RotateKey()
        }
    }()
}

func GenerateRandomKey() string {
    key := make([]byte, 16)
    _, err := rand.Read(key)
    if err != nil {
        return ""
    }
    return hex.EncodeToString(key)
}

func RotateKey() {
    mutex.Lock()
    defer mutex.Unlock()
    currentKey = GenerateRandomKey()
    lastRotated = time.Now()
}

func GetCurrentKey() (string, time.Time) {
    mutex.Lock()
    defer mutex.Unlock()
    return currentKey, lastRotated
}