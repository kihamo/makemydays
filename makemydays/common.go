package makemydays

type Recommendation struct {
	Movie *Movie `json:"movie"`
	Song  *Song `json:"song"`
	Word  *Word `json:"word"`
	Book  *Book `json:"book"`
	Task  *Task `json:"task"`
	Food  *Food `json:"food"`
}

type Movie struct {
	Id       int64  `db:"id" json:"-"`
	Title    string `db:"title" json:"title"`
	TitleRus string `db:"title" json:"title_rus"`
	Year     int64  `db:"year" json:"year"`
}

type Song struct {
	Id     int64  `db:"id" json:"-"`
	Title  string `db:"title" json:"title"`
	Author string `db:"author" json:"author"`
}

type Book struct {
	Id     int64  `db:"id" json:"-"`
	Title  string `db:"title" json:"title"`
	Author string `db:"author" json:"author"`
}

type Word struct {
	Id    int64  `db:"id" json:"-"`
	Title string `db:"title" json:"title"`
}

type Task struct {
	Id    int64  `db:"id" json:"-"`
	Title string `db:"title" json:"title"`
}

type Food struct {
	Id    int64  `db:"id" json:"-"`
	Title string `db:"title" json:"title"`
}
