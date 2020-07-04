package purchase

// Purchase Modelo da Compra
type Purchase struct {
	CpfCnpj            string `json:"cpf_cnpj,omitempty"`
	Private            string `json:"private,omitempty"`
	Incompleted        string `json:"incompleted,omitempty"`
	LastPurchaseDate   string `json:"last_purchase_date,omitempty"`
	AverageTicket      string `json:"average_ticket,omitempty"`
	LastPurchaseTicket string `json:"last_purchase_ticket,omitempty"`
	MostFrequentStore  string `json:"most_frequent_store,omitempty"`
	LastPurchaseStore  string `json:"last_purchase_store,omitempty"`
}
