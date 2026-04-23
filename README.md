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

go-url-shortener/
│── main.go
│── go.mod

---

## ⚙️ Installation & Setup

### 1. Clone the Repository

git clone https://github.com/your-username/go-url-shortener.git  
cd go-url-shortener

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
