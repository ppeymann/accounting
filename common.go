package accounting

import (
	"encoding/json"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type (
	// BaseRepository is abstract interface that all repositories must implement its methods
	BaseRepository interface {
		// BaseRepository runs AutoMigrate for expected repository model
		Migrate() error

		// Name repository associated table name
		Name() string

		// Model returns *gorm.DB instance for repository
		Model() *gorm.DB
	}

	// BaseResult a basic GoLang struct which includes the following fields: Success, Messages, ResultCount, Result
	// It is the unified response model for entire service api calls
	//
	// swagger:model BaseResult
	BaseResult struct {
		// Status of response
		Status int `json:"-"`

		// Errors provides list off error that occurred in processing request
		Errors []string `json:"errors" mapstructure:"errors"`

		// ResultCount specified number of records that returned in result_count field expected result been array.
		ResultCount int64 `json:"result_count,omitempty" mapstructure:"result_count"`

		// Result single/array of any type (object/number/string/boolean) that returns as response
		Result interface{} `json:"result" mapstructure:"result"`
	}

	// ContextUser a basic Golang struct which includes the following fields: ID, Roles, Permissions
	// It used as User object that holds required information to identifying user and relegated roles and permissions
	// It may embedded into any Input model, or you may build your own model without it
	ContextUser struct {
		ID          string        `json:"id" mapstructure:"id"`
		Roles       []AccountRole `json:"roles" mapstructure:"roles"`
		Permissions []string      `json:"permission" mapstructure:"permissions"`
	}

	// AccountRole represents enum for determining account role (medicean, nurse, experiment, imaging, ...)
	AccountRole string
)

var (
	ErrUnimplementedRequest = errors.New("request is not implemented")
	ErrUnhandled            = errors.New("an unhandled error occurred during processing the request")
	ErrNotFound             = errors.New("not found")
	ErrInternalServer       = errors.New("internal server error")
	ErrEntityAlreadyExist   = errors.New("entity with specified properties already exist")
	ErrPermissionDenied     = errors.New("permission denied")
)

const (
	ContextUserKey          string = "CONTEXT_USER"
	UserSessionKey          string = "USER_SESSION"
	AuthorizationFailed     string = "authorization failed"
	ProvideRequiredParam    string = "please provide required params"
	ProvideRequiredJsonBody string = "please provide required JSON body"
)

func SearchConfig(query string) string {
	split := strings.Split(strings.TrimSpace(query), " ")
	query = strings.Join(split, ":* & ")
	query += ":*"

	return query
}

// ToJson is method for parsing ContextUser to json string.
func (p *ContextUser) ToJson() (string, error) {
	js, err := json.Marshal(p)
	return string(js), err
}

// FromJson is method for parsing ContextUser from json string.
func (p *ContextUser) FromJson(val string) error {
	return json.Unmarshal([]byte(val), p)
}
