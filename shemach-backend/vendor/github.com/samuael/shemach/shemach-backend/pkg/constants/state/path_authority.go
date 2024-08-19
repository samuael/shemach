package state

type (
	Permission struct {
		Roles   []string
		Methods []string
	}
	Authority map[string]*Permission
)

// Authorities this map represents a map of pathroutes and their permissions
// and roles that are allowed to
var Authorities = Authority{
	// "/api/secretary/new/": &Permission{
	// 	Roles: []string{ADMIN},
	// },
	// "/api/inspector/new/": &Permission{
	// 	Roles: []string{ADMIN},
	// },
	// "/api/secretary/": &Permission{
	// 	Roles: []string{ADMIN},
	// },
	// "/api/inspector/": &Permission{
	// 	Roles: []string{ADMIN},
	// },
	// "/api/admin/inspectors/": &Permission{
	// 	Roles: []string{ADMIN},
	// },
}
