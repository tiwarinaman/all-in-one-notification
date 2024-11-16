package service

import (
	"errors"
	"tenant-management-service/internal/model"
	"tenant-management-service/internal/repository"
	pkgerr "tenant-management-service/pkg/error"
	"tenant-management-service/pkg/utils"
)

type TenantService struct {
	repo *repository.TenantRepository
}

func NewTenantService(repo *repository.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) CreateTenant(name, email, phone, billingTier, defaultLanguage string) (*model.Tenant, error) {

	// Perform validations
	if err := utils.ValidateNonEmptyString(name, "name"); err != nil {
		return nil, err
	}
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := utils.ValidatePhone(phone); err != nil {
		return nil, err
	}
	if err := utils.ValidateAllowedValues(billingTier, "BillingTier", []string{"basic", "standard", "enterprise"}); err != nil {
		return nil, err
	}
	if err := utils.ValidateMaxLength(defaultLanguage, "DefaultLanguage", 5); err != nil {
		return nil, err
	}

	// Generate secure client ID and client secret
	clientID := utils.GenerateUUID()                   // UUID is fine for client ID
	clientSecret, err := utils.GenerateSecureToken(32) // Secure token for client secret
	if err != nil {
		return nil, errors.New("failed to generate client secret: " + err.Error())
	}

	tenant := &model.Tenant{
		Name:            name,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Email:           email,
		Phone:           phone,
		BillingTier:     billingTier,
		DefaultLanguage: defaultLanguage,
	}

	// Save the tenant in the repository
	if err := s.repo.Create(tenant); err != nil {
		return nil, errors.New("failed to create tenant: " + err.Error())
	}

	return tenant, nil
}

// GetTenantByID retrieves a tenant by its ID.
func (s *TenantService) GetTenantByID(id uint) (*model.Tenant, error) {
	tenant, err := s.repo.FindById(id)
	if err != nil {
		return nil, errors.New("tenant not found")
	}
	return tenant, nil
}

// UpdateTenant updates the details of an existing tenant.
func (s *TenantService) UpdateTenant(id uint, name, email, phone, billingTier, defaultLanguage string) error {
	tenant, err := s.repo.FindById(id)
	if err != nil {
		return errors.New("tenant not found")
	}

	// Update fields if provided
	if name != "" {
		if err := utils.ValidateNonEmptyString(name, "Name"); err != nil {
			return err
		}
		tenant.Name = name
	}
	if email != "" {
		if err := utils.ValidateEmail(email); err != nil {
			return err
		}
		tenant.Email = email
	}
	if phone != "" {
		if err := utils.ValidatePhone(phone); err != nil {
			return err
		}
		tenant.Phone = phone
	}
	if billingTier != "" {
		if err := utils.ValidateAllowedValues(billingTier, "BillingTier", []string{"basic", "standard", "enterprise"}); err != nil {
			return err
		}
		tenant.BillingTier = billingTier
	}
	if defaultLanguage != "" {
		if err := utils.ValidateMaxLength(defaultLanguage, "DefaultLanguage", 5); err != nil {
			return err
		}
		tenant.DefaultLanguage = defaultLanguage
	}

	// Save the updated tenant
	return s.repo.Update(tenant)
}

// DeleteTenant deletes a tenant by its ID.
func (s *TenantService) DeleteTenant(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.New("failed to delete tenant: " + err.Error())
	}
	return nil
}

// ValidateClientCredentials validates the client_id and client_secret.
func (s *TenantService) ValidateClientCredentials(clientID, clientSecret string) (bool, error) {
	tenant, err := s.repo.FindByClientId(clientID)
	if err != nil {
		if errors.Is(err, pkgerr.ErrNotFound) {
			return false, nil // Invalid client_id
		}
		return false, err // Internal error
	}

	// Compare client_secret
	if tenant.ClientSecret != clientSecret {
		return false, nil
	}
	return true, nil
}
