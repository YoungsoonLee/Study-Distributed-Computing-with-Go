package common

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Log is used for simple logging to console.
func Log(msg string) {
	log.Println("INFO - ", msg)
}

// Warn is used to log waring messages to console.
func Warn(msg string) {
	log.Println("--------------------------------------")
	log.Println(fmt.Sprintf("WARN: %s", msg))
	log.Println("--------------------------------------")
}

var punctuations = regexp.MustCompile(`^\p{P}+|\p{P}+$`)

// List of stop words that we want to ignore in our index.
var stopWords = []string{
	"a", "about", "above", "after", "again", "against", "all", "am", "an", "and", "any", "are", "aren't", "as", "at",
	"be", "because", "been", "before", "being", "below", "between", "both", "but", "by", "can't", "cannot", "could",
	"couldn't", "did", "didn't", "do", "does", "doesn't", "doing", "don't", "down", "during", "each", "few", "for",
	"from", "further", "had", "hadn't", "has", "hasn't", "have", "haven't", "having", "he", "he'd", "he'll", "he's",
	"her", "here", "here's", "hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm",
	"i've", "if", "in", "into", "is", "isn't", "it", "it's", "its", "itself", "let's", "me", "more", "most", "mustn't",
	"my", "myself", "no", "nor", "not", "of", "off", "on", "once", "only", "or", "other", "ought", "our", "ours",
	"ourselves", "out", "over", "own", "same", "shan't", "she", "she'd", "she'll", "she's", "should", "shouldn't",
	"so", "some", "such", "than", "that", "that's", "the", "their", "theirs", "them", "themselves", "then", "there",
	"there's", "these", "they", "they'd", "they'll", "they're", "they've", "this", "those", "through", "to", "too",
	"under", "until", "up", "very", "was", "wasn't", "we", "we'd", "we'll", "we're", "we've", "were", "weren't", "what",
	"what's", "when", "when's", "where", "where's", "which", "while", "who", "who's", "whom", "why", "why's", "with",
	"won't", "would", "wouldn't", "you", "you'd", "you'll", "you're", "you've", "your", "yours", "yourself", "yourselves"}

// simplifyToken is responsible to normalizing a string token and
// also checks whether the token should be indexed or not.
func SimplifyToken(token string) (string, bool) {
	simpleToken := strings.ToLower(punctuations.ReplaceAllString(token, ""))
	for _, stopWord := range stopWords {
		if stopWord == simpleToken {
			return "", false
		}
	}

	return simpleToken, true
}
