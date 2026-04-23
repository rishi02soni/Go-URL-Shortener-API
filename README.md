# 🔗 Go URL Shortener API

A simple and fast URL Shortener built with Go, similar to Bitly.  
This project allows users to convert long URLs into short, shareable links and redirect them instantly.

---

## 🚀 Features

- 🔗 Shorten long URLs
- ⚡ Fast redirection
- 🧠 In-memory storage (map-based)
- 🪶 Lightweight and beginner-friendly
- 🌐 REST API based

---

## 🛠️ Tech Stack

- Go (Golang)
- Gorilla Mux (Router)
- net/http

---

## 📁 Project Structure
```
go-url-shortener/
│── main.go
│── go.mod
```
---

## ⚙️ Installation & Setup

### 1. Clone the Repository
```
git clone https://github.com/rishi02soni/go-url-shortener.git  
cd go-url-shortener
```
---

### 2. Install Dependencies

go mod tidy

---

### 3. Run the Server

go run main.go

---

## 🌐 Server

http://localhost:8000

---

## 📡 API Endpoints

### ➤ Shorten URL

POST /shorten

#### Request Body

```
{
  "original": "https://example.com"
}
```
## Response
```
{
  "original": "https://example.com",
  "short": "http://localhost:8000/abc123"
}
```
## ➤ Redirect to Original URL
```
GET /{short}

Example:

http://localhost:8000/abc123

➡️ Redirects to original URL
```

## Using Postman
```
Method: POST
URL: http://localhost:8000/shorten
Body: Raw JSON
```

## Limitations
```
Data is stored in memory (will reset on restart)
No authentication
No custom aliases
```
## Author
Rishi Soni
Passionate about building impactful tech

