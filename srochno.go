package srochno

// package main

import (
	"math/rand"
	"time"
)

// func main() {
// 	phr := NewPhrase()
// 	fmt.Println(phr)

// 	fmt.Scanln()
// }

func NewPhrase() string {

	var sentence string

	greeting := []string{"ВНИМАНИЕ", "СРОЧНО", "Всем привет", "Уважаемые родители!"}
	shitReceiver := []string{"мне", "соседу", "мужу на работе"}
	announceMethod := []string{"позвонил", "сказал", "написал"}
	avtoritetniyIstochnik := []string{"знакомый из министерства", "знакомый из очага эпидемии", "брат", "дядя", "родственник -оттуда-", "один дурак", "один человечек", "добрый человек"}
	message := []string{"и попросил", "и сказал", "и принудил"}
	message2 := []string{"никому не говорить!!!", "всем рассказать!!!"}
	message3 := []string{"на самом деле всё", "оказывается всё", "выяснилось, всё", "что НАС ОБМАНЫВАЮТ РЕПТИЛОИДЫ и будет"}
	status := []string{"лучше", "хуже", "совсем не так", "наоборот"}
	kogda := []string{"и скоро", " и прямо сейчас", "и уже завтра"}
	outcome := []string{"все мы умрем", "все будет хорошо", "прилетят инопланетяне", "царь россии выдаст всем премию", "все получат люлей"}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	newShit1 := r1.Intn(len(greeting))
	newShit2 := r1.Intn(len(shitReceiver))
	newShit3 := r1.Intn(len(announceMethod))
	newShit4 := r1.Intn(len(avtoritetniyIstochnik))
	newShit5 := r1.Intn(len(message))
	newShit6 := r1.Intn(len(message2))
	newShit7 := r1.Intn(len(message3))
	newShit8 := r1.Intn(len(status))
	newShit9 := r1.Intn(len(kogda))
	newShit10 := r1.Intn(len(outcome))

	sentence = greeting[newShit1] +
		" " + shitReceiver[newShit2] +
		" " + announceMethod[newShit3] +
		" " + avtoritetniyIstochnik[newShit4] +
		" " + message[newShit5] +
		" " + message2[newShit6] +
		" " + message3[newShit7] +
		" " + status[newShit8] +
		" " + kogda[newShit9] +
		" " + outcome[newShit10]

	return sentence
}
