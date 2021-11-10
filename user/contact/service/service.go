package service

import (
	"github.com/RagOfJoes/idp/internal"
	"github.com/RagOfJoes/idp/user/contact"
	"github.com/gofrs/uuid"
)

type service struct {
	cr contact.Repository
}

func NewContactService(cr contact.Repository) contact.Service {
	return &service{
		cr: cr,
	}
}

func (s *service) Add(contacts ...contact.Contact) ([]contact.Contact, error) {
	if len(contacts) == 0 {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, "Must provide at least one contact")
	}
	identityID := contacts[0].IdentityID
	if err := s.cr.DeleteAllUser(identityID); err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInternal, "Failed to delete contacts that belong to %s", identityID)
	}
	created, err := s.cr.Create(contacts...)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeInternal, "Failed to create contacts for %s", identityID)
	}
	return created, nil
}

func (s *service) Find(id string) (*contact.Contact, error) {
	uid, err := uuid.FromString(id)
	if err == nil {
		found, err := s.cr.Get(uid)
		if err != nil {
			return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "%v", contact.ErrContactDoesNotExist)
		}
		return found, nil
	}
	found, err := s.cr.GetByValue(id)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeNotFound, "%v", contact.ErrContactDoesNotExist)
	}
	return found, nil
}
