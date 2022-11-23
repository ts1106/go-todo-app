package api

import (
	"time"
)

type Response struct {
	ApiVersion string `json:"api_version,omitempty"`
	Context    string `json:"context,omitempty"`
	Id         string `json:"id,omitempty"`
	Method     string `json:"method,omitempty"`
	Params     any    `json:"params,omitempty"`
	Data       *Data  `json:"data,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

type Data struct {
	Kind    string     `json:"kind,omitempty"`
	Fields  string     `json:"fields,omitempty"`
	Etag    string     `json:"etag,omitempty"`
	Id      string     `json:"id,omitempty"`
	Lang    string     `json:"lang,omitempty"`    //formatted BCP 47
	Updated *time.Time `json:"updated,omitempty"` //formatted RFC 3339
	Deleted bool       `json:"deleted,omitempty"`
	Items   any        `json:"items,omitempty"`

	// Paging
	CurrentItemCount   uint64 `json:"current_item_count,omitempty"`
	ItemsPerPage       uint64 `json:"items_per_page,omitempty"`
	StartIndex         uint64 `json:"start_index,omitempty"`
	TotalItems         uint64 `json:"total_items,omitempty"`
	PagingLinkTemplate string `json:"paging_ling_template,omitempty"`
	PageIndex          uint64 `json:"page_index,omitempty"`
	TotalPages         uint64 `json:"total_pages,omitempty"`

	// Links
	Self         any    `json:"self,omitempty"`
	SelfLink     string `json:"self_link,omitempty"`
	Edit         any    `json:"edit,omitempty"`
	EditLink     string `json:"edit_link,omitempty"`
	Next         any    `json:"next,omitempty"`
	NextLink     string `json:"next_link,omitempty"`
	Previous     any    `json:"previous,omitempty"`
	PreviousLink string `json:"previous_link,omitempty"`
}

type Error struct {
	Code    int      `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Errors  []Errors `json:"errors,omitempty"`
}

type Errors struct {
	Domain       string `json:"domain,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Message      string `json:"message,omitempty"`
	Location     string `json:"location,omitempty"`
	LocationType string `json:"location_type,omitempty"`
	ExtendedHelp string `json:"extended_help,omitempty"`
	SendReport   string `json:"send_report,omitempty"`
}
