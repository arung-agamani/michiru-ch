package auth

import (
	"sync"
)

var sessionStore = struct {
	sync.RWMutex
	sessions map[string]string
}{sessions: make(map[string]string)}

func SetSession(key, value string) {
	sessionStore.Lock()
	defer sessionStore.Unlock()
	sessionStore.sessions[key] = value
}

func GetSession(key string) (string, bool) {
	sessionStore.RLock()
	defer sessionStore.RUnlock()
	value, ok := sessionStore.sessions[key]
	return value, ok
}
