package purchase

import (

	//"read-pdf-go/controller"

	"testing"

	model "github.com/dougadriquei/desafioneoway/model/purchase"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Parallel()
		t.Run("Test1", create01)

	})
}

func create01(t *testing.T) {
	purchases := make([]model.Purchase, 0)
	purchase1 := model.Purchase{
		CpfCnpj:            "042.098.288-40",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}

	purchase2 := model.Purchase{
		CpfCnpj:            "042.098.288-41",
		Private:            "0",
		Incompleted:        "0",
		LastPurchaseDate:   "NULL",
		AverageTicket:      "NULL",
		LastPurchaseTicket: "NULL",
		MostFrequentStore:  "NULL",
		LastPurchaseStore:  "NULL",
	}
	purchases = append(purchases, purchase1, purchase2)
	result, err := CreatePurchases(purchases)

	assert.True(t, result == true)
	assert.Nil(t, err)
}
