package domain

/*********************** LOAN **************************/
type Loan struct {
	ID      int         `json:"id"`
	IDBook  string      `json:"idBook"`
	IDUser  int         `json:"idUser"`
	DueDate string      `json:"dueDate"`
	Info    Information `json:"bookInfo"`
}

func (l Loan) IDValid() bool {
	return l.ID > 0
}

func (l Loan) HasIDBook() bool {
	return l.IDBook != ""
}

func (l Loan) HasIDUser() bool {
	return l.IDUser > 0
}

func (l Loan) HasDueDate() bool {
	return l.DueDate != ""
}

/*********************** USER **************************/
type User struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func (u User) HasName() bool {
	return u.Nombre != ""
}

func (u User) HasSurname() bool {
	return u.Apellido != ""
}

func (u User) IDValid() bool {
	return u.ID > 0
}

/*********************** BOOK **************************/
const (
	specialChar = ".-,/|'#&%?¿`!¡;:[]{}+*º<>$="
)

type Book struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func (b Book) HasTitle() bool {
	return b.Titulo != ""
}

func (b Book) HasAuthor() bool {
	return b.Autor != ""
}

func (b Book) SpecialChar() bool {
	for _, cw := range b.Titulo {
		for _, cc := range specialChar {
			if cw == cc {
				return true
			}
		}
	}

	for _, cw := range b.Autor {
		for _, cc := range specialChar {
			if cw == cc {
				return true
			}
		}
	}

	return false
}

/***************** API GOOGLE BOOKS ********************/
type GoogleBooks struct {
	Items []Items `json:"items"`
}

type Items struct {
	Id   string      `json:"id"`
	Info Information `json:"volumeInfo"`
}

type Information struct {
	Titulo           string    `json:"title"`
	Subtitulo        string    `json:"subtitle"`
	Autores          [5]string `json:"authors"`
	FechaPublicacion string    `json:"publishedDate"`
}
