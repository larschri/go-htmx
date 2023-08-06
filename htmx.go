package htmx

import "net/http"

// RequestHeader is one of the htmx request headers listed in
// https://htmx.org/reference/#request_headers
type RequestHeader string

// BoolRequestHeader is one of the htmx request headers listed in
// https://htmx.org/reference/#request_headers
type BoolRequestHeader RequestHeader

// ResponseHeader is one of the htmx response headers listed in
// https://htmx.org/reference/#response_headers
type ResponseHeader string

const (
	// HXBoosted indicates that the request is via an element using hx-boost
	HXBoosted BoolRequestHeader = "HX-Boosted"

	// HXCurrentURL is the current URL of the browser
	HXCurrentURL RequestHeader = "HX-Current-URL"

	// HXHistoryResoreRequest is true if the request is for history
	// restoration after a miss in the local history cache
	HXHistoryResoreRequest BoolRequestHeader = "HX-History-Restore-Request"

	// HXPrompt the user response to an hx-prompt
	HXPrompt RequestHeader = "HX-Prompt"

	// HXRequest is always true for requests from htmx
	HXRequest BoolRequestHeader = "HX-Request"

	// HXTarget is the id of the target element if it exists
	HXTarget RequestHeader = "HX-Target"

	// HXTriggerName is the name of the triggered element if it exists
	HXTriggerName RequestHeader = "HX-Trigger-Name"

	// HXTriggerID is the id of the triggered element if it exists
	HXTriggerID RequestHeader = "HX-Trigger"

	// HXLocation allows you to do a client-side redirect that does not do a full page reload
	HXLocation ResponseHeader = "HX-Location"

	// HXPushURL pushes a new url into the history stack
	HXPushURL ResponseHeader = "HX-Push-Url"

	// HXRedirect can be used to do a client-side redirect to a new location
	HXRedirect ResponseHeader = "HX-Redirect"

	// HXRefresh triggers a full client side refresh if set to “true”
	HXRefresh ResponseHeader = "HX-Refresh"

	// HXReplaceURL replaces the current URL in the location bar
	HXReplaceURL ResponseHeader = "HX-Replace-Url"

	// HXReswap allows you to specify how the response will be swapped. See
	// hx-swap (https://htmx.org/attributes/hx-swap/) for possible values
	HXReswap ResponseHeader = "HX-Reswap"

	// HXRetarget is a CSS selector that updates the target of the content
	// update to a different element on the page
	HXRetarget ResponseHeader = "HX-Retarget"

	// HXReselect is a CSS selector that allows you to choose which part of
	// the response is used to be swapped in. Overrides an existing
	// hx-select on the triggering element
	HXReselect ResponseHeader = "HX-Reselect"

	// HXTrigger allows you to trigger client side events, see the
	// documentation for more info
	//
	// Note: See HXTriggerID for the "HX-Trigger" request header
	HXTrigger ResponseHeader = "HX-Trigger"

	// HXTriggerAfterSettle allows you to trigger client side events, see
	// the documentation for more info
	HXTriggerAfterSettle ResponseHeader = "HX-Trigger-After-Settle"

	// HXTriggerAfterSwap allows you to trigger client side events, see the
	// documentation for more info
	HXTriggerAfterSwap ResponseHeader = "HX-Trigger-After-Swap"
)

// Header returns the BoolRequestHeader value of false if it isn't set.
func (h BoolRequestHeader) Is(r *http.Request) bool {
	return r.Header.Get(string(h)) == "true"
}

// Header returns the RequestHeader value or an empty string if it isn't set.
func (h RequestHeader) Get(r *http.Request) string {
	return r.Header.Get(string(h))
}

// Header sets the ResponseHeader value
func (h ResponseHeader) Set(w http.ResponseWriter, value string) {
	w.Header().Set(string(h), value)
}

// Redirect replies to the request with a redirect url. The "HX-Redirect"
// header is used for htmx reuests. Other requests will get a plain http
// redirect. See http.Redirect.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if !HXRequest.Is(r) {
		http.Redirect(w, r, url, code)
		return
	}

	HXRedirect.Set(w, url)
}
