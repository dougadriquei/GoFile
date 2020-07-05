package purchase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePurchase(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Parallel()
		t.Run("Test1", create01)
		t.Run("Test3", create02InvalidCpf)
		t.Run("Test3", create03InvalidCnpj)
		t.Run("Test4", create04InvalidCnpjCpf)
	})
}

func create01(t *testing.T) {
	purchases := make([]Purchase, 0)
	purchase1 := Purchase{
		CpfCnpj:            "042.098.288-40",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}

	purchase2 := Purchase{
		CpfCnpj:            "74.570.773/0001-40",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}
	purchases = append(purchases, purchase1, purchase2)
	count, err := CreatePurchases(purchases)
	assert.True(t, count == 2)
	assert.Nil(t, err)
}

func create02InvalidCpf(t *testing.T) {
	purchases := make([]Purchase, 0)
	purchase1 := Purchase{
		CpfCnpj:            "042.098.288-41",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}

	purchase2 := Purchase{
		CpfCnpj:            "74.570.773/0001-40",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}
	purchases = append(purchases, purchase1, purchase2)
	count, err := CreatePurchases(purchases)
	assert.True(t, count == 1)
	assert.Nil(t, err)
}

func create03InvalidCnpj(t *testing.T) {
	purchases := make([]Purchase, 0)
	purchase1 := Purchase{
		CpfCnpj:            "042.098.288-40",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}

	purchase2 := Purchase{
		CpfCnpj:            "74.570.773/0001-41",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}
	purchases = append(purchases, purchase1, purchase2)
	count, err := CreatePurchases(purchases)
	assert.True(t, count == 1)
	assert.Nil(t, err)
}

func create04InvalidCnpjCpf(t *testing.T) {
	purchases := make([]Purchase, 0)
	purchase1 := Purchase{
		CpfCnpj:            "042.098.288-41",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}

	purchase2 := Purchase{
		CpfCnpj:            "74.570.773/0001-41",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}
	purchases = append(purchases, purchase1, purchase2)
	count, err := CreatePurchases(purchases)
	assert.True(t, count == 0)
	assert.Nil(t, err)
}
