/*
Package browser contains a structured article browser that recognizes news
outlets and applies website-specific parsers (strategies) to a given HTML news
page, in order to return a well-formed struct of content (Post).

The implementation maintains session state so that the browser can provide
credentialed browsing to news sites that have paywalls.
*/

package browser
