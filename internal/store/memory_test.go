package store

import( "testing"
"github.com/mrckurz/CI-CD-MCM/internal/model")

func TestCreateAndGet(t *testing.T) {
	_ = NewMemoryStore()

	tests := []struct {
		name    string
		input   model.Product
	}{
		{name: "basic product", input: model.Product{Name: "Apple", Price: 1.99}},
		{name: "second product", input: model.Product{Name: "Banana", Price: 0.49}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewMemoryStore()

			created := s.Create(tt.input)   // ID wird automatisch vergeben!

			got, err := s.GetByID(created.ID)
			if err != nil {
				t.Fatalf("GetByID() failed: %v", err)
			}
			if got.ID != created.ID || got.Name != tt.input.Name {
				t.Errorf("got %+v, want %+v", got, created)
			}
		})
	}
}

func TestGetAllEmpty(t *testing.T) {
	s := NewMemoryStore()
	products := s.GetAll()
	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestDeleteNonExistent(t *testing.T) {
	s := NewMemoryStore()
	err := s.Delete(999)
	if err != ErrNotFound {
		t.Error("expected ErrNotFound when deleting non-existent product")
	}
}


func TestUpdateProduct(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{Name: "Apple", Price: 1.99})

	updated, err := s.Update(created.ID, model.Product{Name: "Green Apple", Price: 2.49})
	if err != nil {
		t.Fatalf("Update() failed: %v", err)
	}

	got, err := s.GetByID(created.ID)
	if err != nil {
		t.Fatalf("GetByID() after Update() failed: %v", err)
	}
	if got.Name != "Green Apple" || got.Price != updated.Price {
		t.Errorf("update not applied: got %+v", got)
	}
}

func TestDeleteProduct(t *testing.T) {
	s := NewMemoryStore()

	created := s.Create(model.Product{Name: "Apple", Price: 1.99})

	if err := s.Delete(created.ID); err != nil {
		t.Fatalf("Delete() failed: %v", err)
	}

	_, err := s.GetByID(created.ID)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound after Delete(), got %v", err)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	s := NewMemoryStore()

	_, err := s.GetByID(999)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
// TODO: Add tests for Update, Delete of existing product, and GetByID with invalid ID
