package dbx

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver module

	"github.com/gavrilaf/spawn/pkg/dbx/mdl"
	"github.com/gavrilaf/spawn/pkg/errx"
	"github.com/gavrilaf/spawn/pkg/senv"
)

const (
	Scope = "dbx"
)

// Connect - connect to spawn database
func Connect(en *senv.Environment) (Database, error) {
	db, err := sqlx.Connect(en.GetDBOpts().Driver, en.GetDBOpts().DataSource)
	if err != nil {
		return nil, errx.ErrEnvironment(Scope, "Couldn't connect to postgre database: %v", err)
	}

	return &Bridge{db}, nil
}

// Database is an interface to the Spawn storage
type Database interface {
	Close() error

	GetClients() ([]mdl.Client, error)

	// User

	RegisterUser(username string, password string, device mdl.DeviceInfo) (*mdl.UserProfile, error)

	GetUserProfile(id string) (*mdl.UserProfile, error)
	FindUserProfile(username string) (*mdl.UserProfile, error)

	UpdateUserPermissions(id string, permissions mdl.Permissions) error

	UpdateUserPersonalInfo(id string, info mdl.PersonalInfo) error
	UpdateUserCountry(id string, contry string) error
	UpdateUserPhoneNumber(id string, phone mdl.PhoneNumber) error

	ReadAllUserProfiles() (<-chan *mdl.UserProfile, <-chan error)

	AddDevice(device mdl.DeviceInfo) error
	UpdateDevice(device mdl.DeviceInfo) error
	RemoveDevice(userID string, deviceID string) error

	ConfirmDevice(userID string, deviceID string) error
	SetDeviceFingerprint(userID string, deviceID string, fingerprint []byte) error

	GetUserDevice(userID string, deviceID string) (*mdl.DeviceInfo, error)
	GetUserDevices(userID string) ([]mdl.DeviceInfo, error)
	GetUserDevicesEx(userID string) ([]mdl.DeviceInfoEx, error)

	LogUserLogin(userID string, deviceID string, userAgent string, ip string, region string) error

	// Accounts

	GetSupportedAccounts() ([]mdl.AccountMeta, error)

	GetAccountsAllowedForUser(userID string) ([]mdl.AccountMeta, error)
	GetUserAccounts(userID string) ([]mdl.Account, error)

	AddAccount(userID string, currency string, name string) (*mdl.Account, error)

	GetAccount(accountID string) (*mdl.Account, error)
}

//////////////////////////////////////////////////////////////////////////////

type Bridge struct {
	conn *sqlx.DB
}

func (db *Bridge) Close() error {
	if db == nil || db.conn == nil {
		return nil
	}

	return db.conn.Close()
}
