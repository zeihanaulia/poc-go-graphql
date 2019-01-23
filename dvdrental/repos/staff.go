package repos

import (
	"github.com/zeihanaulia/poc-go-graphql/dvdrental/models"
)

type Staff interface {
	Create(models.Staff) models.Staff
}
