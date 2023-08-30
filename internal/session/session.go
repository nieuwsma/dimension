package session

import "sync"

type ResourceType int

const (
	Player ResourceType = iota
	Game
	TrainingSession
)

type Resource struct {
	Type       ResourceType
	Identifier string // e.g., gameID, playerID, or trainingSessionID
}

func (r Resource) Equals(other Resource) bool {
	return r.Type == other.Type && r.Identifier == other.Identifier
}

type Token string

// SessionProvider represents the methods required for session management.
type SessionProvider interface {
	AddToken(token Token, resource Resource) error
	RemoveToken(token Token) error
	//GetTokenResource(session Token) (Resource, bool)
	ValidateToken(token Token, resource Resource) bool
}

// MemorySessionStore is an in-memory storage for tokens.
type MemorySessionStore struct {
	tokens map[Token]Resource
	mu     sync.RWMutex // For concurrent safety
}

// NewInMemorySessionStore initializes a new in-memory session store.
func NewInMemorySessionStore() *MemorySessionStore {
	return &MemorySessionStore{
		tokens: make(map[Token]Resource),
	}
}

// AddToken adds a new session and its associated resource.
func (s *MemorySessionStore) AddToken(token Token, resource Resource) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Optional: Check if session already exists and return an error if it does.
	s.tokens[token] = resource
	return nil
}

// RemoveToken removes a session.
func (s *MemorySessionStore) RemoveToken(token Token) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tokens, token)
	return nil
}

//// GetTokenResource retrieves the resource associated with a session.
//func (s *MemorySessionStore) GetTokenResource(session Token) (Resource, bool) {
//	s.mu.RLock()
//	defer s.mu.RUnlock()
//
//	resource, exists := s.tokens[session]
//	return resource, exists
//}

func (s *MemorySessionStore) ValidateToken(token Token, resource Resource) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	retrievedResource, exists := s.tokens[token]

	if !exists {
		return false
	}

	return resource.Equals(retrievedResource)
}
