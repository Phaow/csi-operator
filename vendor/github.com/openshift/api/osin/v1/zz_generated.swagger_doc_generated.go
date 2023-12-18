package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AllowAllPasswordIdentityProvider = map[string]string{
	"": "AllowAllPasswordIdentityProvider provides identities for users authenticating using non-empty passwords\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (AllowAllPasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_AllowAllPasswordIdentityProvider
}

var map_BasicAuthPasswordIdentityProvider = map[string]string{
	"": "BasicAuthPasswordIdentityProvider provides identities for users authenticating using HTTP basic auth credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (BasicAuthPasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_BasicAuthPasswordIdentityProvider
}

var map_DenyAllPasswordIdentityProvider = map[string]string{
	"": "DenyAllPasswordIdentityProvider provides no identities for users\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (DenyAllPasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_DenyAllPasswordIdentityProvider
}

var map_GitHubIdentityProvider = map[string]string{
	"":              "GitHubIdentityProvider provides identities for users authenticating using GitHub credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"clientID":      "clientID is the oauth client ID",
	"clientSecret":  "clientSecret is the oauth client secret",
	"organizations": "organizations optionally restricts which organizations are allowed to log in",
	"teams":         "teams optionally restricts which teams are allowed to log in. Format is <org>/<team>.",
	"hostname":      "hostname is the optional domain (e.g. \"mycompany.com\") for use with a hosted instance of GitHub Enterprise. It must match the GitHub Enterprise settings value that is configured at /setup/settings#hostname.",
	"ca":            "ca is the optional trusted certificate authority bundle to use when making requests to the server. If empty, the default system roots are used.  This can only be configured when hostname is set to a non-empty value.",
}

func (GitHubIdentityProvider) SwaggerDoc() map[string]string {
	return map_GitHubIdentityProvider
}

var map_GitLabIdentityProvider = map[string]string{
	"":             "GitLabIdentityProvider provides identities for users authenticating using GitLab credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"ca":           "ca is the optional trusted certificate authority bundle to use when making requests to the server If empty, the default system roots are used",
	"url":          "url is the oauth server base URL",
	"clientID":     "clientID is the oauth client ID",
	"clientSecret": "clientSecret is the oauth client secret",
	"legacy":       "legacy determines if OAuth2 or OIDC should be used If true, OAuth2 is used If false, OIDC is used If nil and the URL's host is gitlab.com, OIDC is used Otherwise, OAuth2 is used In a future release, nil will default to using OIDC Eventually this flag will be removed and only OIDC will be used",
}

func (GitLabIdentityProvider) SwaggerDoc() map[string]string {
	return map_GitLabIdentityProvider
}

var map_GoogleIdentityProvider = map[string]string{
	"":             "GoogleIdentityProvider provides identities for users authenticating using Google credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"clientID":     "clientID is the oauth client ID",
	"clientSecret": "clientSecret is the oauth client secret",
	"hostedDomain": "hostedDomain is the optional Google App domain (e.g. \"mycompany.com\") to restrict logins to",
}

func (GoogleIdentityProvider) SwaggerDoc() map[string]string {
	return map_GoogleIdentityProvider
}

var map_GrantConfig = map[string]string{
	"":                     "GrantConfig holds the necessary configuration options for grant handlers",
	"method":               "method determines the default strategy to use when an OAuth client requests a grant. This method will be used only if the specific OAuth client doesn't provide a strategy of their own. Valid grant handling methods are:\n - auto:   always approves grant requests, useful for trusted clients\n - prompt: prompts the end user for approval of grant requests, useful for third-party clients\n - deny:   always denies grant requests, useful for black-listed clients",
	"serviceAccountMethod": "serviceAccountMethod is used for determining client authorization for service account oauth client. It must be either: deny, prompt",
}

func (GrantConfig) SwaggerDoc() map[string]string {
	return map_GrantConfig
}

var map_HTPasswdPasswordIdentityProvider = map[string]string{
	"":     "HTPasswdPasswordIdentityProvider provides identities for users authenticating using htpasswd credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"file": "file is a reference to your htpasswd file",
}

func (HTPasswdPasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_HTPasswdPasswordIdentityProvider
}

