package views

import "commitinder/models"

type DuelResponse struct {
	DuelId         int    `json:"duel_id"`
	DiffContent    string `json:"diff_content"`
	CommitMessageA string `json:"commit_message_a"`
	CommitMessageB string `json:"commit_message_b"`
}

func (d *DuelResponse) FromModelToView(content string, duel *models.Duel) {
	d.DuelId = duel.Id
	d.DiffContent = content
	d.CommitMessageA = duel.CommitMessageA.Message
	d.CommitMessageB = duel.CommitMessageB.Message
}

type Options struct {
	Aspect       string  `gorm:"aspect"`
	ChoiceTime   float64 `json:"choice_time"`
	ChosenOption string  `json:"chosen_option"`
}

type ResultsRequest struct {
	DuelId  int       `json:"duel_id"`
	Options []Options `json:"options"`
}
