// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"

	// EdgePlaylists holds the string denoting the playlists edge name in mutations.
	EdgePlaylists = "playlists"
	// EdgeVideos holds the string denoting the videos edge name in mutations.
	EdgeVideos = "videos"

	// Table holds the table name of the user in the database.
	Table = "users"
	// PlaylistsTable is the table the holds the playlists relation/edge.
	PlaylistsTable = "playlists"
	// PlaylistsInverseTable is the table name for the Playlist entity.
	// It exists in this package in order to avoid circular dependency with the "playlist" package.
	PlaylistsInverseTable = "playlists"
	// PlaylistsColumn is the table column denoting the playlists relation/edge.
	PlaylistsColumn = "User_ID"
	// VideosTable is the table the holds the videos relation/edge.
	VideosTable = "videos"
	// VideosInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	VideosInverseTable = "videos"
	// VideosColumn is the table column denoting the videos relation/edge.
	VideosColumn = "User_ID"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserID,
}
