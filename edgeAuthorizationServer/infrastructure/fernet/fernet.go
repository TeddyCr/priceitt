package fernet

import (
	"sync"

	"github.com/TeddyCr/priceitt/models/config"
	"github.com/fernet/fernet-go"
)

var lock = &sync.Mutex{}

type Fernet struct {
	Key []*fernet.Key
	Salt []byte
}

var instance *Fernet

func GetInstance() *Fernet {
	if instance == nil {
		panic("Fernet not initialized. Call fernet.Initialize() before using fernet.GetInstance()")
	}

	return instance
}

func Initialize(config config.FernetConfig) {
	lock.Lock()
	defer lock.Unlock()

	instance = &Fernet{
		fernet.MustDecodeKeys(config.Key),
		[]byte(config.Salt),
	}
}