package entities

// Balance is the calc of all transaction types
type Balance struct {
	Income  uint64 `json:"income"`
	Outcome uint64 `json:"outcome"`
	Total   uint64 `json:"total"`
}
