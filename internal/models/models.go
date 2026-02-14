package models

type SegmentModel struct {
	Slug string `json:"slug"`
}

type UserSegments struct {
	Segments []int `json:"segments"`
}

type AddRemoveUserUpdJsonData struct {
	Add    []int
	Remove []int
}
