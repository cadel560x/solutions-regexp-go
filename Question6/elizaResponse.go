package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func ElizaResponse(input string) string {

	// e.g. My father is a doctor.
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don't you tell me more about your father?"
	}

	// e.g. I am sad.
	re := regexp.MustCompile(`(?i).*\bI'?\s*a?m \b([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		subMatch := re.ReplaceAllString(input, "$1?")
		reflectedString := Reflect(subMatch)
		return "How do you know you are " + reflectedString
	}

	// e.g. I have lots of friends
	re = regexp.MustCompile(`(?i)^\s*I have ([^.!?]*)[.!?\s]*$`)
	if matched := re.MatchString(input); matched {
		subMatch := re.ReplaceAllString(input, "$1?")
		reflectedString := Reflect(subMatch)
		return "Why do you tell me that you've " + reflectedString
	}

	// e.g. I don't care
	re = regexp.MustCompile(`(?i)^\s*I don't ([^.!?]*)[.!?\s]*$`)
	if matched := re.MatchString(input); matched {
		subMatch := re.ReplaceAllString(input, "$1?")
		reflectedString := Reflect(subMatch)
		return "Don't you really " + reflectedString
	}

	// e.g. I feel strange
	re = regexp.MustCompile(`(?i)^\s*I feel ([^.!?]*)[.!?\s]*$`)
	if matched := re.MatchString(input); matched {
		subMatch := re.ReplaceAllString(input, "$1")
		reflectedString := Reflect(subMatch)
		return "When you feel " + reflectedString + ", what do you do?"
	}

	// e.g. Anything else.
	responses := []string{
		"I'm not sure what you're trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}

	randindex := rand.Intn(len(responses))

	return responses[randindex]
}

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\s`)
	tokens := boundaries.Split(input, -1)

	// log.Println("DEBUG: Reflect: tokens: ", tokens)

	// Some key prepositions
	prepositions := []string{
		"to",
		"by",
		"under",
		"about",
		"on",
		"according",
		"over",
		"of",
		"without",
	}

	// List the reflections.
	reflections := [][]string{
		{`was`, `were`},
		{`I`, `you`},
		{`I'm`, `you are`},
		{`I'd`, `you would`},
		{`I've`, `you have`},
		{`I'll`, `you will`},
		{`my`, `your`},
		{`you're`, `I am`},
		{`were`, `was`},
		{`you've`, `I have`},
		{`you'll`, `I will`},
		{`your`, `my`},
		{`yours`, `mine`},
		// {`you`, `me`},
		{`me`, `you`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {

			// Let's take 'you' separately because it is the same as subject pronoun as object pronoun
			if token == "you" {

				// Loop through the prepositions
				for j, preposition := range prepositions {
					// Compare the previous word, that is 'token[i-2]' to the 'preposition'. 'token[i-1]' is the space character.
					if tokens[i-2] == preposition {
						// If 'you' is an object pronoun to a preposition, the swap it for 'me'
						tokens[i] = "me"
						break
					}

					// The previous word was not a preposition, swapping for a subject pronoun
					if j == len(prepositions)-1 {
						tokens[i] = "I"
					}

				} // for j, prepostition
				// As for the rest of reflections, keep doing the normal substitution
			} else if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			} // if - else if

		} // for 'reflection'
	} // for 'i'

	// Put the tokens back together.
	// log.Println("DEBUG: Reflect: tokens: ", tokens)
	return strings.Join(tokens, ` `)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father's favourite.")
	fmt.Println(ElizaResponse("I was my father's favourite."))
	fmt.Println()

	fmt.Println("I'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()

	// I am and variations: I'm Im i'm im
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("I'm supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("Im supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("i am supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("i'm supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("im supposed to just take what you're saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you're saying at face value?"))
	fmt.Println()

	fmt.Println("I have eaten at an italian restaurant.")
	fmt.Println(ElizaResponse("I have eaten at an italian restaurant"))
	fmt.Println()

	fmt.Println("I don't know logarithms.")
	fmt.Println(ElizaResponse("I don't know logarithms"))
	fmt.Println()

	fmt.Println("I feel nice.")
	fmt.Println(ElizaResponse("I feel nice"))
	fmt.Println()

}
