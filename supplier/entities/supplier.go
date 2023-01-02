package entities

type SupplierModel struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string `gorm:"not null"`
	Address     string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
}

func getCategories()  {}
func getCatogory()    {}
func saveCategory()   {}
func updateCategory() {}
func deleteCategory() {}
