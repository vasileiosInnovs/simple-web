package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Mood represents a mood option
type Mood struct {
	ID    string
	Emoji string
	Label string
}

// GreetingData holds data for the greeting page
type GreetingData struct {
	Name       string
	Mood       string
	Emoji      string
	TimeOfDay  string
	Message    string
	Background template.CSS
}

// Get all available moods
func getMoods() []Mood {
	return []Mood{
		{ID: "happy", Emoji: "😊", Label: "Happy"},
		{ID: "tired", Emoji: "😴", Label: "Tired"},
		{ID: "excited", Emoji: "🤩", Label: "Excited"},
		{ID: "grumpy", Emoji: "😠", Label: "Grumpy"},
		{ID: "peaceful", Emoji: "😌", Label: "Peaceful"},
	}
}

// Get time of day
func getTimeOfDay() string {
	hour := time.Now().Hour()
	if hour < 12 {
		return "morning"
	} else if hour < 18 {
		return "afternoon"
	}
	return "evening"
}

// Get mood emoji
func getMoodEmoji(moodID string) string {
	moods := getMoods()
	for _, mood := range moods {
		if mood.ID == moodID {
			return mood.Emoji
		}
	}
	return "😊"
}

// Get greeting message based on mood
func getGreetingMessage(name, moodID, timeOfDay string) string {
	messages := map[string]string{
		"happy":    "Your happiness is contagious today! ✨",
		"tired":    "Take it easy, rest is important. 💤",
		"excited":  "Your excitement is electrifying! ⚡",
		"grumpy":   "Hope your day gets better soon. ☁️",
		"peaceful": "May your calm carry you through. 🍃",
	}

	message, exists := messages[moodID]
	if !exists {
		message = "Have a wonderful day!"
	}

	return "Good " + timeOfDay + ", " + name + "! " + message
}

// Get background color based on mood
func getBackgroundColor(moodID string) string {
	colors := map[string]string{
		"happy":    "linear-gradient(135deg, #FCD34D 0%, #F59E0B 100%)",
		"tired":    "linear-gradient(135deg, #818CF8 0%, #6366F1 100%)",
		"excited":  "linear-gradient(135deg, #F472B6 0%, #EC4899 100%)",
		"grumpy":   "linear-gradient(135deg, #94A3B8 0%, #64748B 100%)",
		"peaceful": "linear-gradient(135deg, #86EFAC 0%, #34D399 100%)",
	}

	color, exists := colors[moodID]
	if !exists {
		color = "linear-gradient(135deg, #F3F4F6 0%, #E5E7EB 100%)"
	}

	return color
}

// Handler for the home page (form)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/form.html"))

	data := struct {
		Moods []Mood
	}{
		Moods: getMoods(),
	}

	tmpl.Execute(w, data)
}

// Handler for greeting submission
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get form values
	name := r.FormValue("name")
	mood := r.FormValue("mood")

	// Validate
	if name == "" || mood == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Prepare greeting data
	timeOfDay := getTimeOfDay()

	background := getBackgroundColor(mood)

	fmt.Printf("Mood: %s\n", mood)
	fmt.Printf("Background: %s\n", background)

	data := GreetingData{
		Name:       name,
		Mood:       mood,
		Emoji:      getMoodEmoji(mood),
		TimeOfDay:  timeOfDay,
		Message:    getGreetingMessage(name, mood, timeOfDay),
		Background: template.CSS(getBackgroundColor(mood)),
	}

	// DEBUG: Print entire data struct
	fmt.Printf("Data: %+v\n", data)

	// Render greeting template
	tmpl := template.Must(template.ParseFiles("templates/greeting.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Printf("Template error: %v\n", err)
	}
}

func main() {
	// Register routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/greet", greetHandler)

	// Start server
	println("🚀 Server starting on http://localhost:8080")
	println("Press Ctrl+C to stop")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
