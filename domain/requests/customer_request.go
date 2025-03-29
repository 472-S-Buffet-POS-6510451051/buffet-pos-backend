package requests

type CustomerRegisterRequest struct {
	Phone	string 	`json:"phone" validate:"required,len=10"`
	PIN		string 	`json:"pin" validate:"required,len=6"`
}

type CustomerAddPointRequest struct {
	Phone	string 	`json:"phone" validate:"required,len=10"`
	PIN		string 	`json:"pin" validate:"required,len=6"`
	Point	int 	`json:"point" validate:"required"`
}

type CustomerRedeemRequest struct {
	Phone		string `json:"phone" validate:"required,len=10"`
	PIN			string `json:"pin" validate:"required,len=6"`
	InvoiceID 	string `json:"invoice_id" validate:"required"`
}