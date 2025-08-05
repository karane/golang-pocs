# Markdown to HTML Converter in Go

A simple command-line tool that converts basic Markdown files to HTML, **without using any external Markdown libraries**. It uses some common string manipulation functions.


## Features

- Converts basic Markdown syntax:
  - Headers (`#`, `##`)
  - Unordered lists (`- item`)
  - Bold (`**bold**`) and italic (`*italic*`) text (basic handling)
  - Paragraphs and line breaks
- Reads from a Markdown input file
- Writes output HTML to a file

---

## Getting Started

## Prerequisites
- Go 1.24.5
  


### Build and Run

```bash
go mod tidy

go run main.go sample.md output.html
```

