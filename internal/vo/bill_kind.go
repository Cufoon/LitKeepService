package vo

type BillKind struct {
	KindID      string `json:"kindID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BillKindCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UpKind      string `json:"upKind"`
}

type BillKindCreateRes struct {
	Created bool `json:"created"`
}

type BillKindQueryTreeResData struct {
	BillKind
	Children []BillKind `json:"children"`
}

type BillKindQueryTreeRes struct {
	Kind []BillKindQueryTreeResData `json:"kind"`
}

type BillKindModifyReq struct {
	KindID      string `json:"kindID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UpKind      string `json:"upKind"`
}

type BillKindModifyRes struct {
	Modified bool `json:"modified"`
}

type BillKindDeleteReq struct {
	KindID string `json:"kindID"`
}

type BillKindDeleteRes struct {
	Deleted bool `json:"deleted"`
}
