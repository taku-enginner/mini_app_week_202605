package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Question struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Choices []string `json:"choices"`
	Answer string `json:"answer"`
	Explain string `json:"explain"`
}

var questions = []Question {
	{
		ID: 1,
		Text: "現在のディレクトリを表示するコマンドは？",
		Choices: []string{"ls", "pwd", "cd", "find"},
		Answer: "pwd",
		Explain: "pwd は print working directory の略です。",
	},
	{
		ID: 2,
		Text: "ファイル一覧を表示するコマンドは？",
		Choices: []string{"ls", "hoge", "mkdir", "find"},
		Answer: "ls",
		Explain: "ls はディレクトリ内のファイル一覧を表示します。",
	},
}

func main() {
	http.HandleFunc("/questions", getQuestions)

	log.Println("server start: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(questions)
}
