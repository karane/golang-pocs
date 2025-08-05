package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func markdownToHTML(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var html strings.Builder

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "# ") {

			html.WriteString("<h1>" + strings.TrimPrefix(trimmed, "# ") + "</h1>\n")

		} else if strings.HasPrefix(trimmed, "## ") {

			html.WriteString("<h2>" + strings.TrimPrefix(trimmed, "## ") + "</h2>\n")

		} else if strings.HasPrefix(trimmed, "- ") {

			html.WriteString("<ul>\n")
			for _, li := range lines {
				if strings.HasPrefix(strings.TrimSpace(li), "- ") {
					html.WriteString("<li>" + strings.TrimPrefix(strings.TrimSpace(li), "- ") + "</li>\n")
				}
			}
			html.WriteString("</ul>\n")
			break

		} else if trimmed == "" {

			html.WriteString("<br/>\n")

		} else {

			// Bold and italic parsing
			text := strings.ReplaceAll(trimmed, "**", "<b>")
			text = strings.ReplaceAll(text, "*", "<i>")

			// Replacing the closing tags correctly
			text = strings.Replace(text, "<b>", "</b>", 1)
			text = strings.Replace(text, "<i>", "</i>", 1)

			html.WriteString("<p>" + text + "</p>\n")
		}
	}

	return html.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: markdown2html <input.md> [output.html]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := "output.html"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var markdown strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		markdown.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	html := markdownToHTML(markdown.String())

	err = os.WriteFile(outputFile, []byte(html), 0644)
	if err != nil {
		fmt.Println("Failed to write HTML:", err)
		os.Exit(1)
	}

	fmt.Println("✅ HTML file created:", outputFile)
}
