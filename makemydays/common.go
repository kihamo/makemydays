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
	Id            int64  `db:"id" json:"-"`
	OriginalValue string `db:"original_value" json:"-"`
	Title         string `db:"title" json:"title"`
	TitleRus      string `db:"title" json:"title_rus"`
	Year          int64  `db:"year" json:"year"`
}

func (m Movie) String() string {
	return m.OriginalValue
}

type Song struct {
	Id            int64  `db:"id" json:"-"`
	OriginalValue string `db:"original_value" json:"-"`
	Title         string `db:"title" json:"title"`
	Author        string `db:"author" json:"author"`
}

func (s Song) String() string {
	return s.OriginalValue
}

type Book struct {
	Id            int64  `db:"id" json:"-"`
	OriginalValue string `db:"original_value" json:"-"`
	Title         string `db:"title" json:"title"`
	Author        string `db:"author" json:"author"`
}

func (b Book) String() string {
	return b.OriginalValue
}

type Word struct {
	Id   int64  `db:"id" json:"-"`
	Word string `db:"title" json:"word"`
}

func (w Word) String() string {
	return w.Word
}

type Task struct {
	Id    int64  `db:"id" json:"-"`
	Title string `db:"title" json:"title"`
}

func (t Task) String() string {
	return t.Title
}

type Food struct {
	Id    int64  `db:"id" json:"-"`
	Title string `db:"title" json:"title"`
}

func (f Food) String() string {
	return f.Title
}
