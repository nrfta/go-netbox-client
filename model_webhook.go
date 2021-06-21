/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type Webhook struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	// Call this webhook when a matching object is created.
	TypeCreate bool `json:"type_create,omitempty"`
	// Call this webhook when a matching object is updated.
	TypeUpdate bool `json:"type_update,omitempty"`
	// Call this webhook when a matching object is deleted.
	TypeDelete bool `json:"type_delete,omitempty"`
	// A POST will be sent to this URL when the webhook is called.
	PayloadUrl string `json:"payload_url"`
	Enabled bool `json:"enabled,omitempty"`
	HttpMethod string `json:"http_method,omitempty"`
	// The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.
	HttpContentType string `json:"http_content_type,omitempty"`
	// User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is support with the same context as the request body (below).
	AdditionalHeaders string `json:"additional_headers,omitempty"`
	// Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate string `json:"body_template,omitempty"`
	// When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	Secret string `json:"secret,omitempty"`
	// Enable SSL certificate verification. Disable with caution!
	SslVerification bool `json:"ssl_verification,omitempty"`
	// The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	CaFilePath string `json:"ca_file_path,omitempty"`
}