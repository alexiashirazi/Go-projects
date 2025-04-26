// package store

// import "curs1_boilerplate/cmd/exemplu_store/model"

// type memoryProductStore struct {
// 	products []model.Product
// }

// // Add implements ProductStore.
// func (s *memoryProductStore) Add(p model.Product) error {
// 	s.products = append(s.products, p)
// 	return nil
// }

// func (s *memoryProductStore) GetAll() []model.Product {
// 	return s.products
// }

//	func NewMemoryStore() ProductStore {
//		return &memoryProductStore{}
//	}
package store
