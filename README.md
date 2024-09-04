<p align="center">
  <picture>
    <source srcset="./assets/qonto-logo.png">
    <img alt="qonto-logo logo" height="180px" src="./assets/qonto-logo.png">
  </picture>
</p>

<p align="center">
  <strong style="font-size: 24px;">Final Submission for Backend Engineer Intern Technical Test</strong>
  <br />
  <em>By Aymene Bousbia, currently working at Veepee</em>
  <br />
  <br />
</p>

<div align="center">
  <p>
    <a href="https://github.com/aymene01/qonto-technical-test/issues"><img src="https://img.shields.io/github/issues/aymene01/qonto-technical-test.svg" alt="GitHub Issues"></a>
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/built%20with-Go-00ADD8.svg" alt="Built with Go"></a>
  </p>
</div>

# Qonto Technical Test - Backend Engineer Intern

## Introduction

Hi Qonto Team,

My name is Aymene, and I am currently working at Veepee. I am excited to submit my final test for the Backend Engineer Intern position at Qonto. This project was both challenging and rewarding, and I’m pleased to share the results with you.

I’ve implemented a series of backend functions and exposed them via RESTful endpoints. The application has been deployed and is accessible at [https://qonto-technical-test.fly.dev](https://qonto-technical-test.fly.dev). Below, you'll find detailed information about the implementation, including the exposed endpoints, deployment details, and testing.

## Project Overview

This project includes:

- **Dockerized Environment:** The application runs in a containerized environment using Docker.
- **Unit Testing:** Comprehensive unit tests to ensure the correctness of each function.
- **CI/CD Pipeline:** Continuous Integration and Continuous Deployment are handled via GitHub Actions.
- **Monitoring:** A Grafana dashboard is set up for monitoring the application.

## Endpoints

### 1. Health Check

**Endpoint:** `GET /health`

**Description:** Checks the health status of the service.

**Response:**

```json
{
  "status": "OK"
}
```

**exemple:**

```bash
curl -X GET https://qonto-technical-test.fly.dev/health
```

## 2. Functions

The application exposes the following functions:

### Count Frequency

**Endpoint:** `POST /count-frequency`

**Description:** Counts how many times each item occurs in a list.
**Request:**

```json
{
  "data": ["apple", "banana", "apple", "banana", "orange", "apple"]
}
```

**Response:**

```json
{
  "apple": 3,
  "banana": 2,
  "orange": 1
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/count-frequency -H "Content-Type: application/json" -d '["egg", "egg", "soap", "soap", "soap"]'
```

### Is Truthy

**Endpoint:** `POST /is-truthy`

**Description:** Determines if any of its arguments is true

**Request:**

```json
{
  "data": [true, false, false]
}
```

**Response:**

```json
{
  "result": true
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/is-truthy -H "Content-Type: application/json" -d '[true, false]'
```

### Is Student Of

**Endpoint:** `POST /is-student-of`

**Description:** Returns true if a student attended a specified college.

**Request:**

```json
{
  "obj": {
    "college": "cambridge",
    "year": 2008
  },
  "college": "cambridge"
}
```

**Response:**

```json
{
  "result": true
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/is-student-of -H "Content-Type: application/json" -d '{"obj": {"college": "cambridge", "year": 2008}, "college": "cambridge"}'
```

### Generate Even Int List

**Endpoint:** `POST /generate-even-int-list`

**Description:** Generates a list of even integers from 0 to n inclusive.

**Request:**

```json
{
  "n": 10
}
```

**Response:**

```json
{
  "result": [0, 2, 4, 6, 8, 10]
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/generate-even-int-list -H "Content-Type: application/json" -d '{"n": 7}'
```

### Parse Int

**Endpoint:** `POST /parse-int`

**Description:** Parses an integer from a string.

**Request:**

```json
{
  "value": "10"
}
```

**Response:**

```json
{
  "result": 10
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/parse-int -H "Content-Type: application/json" -d '{"value": "10"}'
```

## Was Student During

**Endpoint:** `POST /was-student-during`

**Description:** Determines if a student was attending college during a specific year.

**Request:**

```json
{
  "student": {
    "college": "cambridge",
    "year": 2004
  },
  "college": "cambridge",
  "startYear": 2002,
  "endYear": 2006
}
```

**Response:**

```json
{
  "result": true
}
```

**curl:**

```bash
curl -X POST https://qonto-technical-test.fly.dev/was-student-during -H "Content-Type: application/json" -d '{"student": {"college": "cambridge", "year": 2004}, "college": "cambridge", "startYear": 2002, "endYear": 2006}'
```

## Additional Documentation

For a detailed breakdown of the DoSomething function, including its operations and how it can be renamed to better reflect its purpose, please refer to link [DoSomething](./DoSomething.md)

## 👤 Author

**Aymene Bousbia**

- 🔍 Explore: [GitHub Profile](https://github.com/aymene01)
- 💬 Ask me about anything [here](https://github.com/aymene01/nestjs-final-test/issues).

# 📬 Feedback

Your opinions and feedback are what shape the future; let's craft it beautifully together. Share your thoughts in the issues or through discourse.

<!-- Your personal message or trademark --> <div align="center"> <sub>Built with ❤️ by <a href="https://github.com/aymene01">Aymene Bousbia</a></sub> </div>
