package auth

import (
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/gavrilaf/spawn/pkg/cryptx"
	mdl "github.com/gavrilaf/spawn/pkg/dbx/mdl"
)

type LoginDTO struct {
	ClientID   string `json:"client_id" form:"client_id" binding:"required"`
	DeviceID   string `json:"device_id" form:"device_id" binding:"required"`
	DeviceName string `json:"device_name" form:"device_name" binding:"required"`
	AuthType   string `json:"auth_type" form:"auth_type" binding:"required"`
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	Locale     string `json:"locale" form:"locale"`
	Lang       string `json:"lang" form:"lang"`
	Signature  string `json:"signature" form:"signature" binding:"required"`
}

type RegisterDTO struct {
	ClientID   string `json:"client_id" form:"client_id" binding:"required"`
	DeviceID   string `json:"device_id" form:"device_id" binding:"required"`
	DeviceName string `json:"device_name" form:"device_name" binding:"required"`
	Username   string `json:"username" form:"username" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required,ascii,min=8"`
	Locale     string `json:"locale" form:"locale"`
	Lang       string `json:"lang" form:"lang"`
	Signature  string `json:"signature" form:"signature" binding:"required"`
}

type RefreshDTO struct {
	AuthToken    string `json:"auth_token" form:"auth_token" binding:"required"`
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}

type PermissionsDTO struct {
	IsDeviceConfirmed bool `json:"is_device_confirmed" structs:"is_device_confirmed"`
	IsEmailConfirmed  bool `json:"is_email_confirmed" structs:"is_email_confirmed"`
	Is2FARequired     bool `json:"is_2fa_required" structs:"is_2fa_required"`
	IsLocked          bool `json:"is_locked" structs:"is_locked"`
	Scopes            int  `json:"scopes" structs:"scopes"`
}

type AuthTokenDTO struct {
	AuthToken    string         `json:"auth_token" structs:"auth_token"`
	RefreshToken string         `json:"refresh_token" structs:"refresh_token"`
	Expire       time.Time      `json:"expire" structs:"expire"`
	Permissions  PermissionsDTO `json:"permissions" structs:"permissions"`
}

type LoginContext struct {
	UserAgent  string
	IP         string
	Region     string
	DeviceName string
}

////////////////////////////////////////////////////////////////////////

func (p *LoginDTO) FixLocale() {
	if len(p.Locale) == 0 {
		p.Locale = "en"
	}

	if len(p.Lang) == 0 {
		p.Lang = "en"
	}
}

func (p *LoginDTO) GetSignature(key []byte) string {
	msg := p.ClientID + p.DeviceID + p.Username
	return cryptx.GenerateSignature(msg, key)
}

func (p *LoginDTO) CheckSignature(key []byte) error {
	msg := p.ClientID + p.DeviceID + p.Username
	return cryptx.CheckSignature(msg, p.Signature, key)
}

func (p *LoginDTO) CheckPassword(pswHash string) bool {
	return cryptx.CheckPassword(p.Password, pswHash) == nil
}

func (p *LoginDTO) CheckDevice(devices []string) bool {
	for _, d := range devices {
		if p.DeviceID == d {
			return true
		}
	}
	return false
}

func (p *LoginDTO) String() string {
	return fmt.Sprintf("LoginDTO(%s, %s, %s, %s, %s, %s, %s, %s)",
		p.ClientID, p.DeviceID, p.DeviceID, p.DeviceName, p.AuthType, p.Username, p.Locale, p.Lang)
}

func (p *LoginDTO) CreateDevice() mdl.DeviceInfo {
	return mdl.DeviceInfo{
		DeviceID: p.DeviceID,
		Name:     p.DeviceName,
		Locale:   p.Locale,
		Lang:     p.Lang,
	}
}

////////////////////////////////////////////////////////////////////////

func (p *RegisterDTO) FixLocale() {
	if len(p.Locale) == 0 {
		p.Locale = "en"
	}

	if len(p.Lang) == 0 {
		p.Lang = "en"
	}
}

func (p *RegisterDTO) GetSignature(key []byte) string {
	msg := p.ClientID + p.DeviceID + p.Username
	return cryptx.GenerateSignature(msg, key)
}

func (p *RegisterDTO) CheckSignature(key []byte) error {
	msg := p.ClientID + p.DeviceID + p.Username
	return cryptx.CheckSignature(msg, p.Signature, key)
}

func (p *RegisterDTO) String() string {
	return fmt.Sprintf("RegisterDTO(%s, %s, %s, %s, %s, %s)",
		p.ClientID, p.DeviceID, p.DeviceName, p.Username, p.Locale, p.Lang)
}

func (p *RegisterDTO) CreateDevice() mdl.DeviceInfo {
	return mdl.DeviceInfo{
		DeviceID: p.DeviceID,
		Name:     p.DeviceName,
		Locale:   p.Locale,
		Lang:     p.Lang,
	}
}

////////////////////////////////////////////////////////////////////////

func (p *AuthTokenDTO) ToMap() map[string]interface{} {
	pm := structs.Map(*p)
	pm["expire"] = p.Expire.Format(time.RFC3339) // Fixed time format
	return pm
}

////////////////////////////////////////////////////////////////////////

func (p *LoginContext) String() string {
	return fmt.Sprintf("LoginContext(%s, %s, %s, %s)", p.UserAgent, p.IP, p.Region, p.DeviceName)
}
