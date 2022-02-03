package schema

// BackMatter A collection of resources, which may be included directly or by reference.
type BackMatter struct {
	Resources []*Resource `json:"resources,omitempty"`
}

// Base64 The Base64 alphabet in RFC 2045 - aligned with XSD.
type Base64 struct {

	// Name of the file before it was encoded as Base64 to be embedded in a resource. This is the name that will be assigned to the file when the file is decoded.
	Filename string `json:"filename,omitempty"`

	// Specifies a media type as defined by the Internet Assigned Numbers Authority (IANA) Media Types Registry.
	MediaType string `json:"media-type,omitempty"`
	Value     string `json:"value"`
}

// Citation A citation consisting of end note text and optional structured bibliographic data.
type Citation struct {
	Links []*Link     `json:"links,omitempty"`
	Props []*Property `json:"props,omitempty"`

	// A line of citation text.
	Text string `json:"text"`
}

// Constraint A formal or informal expression of a constraint or test
type Constraint struct {

	// A textual summary of the constraint to be applied.
	Description string            `json:"description,omitempty"`
	Tests       []*ConstraintTest `json:"tests,omitempty"`
}

// ConstraintTest A test expression which is expected to be evaluated by a tool.
type ConstraintTest struct {

	// A formal (executable) expression of a constraint
	Expression string `json:"expression"`
	Remarks    string `json:"remarks,omitempty"`
}

// Control A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
type Control struct {

	// A textual label that provides a sub-type or characterization of the control.
	Class    string     `json:"class,omitempty"`
	Controls []*Control `json:"controls,omitempty"`

	// A human-oriented, locally unique identifier with instance scope that can be used to reference this control elsewhere in this and other OSCAL instances (e.g., profiles). This id should be assigned per-subject, which means it should be consistently used to identify the same control across revisions of the document.
	Id     string       `json:"id"`
	Links  []*Link      `json:"links,omitempty"`
	Params []*Parameter `json:"params,omitempty"`
	Parts  []*Part      `json:"parts,omitempty"`
	Props  []*Property  `json:"props,omitempty"`

	// A name given to the control, which may be used by a tool for display and navigation.
	Title string `json:"title"`
}

// ControlGroup A group of controls, or of groups of controls.
type ControlGroup struct {

	// A textual label that provides a sub-type or characterization of the group.
	Class    string          `json:"class,omitempty"`
	Controls []*Control      `json:"controls,omitempty"`
	Groups   []*ControlGroup `json:"groups,omitempty"`

	// A human-oriented, locally unique identifier with cross-instance scope that can be used to reference this defined group elsewhere in in this and other OSCAL instances (e.g., profiles). This id should be assigned per-subject, which means it should be consistently used to identify the same group across revisions of the document.
	Id     string       `json:"id,omitempty"`
	Links  []*Link      `json:"links,omitempty"`
	Params []*Parameter `json:"params,omitempty"`
	Parts  []*Part      `json:"parts,omitempty"`
	Props  []*Property  `json:"props,omitempty"`

	// A name given to the group, which may be used by a tool for display and navigation.
	Title string `json:"title"`
}

// Guideline A prose statement that provides a recommendation for the use of a parameter.
type Guideline struct {

	// Prose permits multiple paragraphs, lists, tables etc.
	Prose string `json:"prose"`
}

