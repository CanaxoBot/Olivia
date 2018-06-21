package language

import (
	"strings"
	"regexp"
)

type Word struct {
	Content string
}

const vowels = "[aeiouyâàëéêèïîôûù]"

// Marking specific vowels as consonants by puting them in upper case
func (word Word) MarkVowels() (response Word) {
	// Initialize the possibilities with regex (as a string) and a replace func
	matches := make(map[string]func (string) string)

	// Match `u` and `i` preceded and followed by vowels
	matches[vowels + "[ui]" + vowels] = func (match string) string {
		matchBytes := []byte(match)
		matchBytes[1] = []byte(strings.ToUpper(string(matchBytes[1])))[0]
		return string(matchBytes)
	}

	// Match a `y` preceded or followed by a vowel
	matches["(" + vowels + "y)|(y" + vowels + ")"] = func (match string) string {
		return strings.Replace(string(match), "y", "Y", 1)
	}

	// Iterate the matches and execute the replacements
	content := strings.ToLower(word.Content)
	for pattern, replace := range matches {
		regex := regexp.MustCompile(pattern)
		content = regex.ReplaceAllStringFunc(content, replace)
	}

	return Word{strings.Replace(content, "qu", "qU", -1)}
}