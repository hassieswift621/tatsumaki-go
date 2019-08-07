package tatsumakigo

// Client is the main interface to interact with the API.
type Client struct {
	restClient *restClient
}

// New creates a new instance of the Tatsumaki client.
func New(token string) *Client {
	return &Client{
		newRestClient(token),
	}
}

// AdjustGuildUserPoints adjusts a user's points in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserPoints(guildId string, userId string, amount int, action Action) (*GuildUserPoints, error) {
	return t.restClient.adjustGuildUserPoints(guildId, userId, amount, action)
}

// AdjustGuildUserScore adjusts a user's score in a guild.
// The amount must be between 0 and 50,000 (inclusive) and must be above 0 if the action is add or remove.
func (t *Client) AdjustGuildUserScore(guildId string, userId string, amount int, action Action) (*GuildUserScore, error) {
	return t.restClient.adjustGuildUserScore(guildId, userId, amount, action)
}

// GuildLeaderboard gets the leaderboard for a guild.
func (t *Client) GuildLeaderboard(guildId string) ([]*GuildRankedUser, error) {
	return t.restClient.guildLeaderboard(guildId)
}

// GuildUserStats gets a user's stats for a guild.
// @me is accepted for the user ID, which will retrieve stats for yourself.
func (t *Client) GuildUserStats(guildId string, userId string) (*GuildUserStats, error) {
	return t.restClient.guildUserStats(guildId, userId)
}

// User gets a Tatsumaki user profile.
func (t *Client) User(userId string) (*User, error) {
	return t.restClient.user(userId)
}