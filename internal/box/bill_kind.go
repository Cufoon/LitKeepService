package box

type BillKind struct {
	KindID      string `json:"kindID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UpKind      string `json:"upKind"`
	OverKind    string `json:"overKind"`
}

type BillKindTree struct {
	BillKind
	Children []BillKind
}

type BillKindCreate struct {
	UserID      string
	Name        string
	Description string
	UpKind      string
}
