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

type Book struct {
	Error             string       `json:"error"`
	Publishers        []string     `json:"publishers"`
	NumberOfPages     int          `json:"number_of_pages"`
	Series            []string     `json:"series"`
	Pagination        string       `json:"pagination"`
	LcClassifications []string     `json:"lc_classifications"`
	Key               string       `json:"key"`
	Authors           []Authors    `json:"authors"`
	PublishPlaces     []string     `json:"publish_places"`
	Contributions     []string     `json:"contributions"`
	Isbn13            []string     `json:"isbn_13"`
	Genres            []string     `json:"genres"`
	SourceRecords     []string     `json:"source_records"`
	Title             string       `json:"title"`
	DeweyDecimalClass []string     `json:"dewey_decimal_class"`
	Notes             Notes        `json:"notes"`
	Identifiers       Identifiers  `json:"identifiers"`
	Languages         []Languages  `json:"languages"`
	Lccn              []string     `json:"lccn"`
	Subjects          []string     `json:"subjects"`
	PublishDate       string       `json:"publish_date"`
	PublishCountry    string       `json:"publish_country"`
	ByStatement       string       `json:"by_statement"`
	Works             []Works      `json:"works"`
	Type              Type         `json:"type"`
	Covers            []int        `json:"covers"`
	Ocaid             string       `json:"ocaid"`
	LocalID           []string     `json:"local_id"`
	LatestRevision    int          `json:"latest_revision"`
	Revision          int          `json:"revision"`
	Created           Created      `json:"created"`
	LastModified      LastModified `json:"last_modified"`
}
type Notes struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Identifiers struct {
	Goodreads    []string `json:"goodreads"`
	Librarything []string `json:"librarything"`
}
type Languages struct {
	Key string `json:"key"`
}
type Type struct {
	Key string `json:"key"`
}
type Created struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type LastModified struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
