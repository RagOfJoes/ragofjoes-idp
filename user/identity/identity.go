package identity

import (
	idp "github.com/RagOfJoes/idp"
	"github.com/RagOfJoes/idp/user/address"
	"github.com/RagOfJoes/idp/user/credential"
	"github.com/gofrs/uuid"
)

// Identity defines the base Identity model
type Identity struct {
	idp.BaseSoftDelete
	Avatar    string `json:"avatar" gorm:"size:1024;" validate:"url,min=1,max=1024,"`
	FirstName string `json:"first_name" gorm:"size:64;not null" validate:"required,min=1,max=64,alphanumunicode"`
	LastName  string `json:"last_name" gorm:"size:64;not null" validate:"required,min=1,max=64,alphanumunicode"`
	// Email is the primary email that will be used for account
	// security related notifications
	Email string `json:"email" gorm:"index;not null;" validate:"email,required"`

	Credentials         []credential.Credential
	VerifiableAddresses []address.VerifiableAddress
}

// Repository defines an interface that allows
// Identity domain data to be persisted through different
// dbs
type Repository interface {
	Create(Identity) (*Identity, error)
	Get(id uuid.UUID, critical bool) (*Identity, error)
	GetIdentifier(identifier string, critical bool) (*Identity, error)
	Update(Identity) (*Identity, error)
	Delete(uuid.UUID) error
}
