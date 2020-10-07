package usecase

// PermissionUsecase ...
type PermissionUsecase interface {
	ListPermission() error
	CreatePermission() error
}
