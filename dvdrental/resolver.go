//go:generate go run scripts/gqlgen.go -v
package dvdrental

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/zeihanaulia/poc-go-graphql/dvdrental/models"
)

type Resolver struct {
	staffs    []models.Staff
	addresses []models.Address
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Staff() StaffResolver {
	return &staffResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) RegisterStaff(ctx context.Context, input NewStaff) (models.Staff, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	defer db.Close()
	fmt.Println(err)
	staff := models.Staff{
		StaffID:   rand.Intn(2147483647),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Username:  input.Username,
		Password:  input.Password,
		StoreID:   1,
		AddressID: 3,
	}

	db.Table("staff").Create(&staff)

	return staff, nil
}

func (r *mutationResolver) CreateAddress(ctx context.Context, input *NewAddress) (models.Address, error) {
	address := models.Address{
		AddressID:  rand.Int(),
		Address:    input.Address,
		Address2:   input.Address2,
		District:   input.District,
		PostalCode: input.PostalCode,
		Phone:      input.Phone,
	}
	r.addresses = append(r.addresses, address)
	return address, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Staffs(ctx context.Context) ([]models.Staff, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	defer db.Close()
	fmt.Println(err)
	staffs := []models.Staff{}
	db.Debug().Table("staff").Find(&staffs)

	return staffs, nil
}

func (r *queryResolver) Addresses(ctx context.Context) ([]models.Address, error) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	defer db.Close()
	fmt.Println(err)
	addresses := []models.Address{}
	db.Table("address").Find(&addresses)

	return addresses, nil
}

type staffResolver struct{ *Resolver }

func (r *staffResolver) Address(ctx context.Context, obj *models.Staff) (models.Address, error) {
	address, err := getAddressLoader(ctx).Load(obj.AddressID)
	if err != nil {
		fmt.Println(err)
	}

	return models.Address{
		AddressID: address.AddressID,
		Address:   address.Address,
	}, nil
}