// Hash A representation of a cryptographic digest generated over a resource using a specified hash algorithm.
type Hash struct {

	// Method by which a hash is derived
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

// Link A reference to a local or remote resource
type Link struct {

	// A resolvable URL reference to a resource.
	Href string `json:"href"`

	// Specifies a media type as defined by the Internet Assigned Numbers Authority (IANA) Media Types Registry.
	MediaType string `json:"media-type,omitempty"`

	// Describes the type of relationship provided by the link. This can be an indicator of the link's purpose.
	Rel string `json:"rel,omitempty"`

	// A textual label to associate with the link, which may be used for presentation in a tool.
	Text string `json:"text,omitempty"`
}

// Location A location, with associated metadata that can be referenced.
type Location struct {
	Address          *Address           `json:"address"`
	EmailAddresses   []string           `json:"email-addresses,omitempty"`
	Links            []*Link            `json:"links,omitempty"`
	Props            []*Property        `json:"props,omitempty"`
	Remarks          string             `json:"remarks,omitempty"`
	TelephoneNumbers []*TelephoneNumber `json:"telephone-numbers,omitempty"`

	// A name given to the location, which may be used by a tool for display and navigation.
	Title string   `json:"title,omitempty"`
	Urls  []string `json:"urls,omitempty"`

	// A machine-oriented, globally unique identifier with cross-instance scope that can be used to reference this defined location elsewhere in this or other OSCAL instances. The locally defined UUID of the location can be used to reference the data item locally or globally (e.g., from an importing OSCAL instance). This UUID should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Uuid string `json:"uuid"`
}

// Catalog A collection of controls.
type Catalog struct {
	BackMatter *BackMatter     `json:"back-matter,omitempty"`
	Controls   []*Control      `json:"controls,omitempty"`
	Groups     []*ControlGroup `json:"groups,omitempty"`
	Metadata   *Metadata       `json:"metadata"`
	Params     []*Parameter    `json:"params,omitempty"`

	// A globally unique identifier with cross-instance scope for this catalog instance. This UUID should be changed when this document is revised.
	Uuid string `json:"uuid"`
}

// Part A partition of a control's definition or a child of another part.
type Part struct {

	// A textual label that provides a sub-type or characterization of the part's name. This can be used to further distinguish or discriminate between the semantics of multiple parts of the same control with the same name and ns.
	Class string `json:"class,omitempty"`

	// A human-oriented, locally unique identifier with cross-instance scope that can be used to reference this defined part elsewhere in this or other OSCAL instances. When referenced from another OSCAL instance, this identifier must be referenced in the context of the containing resource (e.g., import-profile). This id should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Id    string  `json:"id,omitempty"`
	Links []*Link `json:"links,omitempty"`

	// A textual label that uniquely identifies the part's semantic type.
	Name string `json:"name"`

	// A namespace qualifying the part's name. This allows different organizations to associate distinct semantics with the same name.
	Ns    string      `json:"ns,omitempty"`
	Parts []*Part     `json:"parts,omitempty"`
	Props []*Property `json:"props,omitempty"`

	// Permits multiple paragraphs, lists, tables etc.
	Prose string `json:"prose,omitempty"`

	// A name given to the part, which may be used by a tool for display and navigation.
	Title string `json:"title,omitempty"`
}

// Address A postal address for the location.
type Address struct {
	AddrLines []string `json:"addr-lines,omitempty"`

	// City, town or geographical region for the mailing address.
	City string `json:"city,omitempty"`

	// The ISO 3166-1 alpha-2 country code for the mailing address.
	Country string `json:"country,omitempty"`

	// Postal or ZIP code for mailing address
	PostalCode string `json:"postal-code,omitempty"`

	// State, province or analogous geographical region for mailing address
	State string `json:"state,omitempty"`

	// Indicates the type of address.
	Type string `json:"type,omitempty"`
}

// DocumentId A document identifier qualified by an identifier scheme. A document identifier provides a globally unique identifier with a cross-instance scope that is used for a group of documents that are to be treated as different versions of the same document. If this element does not appear, or if the value of this element is empty, the value of "document-id" is equal to the value of the "uuid" flag of the top-level root element.
type DocumentId struct {
	Identifier string `json:"identifier"`

	// Qualifies the kind of document identifier using a URI. If the scheme is not provided the value of the element will be interpreted as a string of characters.
	Scheme string `json:"scheme,omitempty"`
}

// Metadata Provides information about the publication and availability of the containing document.
type Metadata struct {
	DocumentIds        []*DocumentId           `json:"document-ids,omitempty"`
	LastModified       string                  `json:"last-modified"`
	Links              []*Link                 `json:"links,omitempty"`
	Locations          []*Location             `json:"locations,omitempty"`
	OscalVersion       string                  `json:"oscal-version"`
	Parties            []*Party                `json:"parties,omitempty"`
	Props              []*Property             `json:"props,omitempty"`
	Published          string                  `json:"published,omitempty"`
	Remarks            string                  `json:"remarks,omitempty"`
	ResponsibleParties []*ResponsibleParty     `json:"responsible-parties,omitempty"`
	Revisions          []*RevisionHistoryEntry `json:"revisions,omitempty"`
	Roles              []*MetadataRole         `json:"roles,omitempty"`

	// A name given to the document, which may be used by a tool for display and navigation.
	Title   string `json:"title"`
	Version string `json:"version"`
}

// Party A responsible entity which is either a person or an organization.
type Party struct {
	Addresses             []*Address                 `json:"addresses,omitempty"`
	EmailAddresses        []string                   `json:"email-addresses,omitempty"`
	ExternalIds           []*PartyExternalIdentifier `json:"external-ids,omitempty"`
	Links                 []*Link                    `json:"links,omitempty"`
	LocationUuids         []string                   `json:"location-uuids,omitempty"`
	MemberOfOrganizations []string                   `json:"member-of-organizations,omitempty"`

	// The full name of the party. This is typically the legal name associated with the party.
	Name    string      `json:"name,omitempty"`
	Props   []*Property `json:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty"`

	// A short common name, abbreviation, or acronym for the party.
	ShortName        string             `json:"short-name,omitempty"`
	TelephoneNumbers []*TelephoneNumber `json:"telephone-numbers,omitempty"`

	// A category describing the kind of party the object describes.
	Type string `json:"type"`

	// A machine-oriented, globally unique identifier with cross-instance scope that can be used to reference this defined party elsewhere in this or other OSCAL instances. The locally defined UUID of the party can be used to reference the data item locally or globally (e.g., from an importing OSCAL instance). This UUID should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Uuid string `json:"uuid"`
}

// MetadataRole Defines a function assumed or expected to be assumed by a party in a specific situation.
type MetadataRole struct {

	// A summary of the role's purpose and associated responsibilities.
	Description string `json:"description,omitempty"`

	// A human-oriented, locally unique identifier with cross-instance scope that can be used to reference this defined role elsewhere in this or other OSCAL instances. When referenced from another OSCAL instance, the locally defined ID of the Role from the imported OSCAL instance must be referenced in the context of the containing resource (e.g., import, import-component-definition, import-profile, import-ssp or import-ap). This ID should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Id      string      `json:"id"`
	Links   []*Link     `json:"links,omitempty"`
	Props   []*Property `json:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty"`

	// A short common name, abbreviation, or acronym for the role.
	ShortName string `json:"short-name,omitempty"`

	// A name given to the role, which may be used by a tool for display and navigation.
	Title string `json:"title"`
}

// Parameter Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
type Parameter struct {

	// A textual label that provides a characterization of the parameter.
	Class       string        `json:"class,omitempty"`
	Constraints []*Constraint `json:"constraints,omitempty"`

	// **(deprecated)** Another parameter invoking this one. This construct has been deprecated and should not be used.
	DependsOn  string       `json:"depends-on,omitempty"`
	Guidelines []*Guideline `json:"guidelines,omitempty"`

	// A human-oriented, locally unique identifier with cross-instance scope that can be used to reference this defined parameter elsewhere in this or other OSCAL instances. When referenced from another OSCAL instance, this identifier must be referenced in the context of the containing resource (e.g., import-profile). This id should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Id string `json:"id"`

	// A short, placeholder name for the parameter, which can be used as a substitute for a value if no value is assigned.
	Label   string      `json:"label,omitempty"`
	Links   []*Link     `json:"links,omitempty"`
	Props   []*Property `json:"props,omitempty"`
	Remarks string      `json:"remarks,omitempty"`
	Select  *Selection  `json:"select,omitempty"`

	// Describes the purpose and use of a parameter
	Usage  string   `json:"usage,omitempty"`
	Values []string `json:"values,omitempty"`
}

// PartyExternalIdentifier An identifier for a person or organization using a designated scheme. e.g. an Open Researcher and Contributor ID (ORCID)
type PartyExternalIdentifier struct {
	Id string `json:"id"`

	// Indicates the type of external identifier.
	Scheme string `json:"scheme"`
}

// Property An attribute, characteristic, or quality of the containing object expressed as a namespace qualified name/value pair. The value of a property is a simple scalar value, which may be expressed as a list of values.
type Property struct {

	// A textual label that provides a sub-type or characterization of the property's name. This can be used to further distinguish or discriminate between the semantics of multiple properties of the same object with the same name and ns.
	Class string `json:"class,omitempty"`

	// A textual label that uniquely identifies a specific attribute, characteristic, or quality of the property's containing object.
	Name string `json:"name"`

	// A namespace qualifying the property's name. This allows different organizations to associate distinct semantics with the same name.
	Ns      string `json:"ns,omitempty"`
	Remarks string `json:"remarks,omitempty"`

	// A machine-oriented, globally unique identifier with cross-instance scope that can be used to reference this defined property elsewhere in this or other OSCAL instances. This UUID should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Uuid string `json:"uuid,omitempty"`

	// Indicates the value of the attribute, characteristic, or quality.
	Value string `json:"value"`
}

// Resource A resource associated with content in the containing document. A resource may be directly included in the document base64 encoded or may point to one or more equivalent internet resources.
type Resource struct {

	// The Base64 alphabet in RFC 2045 - aligned with XSD.
	Base64 *Base64 `json:"base64,omitempty"`

	// A citation consisting of end note text and optional structured bibliographic data.
	Citation *Citation `json:"citation,omitempty"`

	// A short summary of the resource used to indicate the purpose of the resource.
	Description string          `json:"description,omitempty"`
	DocumentIds []*DocumentId   `json:"document-ids,omitempty"`
	Props       []*Property     `json:"props,omitempty"`
	Remarks     string          `json:"remarks,omitempty"`
	Rlinks      []*ResourceLink `json:"rlinks,omitempty"`

	// A name given to the resource, which may be used by a tool for display and navigation.
	Title string `json:"title,omitempty"`

	// A machine-oriented, globally unique identifier with cross-instance scope that can be used to reference this defined resource elsewhere in this or other OSCAL instances. This UUID should be assigned per-subject, which means it should be consistently used to identify the same subject across revisions of the document.
	Uuid string `json:"uuid"`
}

// ResourceLink A pointer to an external resource with an optional hash for verification and change detection.
type ResourceLink struct {
	Hashes []*Hash `json:"hashes,omitempty"`

	// A resolvable URI reference to a resource.
	Href string `json:"href"`

	// Specifies a media type as defined by the Internet Assigned Numbers Authority (IANA) Media Types Registry.
	MediaType string `json:"media-type,omitempty"`
}

// ResponsibleParty A reference to a set of organizations or persons that have responsibility for performing a referenced role in the context of the containing object.
type ResponsibleParty struct {
	Links      []*Link     `json:"links,omitempty"`
	PartyUuids []string    `json:"party-uuids"`
	Props      []*Property `json:"props,omitempty"`
	Remarks    string      `json:"remarks,omitempty"`

	// A human-oriented identifier reference to roles served by the user.
	RoleId string `json:"role-id"`
}

// RevisionHistoryEntry An entry in a sequential list of revisions to the containing document in reverse chronological order (i.e., most recent previous revision first).
type RevisionHistoryEntry struct {
	LastModified string      `json:"last-modified,omitempty"`
	Links        []*Link     `json:"links,omitempty"`
	OscalVersion string      `json:"oscal-version,omitempty"`
	Props        []*Property `json:"props,omitempty"`
	Published    string      `json:"published,omitempty"`
	Remarks      string      `json:"remarks,omitempty"`

	// A name given to the document revision, which may be used by a tool for display and navigation.
	Title   string `json:"title,omitempty"`
	Version string `json:"version"`
}

// Root
type Root struct {
	Catalog *Catalog `json:"catalog"`
}

// Selection Presenting a choice among alternatives
type Selection struct {
	Choice []string `json:"choice,omitempty"`

	// Describes the number of selections that must occur. Without this setting, only one value should be assumed to be permitted.
	HowMany string `json:"how-many,omitempty"`
}

// TelephoneNumber Contact number by telephone.
type TelephoneNumber struct {
	Number string `json:"number"`

	// Indicates the type of phone number.
	Type string `json:"type,omitempty"`
}
