# 🌤️ Mood Greeting App – Getting Started with Go Templates

A simple beginner-friendly web app built using Go (Golang) and HTML templates.
Users enter their name and mood, and the app generates a personalized greeting with:

Dynamic emoji

Mood-based message

Mood-based background color

This project demonstrates how to build a minimal Go web server with template rendering — perfect for beginners.

## 📌 Features

🌈 Mood selection form (happy, sad, excited, calm, etc.)

🎭 Dynamic emoji + greeting message

🎨 Background color changes based on mood

🧩 Uses Go’s built-in html/template package

⚡ Very lightweight — no external frameworks needed

## 🛠️ Tech Stack
Component	Technology
Language	Go (Golang)
Templates	Go html/template
Editor	VS Code + Go Extension
Browser	Any modern browser

## 📁 Project Structure
```
mood-greeting-app/
├── main.go
└── templates/
    ├── form.html
    └── greeting.html
```

## 🚀 Getting Started
### 1️⃣ Install Go

Download Go (1.20+) from:
👉 https://go.dev/dl/

Verify installation:
```
go version
```

### 2️⃣ Create the Project Directory

```
mood-greeting-app/
│── main.go
└── templates/
     ├── form.html
     └── greeting.html
```

### 3️⃣ Run the Server
```
go run main.go
```


Open the app in your browser:

👉 http://localhost:8080

## 🧪 Minimal Working Example (main.go)

Here’s what the core logic does:

Handles / route → displays form

Handles /greet route → processes name + mood

Sends data to HTML templates

Applies mood-based emoji + colors

A snippet:
```
http.HandleFunc("/", homeHandler)
http.HandleFunc("/greet", greetHandler)

fmt.Println("🚀 Server running at http://localhost:8080")
http.ListenAndServe(":8080", nil)
```

### 🎨 Expected Output

After submitting the form, you’ll see something like:
```
🌞 Good morning, John!
Your happiness is contagious today! ✨
(With a bright yellow gradient background)
```

## 🛑 Common Issues & Fixes
### ❌ 1. Template not found
panic: open templates/form.html: no such file or directory


✔️ Make sure you run the project from the root folder:
```
go run main.go
```

### ❌ 2. Port 8080 already in use

Find and kill the process:

Windows
```
netstat -ano | findstr :8080
taskkill /PID <PID> /F
```

Linux/Mac
```
lsof -i :8080
kill -9 <PID>
```

## 📚 References

Official Go Docs

- html/template Package Docs

Tutorials:

- “Go Web Servers in 20 Minutes” (YouTube)

- “Go Templates Crash Course”