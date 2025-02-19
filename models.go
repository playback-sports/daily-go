package daily

import "time"

// DomainConfig is used when getting and setting the domain configuration.
// https://docs.daily.co/reference#get-domain-configuration
type DomainConfig struct {
	DomainName *string `json:"domain_name,omitempty"`
	Config     *Config `json:"config,omitempty"`
}

// Config options contained within the DomainConfig that can be changed by the
// user.
type Config struct {
	RedirectOnMeetingExit *string `json:"redirect_on_meeting_exit,omitempty"`
	HideDailyBranding     *bool   `json:"hide_daily_branding,omitempty"`
	HIPPAA                *bool   `json:"hipaa,omitempty"`
	IntercomAutoRecord    *bool   `json:"intercom_auto_record,omitempty"`
	Lang                  *string `json:"lang,omitempty"`
}

// Room contains information about a video location and configuration.
// https://docs.daily.co/reference#rooms
type Room struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	APICreated bool        `json:"api_created"`
	Privacy    RoomPrivacy `json:"privacy"`
	URL        string      `json:"url"`
	CreatedAt  time.Time   `json:"created_at"`
	Config     *RoomConfig `json:"config"`
}

// RoomPrivacy controls who can join a meeting.
type RoomPrivacy string

const (
	Public  RoomPrivacy = "public"
	Private RoomPrivacy = "private"
	Org     RoomPrivacy = "org"
)

type PermissionType string

const (
	Video       PermissionType = "video"
	Audio       PermissionType = "audio"
	ScreenAudio PermissionType = "screenAudio"
	ScreenVideo PermissionType = "screenVideo"
)

type Permissions struct {
	CanSend     *[]PermissionType `json:"canSend,omitempty"`
	HasPresence *bool             `json:"hasPresence,omitempty"`
}

// RoomConfig is the configuration for a room.
type RoomConfig struct {
	NotBefore                *int64  `json:"nbf,omitempty"` // Unix timestamp in seconds
	ExpiresAt                *int64  `json:"exp,omitempty"` // Unix timestamp in seconds
	StartVideoOff            *bool   `json:"start_video_off,omitempty"`
	StartAudioOff            *bool   `json:"start_audio_off,omitempty"`
	MaxParticipants          *int32  `json:"max_participants,omitempty"`
	AutoJoin                 *bool   `json:"autojoin,omitempty"`
	EnableKnocking           *bool   `json:"enable_knocking,omitempty"`
	EnableScreenShare        *bool   `json:"enable_screenshare,omitempty"`
	EnableChat               *bool   `json:"enable_chat,omitempty"`
	OwnerOnlyBroadcast       *bool   `json:"owner_only_broadcast,omitempty"`
	EnableRecording          *string `json:"enable_recording,omitempty"`
	EjectAtRoomExpiry        *bool   `json:"eject_at_room_exp,omitempty"`
	EjectAfterElapsed        *int32  `json:"eject_after_elapsed,omitempty"`
	Lang                     *string `json:"lang,omitempty"`
	MeetingJoinHook          *string `json:"meeting_join_hook,omitempty"`
	SignalingType            *string `json:"signaling_impl,omitempty"` // In JSON, they spell it 'signaling' so we use that
	SFUSwitchover            *int32  `json:"sfu_switchover,omitempty"`
	EnableMeshSFU            *bool   `json:"enable_mesh_sfu,omitempty"`
	EnableTerseLogging       *bool   `json:"enable_terse_logging,omitempty"`
	EnableHiddenParticipants *bool   `json:"enable_hidden_participants,omitempty"`
}

// MeetingToken is the configuration that controls room access and session configuration on a per-user basis.
type MeetingToken struct {
	NotBefore           *int64       `json:"nbf,omitempty"` // Unix timestamp in seconds
	ExpiresAt           *int64       `json:"exp,omitempty"` // Unix timestamp in seconds
	RoomName            *string      `json:"room_name,omitempty"`
	IsOwner             *bool        `json:"is_owner,omitempty"`
	UserName            *string      `json:"user_name,omitempty"`
	UserID              *string      `json:"user_id,omitempty"`
	EnableScreenShare   *bool        `json:"enable_screenshare,omitempty"`
	StartVideoOff       *bool        `json:"start_video_off,omitempty"`
	StartAudioOff       *bool        `json:"start_audio_off,omitempty"`
	EnableRecording     *string      `json:"enable_recording,omitempty"`
	StartCloudRecording *bool        `json:"start_cloud_recording,omitempty"`
	CloseTabOnExit      *bool        `json:"close_tab_on_exit,omitempty"`
	EjectAtRoomExpiry   *bool        `json:"eject_at_room_exp,omitempty"`
	EjectAfterElapsed   *int32       `json:"eject_after_elapsed,omitempty"`
	Lang                *string      `json:"lang,omitempty"`
	Permissions         *Permissions `json:"permissions,omitempty"`
}

// Layout is a configuration for started a recording
type Layout struct {
	Preset string `json:"preset"`
}

type Recording struct {
	Id              string        `json:"id"`
	StartTs         int           `json:"start_ts"`
	Status          string        `json:"status"`
	MaxParticipants int           `json:"max_participants"`
	RoomName        string        `json:"room_name"`
	Tracks          []interface{} `json:"tracks"`
	Duration        int           `json:"duration"`
	ShareToken      string        `json:"share_token"`
}

// String returns a pointer to the string.
func String(s string) *string {
	return &s
}

// Int64 returns a pointer to the int64.
func Int64(i int64) *int64 {
	return &i
}

// Int32 returns a pointer to the int32.
func Int32(i int32) *int32 {
	return &i
}

// Timestamp returns number of seconds since epoch, consistent wih Daily's
// API expectations.
func Timestamp(t time.Time) *int64 {
	return Int64(t.Unix())
}

// Bool returns a pointer to the bool.
func Bool(b bool) *bool {
	return &b
}

// True returns a pointer to true.
func True() *bool {
	return Bool(true)
}

// False returns a pointer to false.
func False() *bool {
	return Bool(false)
}
