package constants

// Workspace status.
const (
	SpaceStatusEnabled  int32 = iota + 1 // => "enabled"
	SpaceStatusDisabled                  // => "disabled"
)

// Workspace System roles.
const (
	RoleTypeSystem int32 = iota + 1
	RoleTypeCustom
)
const (
	RoleIdSpaceOwner     = "rs-1001"
	RoleIdSpaceAdmin     = "rs-1002"
	RoleIdSpaceDeveloper = "rs-1003"
	RoleIdSpaceOperator  = "rs-1004"
	RoleIdSpaceVisitor   = "rs-1005"
)
