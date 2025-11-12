package store

type RAMStore struct {
}

// NewRAMStore crée une nouvelle instance de RAMStore.
func NewRAMStore() *RAMStore {
	return &RAMStore{}
}