var map_IdentityProvider = map[string]string{
	"":              "IdentityProvider provides identities for users authenticating using credentials",
	"name":          "name is used to qualify the identities returned by this provider",
	"challenge":     "challenge indicates whether to issue WWW-Authenticate challenges for this provider",
	"login":         "login indicates whether to use this identity provider for unauthenticated browsers to login against",
	"mappingMethod": "mappingMethod determines how identities from this provider are mapped to users",
	"provider":      "provider contains the information about how to set up a specific identity provider",
}

func (IdentityProvider) SwaggerDoc() map[string]string {
	return map_IdentityProvider
}

var map_KeystonePasswordIdentityProvider = map[string]string{
	"":                    "KeystonePasswordIdentityProvider provides identities for users authenticating using keystone password credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"domainName":          "domainName is required for keystone v3",
	"useKeystoneIdentity": "useKeystoneIdentity flag indicates that user should be authenticated by keystone ID, not by username",
}

func (KeystonePasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_KeystonePasswordIdentityProvider
}

var map_LDAPAttributeMapping = map[string]string{
	"":                  "LDAPAttributeMapping maps LDAP attributes to OpenShift identity fields",
	"id":                "id is the list of attributes whose values should be used as the user ID. Required. LDAP standard identity attribute is \"dn\"",
	"preferredUsername": "preferredUsername is the list of attributes whose values should be used as the preferred username. LDAP standard login attribute is \"uid\"",
	"name":              "name is the list of attributes whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity LDAP standard display name attribute is \"cn\"",
	"email":             "email is the list of attributes whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
}

func (LDAPAttributeMapping) SwaggerDoc() map[string]string {
	return map_LDAPAttributeMapping
}

var map_LDAPPasswordIdentityProvider = map[string]string{
	"":             "LDAPPasswordIdentityProvider provides identities for users authenticating using LDAP credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"url":          "url is an RFC 2255 URL which specifies the LDAP search parameters to use. The syntax of the URL is\n   ldap://host:port/basedn?attribute?scope?filter",
	"bindDN":       "bindDN is an optional DN to bind with during the search phase.",
	"bindPassword": "bindPassword is an optional password to bind with during the search phase.",
	"insecure":     "insecure, if true, indicates the connection should not use TLS. Cannot be set to true with a URL scheme of \"ldaps://\" If false, \"ldaps://\" URLs connect using TLS, and \"ldap://\" URLs are upgraded to a TLS connection using StartTLS as specified in https://tools.ietf.org/html/rfc2830",
	"ca":           "ca is the optional trusted certificate authority bundle to use when making requests to the server If empty, the default system roots are used",
	"attributes":   "attributes maps LDAP attributes to identities",
}

func (LDAPPasswordIdentityProvider) SwaggerDoc() map[string]string {
	return map_LDAPPasswordIdentityProvider
}

var map_OAuthConfig = map[string]string{
	"":                            "OAuthConfig holds the necessary configuration options for OAuth authentication",
	"masterCA":                    "masterCA is the CA for verifying the TLS connection back to the MasterURL. This field is deprecated and will be removed in a future release. See loginURL for details. Deprecated",
	"masterURL":                   "masterURL is used for making server-to-server calls to exchange authorization codes for access tokens This field is deprecated and will be removed in a future release. See loginURL for details. Deprecated",
	"masterPublicURL":             "masterPublicURL is used for building valid client redirect URLs for internal and external access This field is deprecated and will be removed in a future release. See loginURL for details. Deprecated",
	"loginURL":                    "loginURL, along with masterCA, masterURL and masterPublicURL have distinct meanings depending on how the OAuth server is run.  The two states are: 1. embedded in the kube api server (all 3.x releases) 2. as a standalone external process (all 4.x releases) in the embedded configuration, loginURL is equivalent to masterPublicURL and the other fields have functionality that matches their docs. in the standalone configuration, the fields are used as: loginURL is the URL required to login to the cluster: oc login --server=<loginURL> masterPublicURL is the issuer URL it is accessible from inside (service network) and outside (ingress) of the cluster masterURL is the loopback variation of the token_endpoint URL with no path component it is only accessible from inside (service network) of the cluster masterCA is used to perform TLS verification for connections made to masterURL For further details, see the IETF Draft: https://tools.ietf.org/html/draft-ietf-oauth-discovery-04#section-2",
	"assetPublicURL":              "assetPublicURL is used for building valid client redirect URLs for external access",
	"alwaysShowProviderSelection": "alwaysShowProviderSelection will force the provider selection page to render even when there is only a single provider.",
	"identityProviders":           "identityProviders is an ordered list of ways for a user to identify themselves",
	"grantConfig":                 "grantConfig describes how to handle grants",
	"sessionConfig":               "sessionConfig hold information about configuring sessions.",
	"tokenConfig":                 "tokenConfig contains options for authorization and access tokens",
	"templates":                   "templates allow you to customize pages like the login page.",
}

