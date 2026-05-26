package memory

import (
	"github.com/lf-edge/ekuiper/v2/pkg/kv"
	"github.com/lf-edge/ekuiper/v2/pkg/syncx"
)

type memoryKvStore struct {
	data map[string]string
	mu   syncx.RWMutex
}

func NewMemoryKV() kv.KeyValue { _ = "STUB: not implemented"; return *new(kv.KeyValue) }

func (m *memoryKvStore) Open() error { _ = "STUB: not implemented"; return nil }

func (m *memoryKvStore) Close() error { _ = "STUB: not implemented"; return nil }

func (m *memoryKvStore) Set(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryKvStore) Setnx(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryKvStore) Get(key string, value interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *memoryKvStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (m *memoryKvStore) Keys() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *memoryKvStore) All() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a copy to avoid race conditions if the caller modifies the map (though signature returns map[string]string, usually it's better to copy)
// But matching the interface, let's just return a copy.

func (m *memoryKvStore) Clean() error { _ = "STUB: not implemented"; return nil }

func (m *memoryKvStore) Drop() error { _ = "STUB: not implemented"; return nil }

func (m *memoryKvStore) SetKeyedState(key string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryKvStore) GetKeyedState(key string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *memoryKvStore) GetByPrefix(prefix string) (map[string][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
