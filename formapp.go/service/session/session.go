package session
 
import (
    "fmt"
    "github.com/google/uuid"
)
 
type State struct {
    Name string
    Birthday string
    Message string
}
 
var store = make(map[string]*State)
 
type Session struct {
    id string
}
 
func NewSession() (Session, error) {
    id, err := uuid.NewRandom()
    if err != nil {
        return Session{}, err
    }
    session := Session{ id.String() }
    return session, nil
}
 
func (s *Session) ID() string {
    return s.id
}
 
func (s *Session) GetState() (State, error) {
    if _, err := uuid.Parse(s.ID()); err != nil {
        return State{}, fmt.Errorf("Invalid session ID")
    }
    state, exist := store[s.ID()]
    if !exist {
        state = new(State)
        store[s.ID()] = state
    }
    return *state, nil
}
 
func (s *Session) SetState(ns State) error {
    state, exist := store[s.ID()]
    if !exist {
        return fmt.Errorf("State corresponds to %s does not exist", s.ID())
    }
    *state = ns
    return nil
}
 
func (s *Session) Close() {
    delete(store, s.ID())
    s.id = ""
}