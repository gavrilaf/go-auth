package types

const (
	Realm = "Spawn"

	TokenHeadName    = "Bearer"
	SigningAlgorithm = "HS256"
	TokenLookup      = "Authorization"

	AuthTypeSimple = "simple"

	EndpointKey = "EnpointKey"
)

var EmptySuccessResponse = map[string]interface{}{"success": true}

type Endpoint struct {
	Path   string
	Method string
}

type Access struct {
	NeedDevice bool
	NeedEmail  bool
	MinScope   int
}

type EndpointAccess struct {
	Group string
	Endpoint
	Access
}

func GetEndpointKey(group string, endpoint Endpoint) string {
	return group + endpoint.Path + ":" + endpoint.Method
}
