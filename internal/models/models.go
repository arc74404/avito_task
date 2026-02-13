package models

type SegmentModel struct {
	Slug string `json:"slug"`
}

type UserSegments struct {
	Segments []int `json:"segments"`
}
