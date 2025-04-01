# PDFiber

**Status:** Under Development 🚧

## About

An API written in **Go** and built using the **Fiber framework**, inspired by tools like **iLovePDF**. The primary goal of this project is to serve as a way to enhance my skills in Go programming. It's not intended to be overly complex, but rather a practical and hands-on approach to learning and improving my development abilities while working with PDF file manipulation tasks.

Even though it’s a learning project, PDFiber aims to provide useful features for handling PDF files with a focus on simplicity and efficiency.

## Features

- ✅ Merge multiple PDF files
- ✅ Split PDF files into smaller parts
- 🚧 Generate new PDF files
- 🚧 Extract text from PDF files

## Requirements

Before starting, ensure you have:

- **Go** (version 1.16 or later)
- A package manager like **Go Modules**
- [Fiber Framework](https://gofiber.io/)

## Installation

1. Clone this repository:

```
   git clone https://github.com/pentodo/pdfiber.git
```

2. Navigate to the project directory:

```
   cd pdfiber
```

3. Install the dependencies:

```
   go mod tidy
```

## Usage

Start the server:

```
go run main.go
```

The API will be available at `http://localhost:3000`.

## Contributing

Pull requests are welcome! For major changes, open an **issue** first to discuss what you would like to change. Here are some impactful areas where you can make a difference:

- **Documentation:**  
  Build robust API documentation using tools like **Swagger** or **Redoc** to provide a clear understanding of endpoints, parameters, and responses.

- **Request Validation:**  
  Strengthen server-side validation to ensure data inputs are secure and reliable. Libraries like **validator/v10** or custom validation logic can be implemented.

- **User Authentication:**  
  Add an authentication system to restrict access to premium features, utilizing technologies such as **JWT (JSON Web Token)** or **OAuth**.

- **New Functionalities:**  
  Develop additional features, like PDF compression or file upload support directly from the browser.

- **Optimizations:**  
  Enhance the handling of large PDF files and explore memory optimization techniques during complex operations.
