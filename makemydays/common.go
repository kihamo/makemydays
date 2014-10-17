package makemydays

type Recommendation struct {
	Movie Movie `json:"movie"`
	Song  Song `json:"song"`
	Word  Word `json:"word"`
	Book  Book `json:"book"`
	Task  Task `json:"task"`
	Food  Food `json:"food"`
}

type Movie struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Year  int64  `db:"year" json:"year"`
}

func (r *Movie) GetTitle() string {
	return r.Title
}

type Song struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type Word struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type Book struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type Task struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type Food struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}
