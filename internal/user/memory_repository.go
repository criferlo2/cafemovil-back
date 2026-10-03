package user

import "sync"

type MemoryRepository struct {
	mutex sync.RWMutex
	items map[string]User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{items: make(map[string]User)}
}

func (repository *MemoryRepository) Save(user User) (User, error) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()
	repository.items[user.ID] = user
	return user, nil
}
