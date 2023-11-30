package model

type GetBooksRequest struct {
	Subject string
	Page    int `json:"page"`
	Limit   int `json:"limit"`
}

type GetBooksResponse struct {
	Books       []*Books `json:"books"`
	TotalPages  int      `json:"total_pages"`
	CurrentPage int      `json:"current_page"`
	Limit       int      `json:"limit"`
}

type Books struct {
	Title         string    `json:"title"`
	Authors       []Authors `json:"authors"`
	EditionNumber string    `json:"edition_number"`
}

type GetBookBySubjectResponse struct {
	Key         string  `json:"key"`
	Name        string  `json:"name"`
	SubjectType string  `json:"subject_type"`
	WorkCount   int     `json:"work_count"`
	Works       []Works `json:"works"`
	Error       string  `json:"error"`
}
type Authors struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
type Availability struct {
	Status              string `json:"status"`
	AvailableToBrowse   bool   `json:"available_to_browse"`
	AvailableToBorrow   bool   `json:"available_to_borrow"`
	AvailableToWaitlist bool   `json:"available_to_waitlist"`
	IsPrintdisabled     bool   `json:"is_printdisabled"`
	IsReadable          bool   `json:"is_readable"`
	IsLendable          bool   `json:"is_lendable"`
	IsPreviewable       bool   `json:"is_previewable"`
	Identifier          string `json:"identifier"`
	Isbn                any    `json:"isbn"`
	Oclc                any    `json:"oclc"`
	OpenlibraryWork     string `json:"openlibrary_work"`
	OpenlibraryEdition  string `json:"openlibrary_edition"`
	LastLoanDate        any    `json:"last_loan_date"`
	NumWaitlist         any    `json:"num_waitlist"`
	LastWaitlistDate    any    `json:"last_waitlist_date"`
	IsRestricted        bool   `json:"is_restricted"`
	IsBrowseable        bool   `json:"is_browseable"`
	Src                 string `json:"__src__"`
}
type Works struct {
	Key               string       `json:"key"`
	Title             string       `json:"title"`
	EditionCount      int          `json:"edition_count"`
	CoverID           int          `json:"cover_id"`
	CoverEditionKey   string       `json:"cover_edition_key"`
	Subject           []string     `json:"subject"`
	IaCollection      []string     `json:"ia_collection"`
	Lendinglibrary    bool         `json:"lendinglibrary"`
	Printdisabled     bool         `json:"printdisabled"`
	LendingEdition    string       `json:"lending_edition"`
	LendingIdentifier string       `json:"lending_identifier"`
	Authors           []Authors    `json:"authors"`
	FirstPublishYear  int          `json:"first_publish_year"`
	Ia                string       `json:"ia"`
	PublicScan        bool         `json:"public_scan"`
	HasFulltext       bool         `json:"has_fulltext"`
	Availability      Availability `json:"availability"`
}
