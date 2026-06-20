package manager

import (
	"encoding/json"
	"os"
)

type State struct {
	Offsets map[string]map[string]int `json:"offsets"`
}

func LoadState() *State {
	s := &State{
		Offsets: map[string]map[string]int{},
	}

	data, err := os.ReadFile("storage/state.json")
	if err == nil {
		_ = json.Unmarshal(data, s)
	}

	return s
}

func (s *State) GetOffset(group, topic string) int {
	if s.Offsets[group] == nil {
		return 0
	}
	return s.Offsets[group][topic]
}

func (s *State) SetOffset(group, topic string, value int) {
	if s.Offsets[group] == nil {
		s.Offsets[group] = map[string]int{}
	}
	s.Offsets[group][topic] = value
}

func (s *State) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("storage/state.json", data, 0644)
}
