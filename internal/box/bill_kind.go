package box

type BillKind struct {
	KindID      string
	Name        string
	Description string
	UpKind      string
	OverKind    string
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
