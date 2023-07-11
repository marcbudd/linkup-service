package services

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/marcbudd/linkup-service/linkuperrors"
)

type gptAnswer struct {
	ID      string `json:"id"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func CreatePostGPT() (string, *linkuperrors.LinkupError) {

	url := os.Getenv("OPENAI_API_URL")
	apiKey := os.Getenv("OPENAI_API_KEY")
	payload := `{
		"prompt": "Generiere mir einen kurzen, lustigen Tweet mit max. 150 Zeichen.",
		"max_tokens": 50,
		"temperature": 0.7,
		"n": 5
	}`

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return "", linkuperrors.New("openai api error: "+err.Error(), http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Body = ioutil.NopCloser(strings.NewReader(payload))

	resp, err := client.Do(req)
	resp.StatusCode = 400 // is used to block the api call for now because of api requests cost
	if err != nil || resp.StatusCode != 200 {
		// return "", linkuperrors.New("failed to create post from api", http.StatusInternalServerError)
		return FakeCreatePostGPT(), nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	string := string(body)

	var response gptAnswer
	err = json.Unmarshal([]byte(string), &response)
	if err != nil {
		return "", linkuperrors.New(err.Error(), http.StatusInternalServerError)
	}

	if len(response.Choices) > 0 {
		text := response.Choices[0].Text
		return text, nil
	}

	return "", linkuperrors.New(err.Error(), http.StatusInternalServerError)

}

// Fake API call because if free API call limit is reached
func FakeCreatePostGPT() string {
	posts := []string{
		"Wenn ich ein Vogel wäre, würde ich den ganzen Tag vor Fenstern sitzen und die Menschen erschrecken. Sorry, aber manchmal muss man seine Flügel ausbreiten.",
		"Mein Vorsatz für dieses Jahr ist es, weniger Prokrastination zu betreiben. Morgen fange ich damit an.",
		"Bin gerade im Fitnessstudio, aber es gibt hier so viele Spiegel, dass ich den Eindruck habe, ich bin in einer Selfie-Convention gelandet.",
		"Ich habe beschlossen, Yoga zu machen, um meine Flexibilität zu verbessern. Aktueller Stand: Ich kann meine Pizza im Sitzen essen.",
		"Die meisten Menschen haben Angst vor Geistern oder Monstern. Ich hingegen habe Angst vor meiner eigenen Bankkontonummer.",
		"Ich habe gerade versucht, mit meinem Kaffee zu sprechen. Aber er antwortete nicht. Wahrscheinlich war er noch zu heiß, um zu reden.",
		"Ich habe meine Fitness-App gestartet und sie sagte mir, ich solle mehr laufen. Also bin ich jetzt hier, um meinen Laptop im Park spazieren zu führen.",
		"Ich habe gerade einen Lauf gemacht. Naja, eigentlich waren es mehrere Läufe zur Kaffeemaschine und zurück.",
		"Ich habe einen langen Spaziergang gemacht, um meine Gedanken zu klären. Jetzt weiß ich, dass ich hungrig, aber immer noch verwirrt bin.",
		"Ich habe eine Liste mit Dingen, die ich heute erreichen will. Aber das Sofa und ich sind uns einig, dass wir die Liste auf morgen verschieben.",
		"Ich habe versucht, meinen Computer zu motivieren, schneller zu arbeiten. Also habe ich ihm ein Foto von einem neuen Laptop gezeigt. Jetzt weint er.",
		"Ich versuche, meine Karriereleiter zu erklimmen, aber irgendwie scheint sie auf einer felsigen Oberfläche zu stehen. Ich glaube, ich habe die Leiter zum Erfolg verpasst und bin auf der Rutschbahn des Sarkasmus gelandet.",
		"Ich habe heute Morgen ein Buch über Geduld gelesen. Ich konnte es kaum bis zur Hälfte aushalten.",
		"Es gibt zwei Arten von Menschen auf der Welt: diejenigen, die Mathe verstehen, und diejenigen, die sich einfach gut in die Wissenschaft des Stirnrunzelns einfinden können.",
		"Mein Arzt hat mir empfohlen, mehr Vitamin C zu mir zu nehmen. Jetzt warte ich nur noch darauf, dass die Chips zu Orangen werden.",
		"Die gute Nachricht: Mein Leben ist ein Zirkus. Die schlechte Nachricht: Ich bin der Clown.",
		"Das Leben ist wie eine Schachtel Pralinen. Du weißt nie, welche Geschmacksrichtung dich enttäuschen wird.",
		"Ich habe gelernt, dass manche Menschen nur dann lächeln, wenn sie dich brauchen. Ansonsten haben sie ein Gesicht wie ein bedrucktes Kissen.",
		"Das Universum hat einen seltsamen Sinn für Humor. Es stellt sicher, dass die Batterie deines Rauchmelders immer nachts leer wird.",
		"Ich habe heute meinen Wecker umarmt, um ihm zu zeigen, wie sehr ich ihn verachte. Jetzt liegt er immer noch am Boden und weint.",
		"Mein Spiegel hat mir gesagt, dass ich schön bin. Ich denke, er hat eine Schwäche für Lügen.",
		"Ich habe beschlossen, meine Träume zu verfolgen. Aber leider habe ich sie verloren, weil sie viel zu schnell für mich gerannt sind.",
		"Manche Menschen sind wie Wolken. Wenn sie verschwinden, wird der Tag schöner.",
		"Das Leben ist wie ein Spiel. Die Regeln sind kompliziert.",
		"Mein Gehirn ist wie ein Internetbrowser. Es hat 19 Tabs geöffnet, spielt Musik ab und stürzt alle fünf Minuten ab.",
		"Ich habe festgestellt, dass ich mein Passwort so oft ändere, dass ich meine eigene Identität infrage stelle.",
		"Meine Kleidung schrumpft. Das muss am Trockner liegen, definitiv nicht an den vielen Pizzen, die ich esse.",
		"Mein Kalender schrie mich an und sagte: 'Du hast zu viele Termine!' Ich schrie zurück: 'Du hast zu wenig Feiertage!'",
		"Ich habe versucht, mich selbst zu googeln, aber Google hat mich gefragt: 'Meinst du nicht, dass du genug Zeit mit dir selbst verbringst?'",
		"Ich habe versucht, mich mit meiner Kaffeetasse anzufreunden. Jetzt treffen wir uns regelmäßig, um über den Tag zu plaudern.",
		"Meine Mutter hat mir gesagt, ich solle immer meine Träume verfolgen. Deshalb schlafe ich jetzt den ganzen Tag.",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(posts))
	return posts[index]
}
