package course

type Course struct {
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Url         string `bson:"url,omitempty"`
	Size        string `bson:"size,omitempty"`
	Duration    string `bson:"duration,omitempty"`
	Notes       string `bson:"notes,omitempty"`
}