func (OAuthConfig) SwaggerDoc() map[string]string {
	return map_OAuthConfig
}

var map_OAuthTemplates = map[string]string{
	"":                  "OAuthTemplates allow for customization of pages like the login page",
	"login":             "login is a path to a file containing a go template used to render the login page. If unspecified, the default login page is used.",
	"providerSelection": "providerSelection is a path to a file containing a go template used to render the provider selection page. If unspecified, the default provider selection page is used.",
	"error":             "error is a path to a file containing a go template used to render error pages during the authentication or grant flow If unspecified, the default error page is used.",
}

func (OAuthTemplates) SwaggerDoc() map[string]string {
	return map_OAuthTemplates
}

var map_OpenIDClaims = map[string]string{
	"":                  "OpenIDClaims contains a list of OpenID claims to use when authenticating with an OpenID identity provider",
	"id":                "id is the list of claims whose values should be used as the user ID. Required. OpenID standard identity claim is \"sub\"",
	"preferredUsername": "preferredUsername is the list of claims whose values should be used as the preferred username. If unspecified, the preferred username is determined from the value of the id claim",
	"name":              "name is the list of claims whose values should be used as the display name. Optional. If unspecified, no display name is set for the identity",
	"email":             "email is the list of claims whose values should be used as the email address. Optional. If unspecified, no email is set for the identity",
	"groups":            "groups is the list of claims value of which should be used to synchronize groups from the OIDC provider to OpenShift for the user",
}

func (OpenIDClaims) SwaggerDoc() map[string]string {
	return map_OpenIDClaims
}

var map_OpenIDIdentityProvider = map[string]string{
	"":                         "OpenIDIdentityProvider provides identities for users authenticating using OpenID credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"ca":                       "ca is the optional trusted certificate authority bundle to use when making requests to the server If empty, the default system roots are used",
	"clientID":                 "clientID is the oauth client ID",
	"clientSecret":             "clientSecret is the oauth client secret",
	"extraScopes":              "extraScopes are any scopes to request in addition to the standard \"openid\" scope.",
	"extraAuthorizeParameters": "extraAuthorizeParameters are any custom parameters to add to the authorize request.",
	"urls":                     "urls to use to authenticate",
	"claims":                   "claims mappings",
}

func (OpenIDIdentityProvider) SwaggerDoc() map[string]string {
	return map_OpenIDIdentityProvider
}

var map_OpenIDURLs = map[string]string{
	"":          "OpenIDURLs are URLs to use when authenticating with an OpenID identity provider",
	"authorize": "authorize is the oauth authorization URL",
	"token":     "token is the oauth token granting URL",
	"userInfo":  "userInfo is the optional userinfo URL. If present, a granted access_token is used to request claims If empty, a granted id_token is parsed for claims",
}

func (OpenIDURLs) SwaggerDoc() map[string]string {
	return map_OpenIDURLs
}

var map_OsinServerConfig = map[string]string{
	"":            "Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"oauthConfig": "oauthConfig holds the necessary configuration options for OAuth authentication",
}

func (OsinServerConfig) SwaggerDoc() map[string]string {
	return map_OsinServerConfig
}

