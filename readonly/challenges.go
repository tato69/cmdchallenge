package main

import (
	"encoding/json"
	"errors"
)

type challenge struct {
	Slug           string         `json:"slug"`
	Version        int            `json:"version"`
	Author         string         `json:"author"`
	Description    string         `json:"description"`
	Example        string         `json:"example"`
	ExpectedOutput expectedOutput `json:"expected_output"`
}

type challenges []challenge

func (ch challenges) Len() int {
	return len(ch)
}

func (ch challenges) Slug(slug string) (challenge, error) {
	for _, c := range ch {
		if c.Slug == slug {
			return c, nil
		}
	}
	return challenge{}, errors.New("unable to find slug")
}

func (p *challenge) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return true
}
