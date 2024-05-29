package main

type ItemType int

const (
	Furniture = iota
	Clothes
	Technology
	Utility
	Books
	Notes
	Junk
	Valuables
	Music
)

var ItemTypeName = map[ItemType]string{
	Furniture:  "Furniture",
	Clothes:    "Clothes",
	Technology: "Technology",
	Utility:    "Utility",
	Books:      "Books",
	Notes:      "Notes",
	Junk:       "Junk",
	Valuables:  "Valuables",
	Music:      "Music",
}

func (i ItemType) String() string {
	return ItemTypeName[i]
}
