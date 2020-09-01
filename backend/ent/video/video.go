// Code generated by entc, DO NOT EDIT.

package video

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVideoID holds the string denoting the video_id field in the database.
	FieldVideoID = "video_id"

	// EdgePlaylistVideos holds the string denoting the playlist_videos edge name in mutations.
	EdgePlaylistVideos = "playlist_videos"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the video in the database.
	Table = "videos"
	// PlaylistVideosTable is the table the holds the playlist_videos relation/edge.
	PlaylistVideosTable = "playlist_videos"
	// PlaylistVideosInverseTable is the table name for the PlaylistVideo entity.
	// It exists in this package in order to avoid circular dependency with the "playlistvideo" package.
	PlaylistVideosInverseTable = "playlist_videos"
	// PlaylistVideosColumn is the table column denoting the playlist_videos relation/edge.
	PlaylistVideosColumn = "video_playlist_videos"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "videos"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_videos"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldVideoID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Video type.
var ForeignKeys = []string{
	"user_videos",
}
