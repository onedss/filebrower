package auth

import (
	"net/http"
	"os"

	"github.com/onedss/filebrowser/errors"
	"github.com/onedss/filebrowser/settings"
	"github.com/onedss/filebrowser/users"
)

// MethodProxyAuth is used to identify no auth.
const MethodProxyAuth settings.AuthMethod = "proxy"

// ProxyAuth is a proxy implementation of an auther.
type ProxyAuth struct {
	Header string `json:"header"`
}

// Auth authenticates the user via an HTTP header.
func (a ProxyAuth) Auth(r *http.Request, sto users.Store, root string) (*users.User, error) {
	username := r.Header.Get(a.Header)
	user, err := sto.Get(root, username)
	if err == errors.ErrNotExist {
		return nil, os.ErrPermission
	}

	return user, err
}

// LoginPage tells that proxy auth doesn't require a login page.
func (a ProxyAuth) LoginPage() bool {
	return false
}
