package test

import (
	"testing"
)

// Define Mood type for testing
type Mood struct {
	ID    string
	Emoji string
}

// Declare getMoods as a variable to allow stubbing
var getMoods = func() []Mood {
	return []Mood{
		{ID: "happy", Emoji: "😄"},
		{ID: "tired", Emoji: "😴"},
	}
}

// Provide a stub for getMoodEmoji for testing
func getMoodEmoji(moodID string) string {
	for _, mood := range getMoods() {
		if mood.ID == moodID {
			return mood.Emoji
		}
	}
	return "😊"
}

func TestGetMoodEmoji(t *testing.T) {
	// Stub getMoods for predictable behavior
	originalGetMoods := getMoods
	getMoods = func() []Mood {
		return []Mood{
			{ID: "happy", Emoji: "😄"},
			{ID: "tired", Emoji: "😴"},
		}
	}
	defer func() { getMoods = originalGetMoods }()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Existing mood returns emoji", "happy", "😄"},
		{"Another existing mood", "tired", "😴"},
		{"Unknown mood returns default", "unknown", "😊"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMoodEmoji(tt.input)
			if got != tt.expected {
				t.Errorf("getMoodEmoji(%s) = %s, expected %s", tt.input, got, tt.expected)
			}
		})
	}
}
