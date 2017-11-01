package auth

import (
	"github.com/gavrilaf/spawn/pkg/cryptx"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/dgrijalva/jwt-go.v3"

	"encoding/hex"
	"time"
)

func (mw *Middleware) HandleLogin(p *LoginParcel) (*TokenParcel, error) {
	mw.Log.Infof("auth.HandleLogin, %v", p)

	// Check client
	client, err := mw.Storage.FindClientByID(p.ClientID)
	if err != nil {
		return nil, err
	}

	// Check signature
	if err = p.CheckSignature(client.Secret()); err != nil {
		return nil, errInvalidSignature
	}

	// Check user
	user, err := mw.Storage.FindUserByUsername(p.Username)
	if err != nil {
		return nil, err
	}

	if !p.CheckPassword(user.PasswordHash) {
		mw.Log.Errorf("Invalid password for %v", p.Username)
		return nil, errUserUnknown
	}

	// Check device
	if !p.CheckDevice(user.Devices) {
		// TODO: Send email about new device
		mw.Log.Errorf("Unknown device for %v", p.Username)
		return nil, errDeviceUnknown
	}

	// Generate token

	sessionId := mw.GenerateSessionID()
	refreshToken := mw.GenerateRefreshToken(sessionId)

	session := Session{ID: sessionId, RefreshToken: refreshToken, ClientID: client.ID(), ClientSecret: client.Secret(), UserID: user.ID}

	err = mw.Storage.StoreSession(session)
	if err != nil {
		return nil, err
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	now := time.Now()
	expire := now.Add(mw.Timeout)
	claims["session_id"] = session.ID
	claims["aud"] = session.ClientID
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = now.Unix()
	//claims["iss"] = "go-auth" // TODO: Fix it later

	tokenString, err := token.SignedString(session.ClientSecret)
	if err != nil {
		return nil, err
	}

	return &TokenParcel{AuthToken: tokenString, RefreshToken: refreshToken, Expire: expire}, nil
}

func (mw *Middleware) HandleRefresh(p *RefreshParcel) (*TokenParcel, error) {
	token, _ := mw.parseToken(p.AuthToken)
	claims := token.Claims.(jwt.MapClaims)

	sessionId := claims["session_id"].(string)
	origIat := int64(claims["orig_iat"].(float64))

	if origIat < time.Now().Add(-mw.MaxRefresh).Unix() {
		return nil, errTokenExpired
	}

	session, err := mw.Storage.FindSessionByID(sessionId)
	if err != nil {
		return nil, err
	}

	if p.RefreshToken != session.RefreshToken {
		return nil, errTokenInvalid
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(SigningAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	for key := range claims {
		newClaims[key] = claims[key]
	}

	now := time.Now()
	expire := now.Add(mw.Timeout)
	claims["exp"] = expire.Unix()

	tokenString, err := token.SignedString(session.ClientSecret)
	if err != nil {
		return nil, err
	}

	return &TokenParcel{AuthToken: tokenString, RefreshToken: "", Expire: expire}, nil
}

func (mw *Middleware) HandleRegister(p *RegisterParcel) error {
	mw.Log.Infof("auth.HandleRegister, %v", p)

	// Check client
	client, err := mw.Storage.FindClientByID(p.ClientID)
	if err != nil {
		mw.Log.Errorf("auth.HandleRegister, can't find client with ID = ", p.ClientID)
		return err
	}

	// Check signature
	if p.CheckSignature(client.Secret()) != nil {
		mw.Log.Errorf("auth.HandleRegister, invalid signature for %v", p)
		return errInvalidSignature
	}

	// Validate password

	pswHash, err := cryptx.GenerateHashedPassword(p.Password)
	if err != nil {
		mw.Log.Errorf("auth.HandleRegister, password generate error for %v", p)
	}

	return mw.Storage.AddUser(p.ClientID, p.DeviceID, p.Username, pswHash)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (mw *Middleware) CheckAccess(userId string, clientId string, c *gin.Context) bool {
	return true
}

func (mw *Middleware) GenerateSessionID() string {
	return uuid.NewV4().String()
}

func (mw *Middleware) GenerateRefreshToken(sessionId string) string {
	k, _ := cryptx.GenerateSaltedKey(sessionId)
	return hex.EncodeToString(k)
}
