package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AltaCvMod() error {
	// Declare an array of strings
	StringArray := []string{"pdfx", "biber", "bibhang", "biblabelsep", "pubtype", "bibsetup", "bibitemsep", "trimclip"}

	// Open the altacv.cls file
	file, err := os.OpenFile("altacv.cls", os.O_RDWR, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read the content of the file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Close the file and truncate it to clear its content
	file.Close()
	file, err = os.Create("altacv.cls")
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Remove lines containing the strings from the StringArray
	for _, line := range lines {
		keepLine := true
		for _, val := range StringArray {
			if strings.Contains(line, val) {
				keepLine = false
				break
			}
		}
		if keepLine {
			// Replace pdfstringdef method to escape
			newLine := strings.Replace(line, "pdfstringdef", "escape", -1)
			fmt.Fprintln(file, newLine)
		}
	}

	// Add modifications to altacv.cls
	modifications := `
% Modification
\newcommand{\cvproject}[3]{%
  {\large\color{emphasis}#1\par}
  \smallskip\normalsize
  \ifstrequal{#2}{}{}{
  \textbf{\color{accent}#2}\par
  \smallskip}
  \ifstrequal{#3}{}{}{{\small\makebox[0.5\linewidth][l]{\faCalendar~#3}}}%
  \medskip\normalsize
}
\newcommand{\cvskillstr}[2]{%
  \textcolor{emphasis}{\textbf{#1}}\hfill
  \textbf{\color{body}#2}\par
}
`

	fmt.Fprintln(file, modifications)

	log.Println("Refactor complete.")
	return nil
}
