package handler

type Handler struct {
	UserHandler       *UserHandler
	BillKindHandler   *BillKindHandler
	BillRecordHandler *BillRecordHandler
	TokenHandler      *TokenHandler
	OtherHandler      *OtherHandler
}
