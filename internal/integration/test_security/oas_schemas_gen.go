// Code generated by ogen, DO NOT EDIT.

package api

type BasicAuth struct {
	Username string
	Password string
}

// GetUsername returns the value of Username.
func (s *BasicAuth) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *BasicAuth) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *BasicAuth) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *BasicAuth) SetPassword(val string) {
	s.Password = val
}

type BearerToken struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerToken) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerToken) SetToken(val string) {
	s.Token = val
}

type CookieKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *CookieKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *CookieKey) SetAPIKey(val string) {
	s.APIKey = val
}

// DisjointSecurityOK is response for DisjointSecurity operation.
type DisjointSecurityOK struct{}

type HeaderKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *HeaderKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *HeaderKey) SetAPIKey(val string) {
	s.APIKey = val
}

// IntersectSecurityOK is response for IntersectSecurity operation.
type IntersectSecurityOK struct{}

// OptionalSecurityOK is response for OptionalSecurity operation.
type OptionalSecurityOK struct{}

type QueryKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *QueryKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *QueryKey) SetAPIKey(val string) {
	s.APIKey = val
}
