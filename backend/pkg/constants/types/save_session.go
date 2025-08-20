package types

// SaveSession is a struct used to insert and load session information to and from the database
type SaveSession struct {
	ID     uint64
	UserID uint64
	Token  string
}
