package populating

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func TestJsonService_Populate(t *testing.T) {
	populator := jsonService{fileName: "test.json", manager: managing.MockService{}}
	file, err := os.Open(populator.fileName)
	if err != nil {
		t.Fatal(err)
	}
	populator.file = file
	defer populator.file.Close()

	err = populator.Populate()
	assert.NoError(t, err)
}

func BenchmarkJsonService_Populate(b *testing.B) {
	populator := jsonService{fileName: "ports.json", manager: managing.MockService{}}
	file, err := os.Open(populator.fileName)
	if err != nil {
		b.Fatal(err)
	}
	populator.file = file
	defer populator.file.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		populator.Populate()
	}
}