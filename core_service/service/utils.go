package coreservice

import (
	"crypto/rand"
	"crypto/sha256"
	"unsafe"
)

func generateRandomValue() ([]byte, error) {
	// Erzeugen Sie ein zufälliges Byte-Array mit der gewünschten Länge
	randomBytes := make([]byte, 32) // 32 Bytes für SHA-256

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

func hashRandomValue(value []byte) []byte {
	// Erzeugen Sie einen SHA-256-Hash des übergebenen Werts
	hasher := sha256.New()
	hasher.Write(value)
	return hasher.Sum(nil)
}

func vmProcessPointerId(ptr *ProcessSession) uint64 {
	// Die Adresse des Pointers als uintptr
	pointerAddress := uintptr(unsafe.Pointer(ptr))

	// Konvertieren von uintptr zu uint64
	return uint64(pointerAddress)
}
