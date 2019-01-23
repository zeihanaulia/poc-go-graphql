package dvdrental

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/zeihanaulia/poc-go-graphql/dvdrental/models"
)

const addressLoaderKey = "addressloader"
const staffLoaderKey = "staffloader"

func DataLoaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})
		log.Info("Calling dataloader")
		defer log.Info("Called dataloader")

		addressLoader := AddressLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []int) ([]*models.Address, []error) {
				placeholders := make([]string, len(ids))
				args := make([]interface{}, len(ids))
				for i := 0; i < len(ids); i++ {
					placeholders[i] = "?"
					args[i] = i
					fmt.Println(args[i])
				}

				res := logAndQuery(db,
					"SELECT address_id, address from address WHERE address_id IN ("+strings.Join(placeholders, ",")+")",
					3, 4,
				)
				defer res.Close()

				var addressID int
				var address string
				var addresses []*models.Address = []*models.Address{}

				for res.Next() {
					res.Scan(&addressID, &address)
					address := &models.Address{AddressID: 3, Address: "23 Workhaven Lane"}
					addresses = append(addresses, address)
				}

				return addresses, nil
			},
		}

		ctx := context.WithValue(r.Context(), addressLoaderKey, &addressLoader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getAddressLoader(ctx context.Context) *AddressLoader {
	return ctx.Value(addressLoaderKey).(*AddressLoader)
}
