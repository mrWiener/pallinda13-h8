
package Lab8

import (
  "bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func Main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go parseQuestions(questions, answers)
	go prophecy(answers)
	go printAnswers(answers)

	return questions
}

func printAnswers(answers <- chan string) {
	for q := range answers {
		fmt.Println(q)
		fmt.Print("> ")
	}
}

func parseQuestions(questions <- chan string, answers chan<- string) {
	answersStrings := []string{
		"I dont know.",
		"Ask me something else.",
		"Just go for it.",
	}

	for q := range questions {
		go answerQuestion(q, answers, answersStrings)
	}
}

func answerQuestion(question string, answers chan<- string, answersStrings[] string){
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	a := "Your question was: " +question + " | Pynthia answers: " + answersStrings[rand.Intn(len(answersStrings))]
	answers <- a
}

func prophecy(answers chan<- string) {

	r := []string{
		"The moon is dark.",
		"The sun is bright.",
		"Mmmyes.",
		"Dont ask me anything.",
		"Orly?."
	}

for{
	time.Sleep(time.Duration(5+rand.Intn(10)) * time.Second)
	answers <- "... " + r[rand.Intn(len(r))]
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
