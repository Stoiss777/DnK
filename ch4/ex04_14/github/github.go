package github

import "time"

const GitHubAPIVer = "2026-03-10"
const GetIssuesURL = "https://api.github.com/repos/%s/%s/issues"
const GetMilestonesURL = "https://api.github.com/repos/%s/%s/milestones"
const GetContributorsURL = "https://api.github.com/repos/%s/%s/contributors"
const DeviceCodeURL = "https://github.com/login/device/code"
const AccessTokenURL = "https://github.com/login/oauth/access_token"
const WhoamiURL = "https://api.github.com/user"
const AccessTokenFile = ".token"
const ClientIDFile = ".client_id"
const AccessTokenRequestInterval = 5

type DeviceCodeResponse struct {
    DeviceCode string       `json:"device_code"`
    UserCode string         `json:"user_code"`
    VerificationUri string  `json:"verification_uri"`
    ExpiresIn int           `json:"expires_in"`
    Interval int            `json:"interval"`
}

type AccessTokenResponse struct {
    Error string               `json:"error"`
    ErrorDescription string    `json:"error_description"`
    AccessToken string         `json:"access_token"`
    ExpiresIn int              `json:"expires_in"`
    RefreshToken string        `json:"refresh_token"`
    RefreshTokenExpiresIn int  `json:"refresh_token_expires_in"`
    TokenType string           `json:"token_type"`
    Scope string               `json:"scope"`
}

type User struct {
    Login string    `json:"login"`
    ID int          `json:"id"`
    HTMLURL string  `json:"html_url"`
}

type Issue struct {
    ID int
    Number int
    HTMLURL string       `json:"html_url"`
    Title string
    State string
    User *User
    CreatedAt time.Time  `json:"created_at"`
    Body string          // in Markdown format
}

type Milestone struct {
    ID int
    Number int
    HTMLURL string       `json:"html_url"`
    Title string
    State string
    Creator *User
    Description string
}
