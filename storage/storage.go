package storage

import (
	"github.com/onedss/filebrowser/auth"
	"github.com/onedss/filebrowser/settings"
	"github.com/onedss/filebrowser/share"
	"github.com/onedss/filebrowser/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