var map_RequestHeaderIdentityProvider = map[string]string{
	"":                         "RequestHeaderIdentityProvider provides identities for users authenticating using request header credentials\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"loginURL":                 "loginURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect interactive logins will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter\n  https://www.example.com/sso-login?then=${url}\n${query} is replaced with the current query string\n  https://www.example.com/auth-proxy/oauth/authorize?${query}",
	"challengeURL":             "challengeURL is a URL to redirect unauthenticated /authorize requests to Unauthenticated requests from OAuth clients which expect WWW-Authenticate challenges will be redirected here ${url} is replaced with the current URL, escaped to be safe in a query parameter\n  https://www.example.com/sso-login?then=${url}\n${query} is replaced with the current query string\n  https://www.example.com/auth-proxy/oauth/authorize?${query}",
	"clientCA":                 "clientCA is a file with the trusted signer certs.  If empty, no request verification is done, and any direct request to the OAuth server can impersonate any identity from this provider, merely by setting a request header.",
	"clientCommonNames":        "clientCommonNames is an optional list of common names to require a match from. If empty, any client certificate validated against the clientCA bundle is considered authoritative.",
	"headers":                  "headers is the set of headers to check for identity information",
	"preferredUsernameHeaders": "preferredUsernameHeaders is the set of headers to check for the preferred username",
	"nameHeaders":              "nameHeaders is the set of headers to check for the display name",
	"emailHeaders":             "emailHeaders is the set of headers to check for the email address",
}

func (RequestHeaderIdentityProvider) SwaggerDoc() map[string]string {
	return map_RequestHeaderIdentityProvider
}

var map_SessionConfig = map[string]string{
	"":                     "SessionConfig specifies options for cookie-based sessions. Used by AuthRequestHandlerSession",
	"sessionSecretsFile":   "sessionSecretsFile is a reference to a file containing a serialized SessionSecrets object If no file is specified, a random signing and encryption key are generated at each server start",
	"sessionMaxAgeSeconds": "sessionMaxAgeSeconds specifies how long created sessions last. Used by AuthRequestHandlerSession",
	"sessionName":          "sessionName is the cookie name used to store the session",
}

func (SessionConfig) SwaggerDoc() map[string]string {
	return map_SessionConfig
}

var map_SessionSecret = map[string]string{
	"":               "SessionSecret is a secret used to authenticate/decrypt cookie-based sessions",
	"authentication": "Authentication is used to authenticate sessions using HMAC. Recommended to use a secret with 32 or 64 bytes.",
	"encryption":     "Encryption is used to encrypt sessions. Must be 16, 24, or 32 characters long, to select AES-128, AES-",
}

func (SessionSecret) SwaggerDoc() map[string]string {
	return map_SessionSecret
}

var map_SessionSecrets = map[string]string{
	"":        "SessionSecrets list the secrets to use to sign/encrypt and authenticate/decrypt created sessions.\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"secrets": "Secrets is a list of secrets New sessions are signed and encrypted using the first secret. Existing sessions are decrypted/authenticated by each secret until one succeeds. This allows rotating secrets.",
}

func (SessionSecrets) SwaggerDoc() map[string]string {
	return map_SessionSecrets
}

var map_TokenConfig = map[string]string{
	"":                                    "TokenConfig holds the necessary configuration options for authorization and access tokens",
	"authorizeTokenMaxAgeSeconds":         "authorizeTokenMaxAgeSeconds defines the maximum age of authorize tokens",
	"accessTokenMaxAgeSeconds":            "accessTokenMaxAgeSeconds defines the maximum age of access tokens",
	"accessTokenInactivityTimeoutSeconds": "accessTokenInactivityTimeoutSeconds - DEPRECATED: setting this field has no effect.",
	"accessTokenInactivityTimeout":        "accessTokenInactivityTimeout defines the token inactivity timeout for tokens granted by any client. The value represents the maximum amount of time that can occur between consecutive uses of the token. Tokens become invalid if they are not used within this temporal window. The user will need to acquire a new token to regain access once a token times out. Takes valid time duration string such as \"5m\", \"1.5h\" or \"2h45m\". The minimum allowed value for duration is 300s (5 minutes). If the timeout is configured per client, then that value takes precedence. If the timeout value is not specified and the client does not override the value, then tokens are valid until their lifetime.",
}

func (TokenConfig) SwaggerDoc() map[string]string {
	return map_TokenConfig
}

// AUTO-GENERATED FUNCTIONS END HERE
