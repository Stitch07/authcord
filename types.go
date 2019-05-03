package authcord

// User is a Discord user object as documented at https://discordapp.com/developers/docs/resources/user#user-object
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar,omitempty"`
	Bot           bool   `json:"bot,omitempty"`
	MFAEnabled    bool   `json:"mfa_enabled,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	Email         string `json:"email,omitempty"` // only provided if email scope is passed
	Flags         uint   `json:"flags,omitempty"`
	PremiumType   uint8  `json:"premium_type,omitempty"`
}

// Guild is a partial Guild object
type Guild struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Owner       bool   `json:"owner"`
	Permissions uint   `json:"permissions"`
}
