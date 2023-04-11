package models

type Review struct {
	ID       int64
	PotionID int64
	Author   string
	Text     string
	Rating   int
}
