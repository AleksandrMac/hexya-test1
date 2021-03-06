// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package m

import (
	"github.com/beevik/etree"
	"github.com/hexya-addons/web/domains"
	"github.com/hexya-addons/web/webtypes"
	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/operator"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/q"
)

// ImageMixinSet is an autogenerated type to handle ImageMixin objects.
type ImageMixinSet interface {
	models.RecordSet
	// ImageMixinSetHexyaFunc is a dummy function to uniquely match interfaces.
	ImageMixinSetHexyaFunc()
	// ForceLoad reloads the cache for the given fields and updates the ids of this ImageMixinSet.
	//
	// If no fields are given, all DB columns of the ImageMixin model are retrieved.
	//
	// It also returns this ImageMixinSet.
	ForceLoad(fields ...models.FieldName) ImageMixinSet
	// ID is a getter for the value of the "ID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	ID() int64
	// SetID is a setter for the value of the "ID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetID panics if the RecordSet is empty.
	SetID(value int64)
	// Image1024 is a getter for the value of the "Image1024" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Image1024() string
	// SetImage1024 is a setter for the value of the "Image1024" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetImage1024 panics if the RecordSet is empty.
	SetImage1024(value string)
	// Image128 is a getter for the value of the "Image128" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Image128() string
	// SetImage128 is a setter for the value of the "Image128" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetImage128 panics if the RecordSet is empty.
	SetImage128(value string)
	// Image1920 is a getter for the value of the "Image1920" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Image1920() string
	// SetImage1920 is a setter for the value of the "Image1920" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetImage1920 panics if the RecordSet is empty.
	SetImage1920(value string)
	// Image256 is a getter for the value of the "Image256" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Image256() string
	// SetImage256 is a setter for the value of the "Image256" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetImage256 panics if the RecordSet is empty.
	SetImage256(value string)
	// Image512 is a getter for the value of the "Image512" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Image512() string
	// SetImage512 is a setter for the value of the "Image512" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetImage512 panics if the RecordSet is empty.
	SetImage512(value string)
	// AddDomainLimitOffset adds the given domain, limit, offset
	// and order to the current RecordSet query.
	AddDomainLimitOffset(domain domains.Domain, limit int, offset int, order string) ImageMixinSet
	// AddModifiers adds the modifiers attribute nodes to given xml doc.
	AddModifiers(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// AddNameToRelations returns the given RecordData after getting the name of all 2one relation ids
	AddNamesToRelations(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData

	Aggregates(fieldNames ...models.FieldName) []ImageMixinGroupAggregateRow
	// Browse returns a new RecordSet with only the records with the given ids.
	// Note that this function is just a shorcut for Search on a list of ids.
	Browse(ids []int64) ImageMixinSet
	// BrowseOne returns a new RecordSet with only the record with the given id.
	// Note that this function is just a shorcut for Search on a given id.
	BrowseOne(id int64) ImageMixinSet

	CartesianProduct(others ...ImageMixinSet) []ImageMixinSet
	// CheckAccessRights verifies that the operation given by "operation" is allowed for
	// the current user according to the access rights.
	//
	// operation must be one of "read", "create", "unlink", "write".
	CheckAccessRights(args webtypes.CheckAccessRightsArgs) bool
	// CheckExecutionPermission panics if the current user is not allowed to execute the given method.
	//
	// If dontPanic is false, this function will panic, otherwise it returns true
	// if the user has the execution permission and false otherwise.
	CheckExecutionPermission(method *models.Method, dontPanic ...bool) bool
	// CheckRecursion verifies that there is no loop in a hierarchical structure of records,
	// by following the parent relationship using the 'Parent' field until a loop is detected or
	// until a top-level record is found.
	//
	// It returns true if no loop was found, false otherwise`,
	CheckRecursion() bool
	// ComputeImages computes and store resized images
	ComputeImages() ImageMixinData

	Copy(overrides ImageMixinData) ImageMixinSet

	CopyData(overrides ImageMixinData) ImageMixinData

	Create(data ImageMixinData) ImageMixinSet

	DefaultGet() ImageMixinData
	// Enqueue queues the execution of the given method with the given arguments on this recordset.
	// description will be the name given to the job.
	Enqueue(description string, method models.Methoder, arguments ...interface{}) QueueJobSet
	// Equals returns true if this RecordSet is the same as other
	// i.e. they are of the same model and have the same ids
	Equals(other ImageMixinSet) bool
	// ExecuteO2MActions executes the actions on one2many fields given by
	// the list of triplets received from the client
	ExecuteO2MActions(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Fetch query the database with the current filter and returns a RecordSet
	// with the queries ids.
	//
	// Fetch is lazy and only return ids. Use Load() instead if you want to fetch all fields.
	Fetch() ImageMixinSet
	// FieldGet returns the definition of the given field.
	// The string, help, and selection (if present) attributes are translated.
	FieldGet(field models.FieldName) *models.FieldInfo
	// FieldsGet returns the definition of each field.
	// The embedded fields are included.
	// The string, help, and selection (if present) attributes are translated.
	//
	// The result map is indexed by the fields JSON names.
	FieldsGet(args models.FieldsGetArgs) map[string]*models.FieldInfo
	// FieldsViewGet is the base implementation of the 'FieldsViewGet' method which
	// gets the detailed composition of the requested view like fields, mixin,
	// view architecture.
	FieldsViewGet(args webtypes.FieldsViewGetParams) *webtypes.FieldsViewData

	Filtered(test func(ImageMixinSet) bool) ImageMixinSet
	// FormatRelationFields returns the given data with all relation fields converted to int64 or []int64
	FormatRelationFields(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData
	// GetFormviewAction returns an action to open the document.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.`,
	GetFormviewAction() *actions.Action
	// GetFormviewID returns an view id to open the document with.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.
	GetFormviewId() string
	// GetRecord returns the Recordset with the given externalID. It panics if the externalID does not exist.
	GetRecord(externalID string) ImageMixinSet
	// GetToolbar returns a toolbar populated with the actions linked to this model
	GetToolbar() webtypes.Toolbar
	// GroupBy returns a new RecordSet grouped with the given GROUP BY expressions.
	GroupBy(exprs ...models.FieldName) ImageMixinSet
	// Intersect returns a new RecordCollection with only the records that are both
	// in this RecordCollection and in the other RecordSet.
	Intersect(other ImageMixinSet) ImageMixinSet
	// Limit returns a new RecordSet with only the first 'limit' records.
	Limit(limit int) ImageMixinSet
	// Load looks up cache for fields of the RecordCollection and
	// query database for missing values.
	// fields are the fields to retrieve in the expression format,
	// i.e. "User.Profile.Age" or "user_id.profile_id.age".
	// If no fields are given, all DB columns of the RecordCollection's
	// model are retrieved.
	Load(fields ...models.FieldName) ImageMixinSet
	// LoadViews returns the data for all the views and filters required in the parameters.
	LoadViews(args webtypes.LoadViewsArgs) *webtypes.LoadViewsData
	// ManageGroupsOnFields adds the invisible attribute to fields nodes if the current
	// user does not belong to one of the groups of the 'groups' attribute
	ManageGroupsOnFields(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// NameGet retrieves the human readable name of this record.`,
	NameGet() string
	// NameSearch searches for records that have a display name matching the given
	// "name" pattern when compared with the given "operator", while also
	// matching the optional search domain ("args").
	//
	// This is used for example to provide suggestions based on a partial
	// value for a relational field. Sometimes be seen as the inverse
	// function of NameGet but it is not guaranteed to be.
	NameSearch(params webtypes.NameSearchParams) []webtypes.RecordIDWithName

	New(data ImageMixinData) ImageMixinSet
	// NormalizeM2MData converts the list of triplets received from the client into the final list of ids
	// to keep in the Many2Many relationship of this model through the given field.
	NormalizeM2MData(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Offset returns a new RecordSet with only the records starting at offset
	Offset(offset int) ImageMixinSet
	// Onchange returns the values that must be modified according to each field's Onchange
	// method in the pseudo-record given as params.Values`,
	Onchange(params models.OnchangeParams) models.OnchangeResult
	// OrderBy returns a new RecordSet ordered by the given ORDER BY expressions.
	// Each expression contains a field name and optionally one of "asc" or "desc", such as:
	//
	// rs.OrderBy("Company", "Name desc")
	OrderBy(exprs ...string) ImageMixinSet
	// PostProcessCreateValues updates FK of related records created at the same time.
	//
	// This method is meant to be called with the second returned value of ProcessCreateValues
	// after record creation.
	PostProcessCreateValues(data models.RecordData)
	// PostProcessFilters transforms a map[models.FieldName]models.Conditioner
	// in a map[string][]interface{} which acts as a map of domains.
	PostProcessFilters(in map[models.FieldName]models.Conditioner) map[string][]interface{}
	// ProcessCreateValues updates the given data values for Create method to be
	// compatible with the ORM, in particular for relation fields.
	//
	// It returns a first FieldMap to be used as argument to the Create method, and
	// a second map to be used with a subsequent call to PostProcessCreateValues (for
	// updating FKs pointing to the newly created record).
	ProcessCreateValues(data models.RecordData) (models.RecordData, models.RecordData)
	// ProcessElementAttrs returns a modifiers map according to the domain
	// in attrs of the given element
	ProcessElementAttrs(element *etree.Element, fieldInfos map[string]*models.FieldInfo) map[string]interface{}
	// ProcessFieldElementModifiers modifies the given modifiers map by taking into account:
	// - 'invisible', 'readonly' and 'required' attributes in field tags
	// - 'ReadOnly' and 'Required' parameters of the model's field'
	// It returns the modified map.
	ProcessFieldElementModifiers(element *etree.Element, fieldInfos map[string]*models.FieldInfo, modifiers map[string]interface{}) map[string]interface{}
	// ProcessView makes all the necessary modifications to the view
	// arch and returns the new xml string.`,
	ProcessView(arch *etree.Document, fieldInfos map[string]*models.FieldInfo) string
	// ProcessWriteValues updates the given data values for Write method to be
	// compatible with the ORM, in particular for relation fields
	ProcessWriteValues(data models.RecordData) models.RecordData
	// Read reads the database and returns a slice of FieldMap of the given model.
	Read(fields models.FieldNames) []models.RecordData
	// ReadGroup gets a list of record aggregates according to the given parameters.
	ReadGroup(params webtypes.ReadGroupParams) []models.FieldMap
	// SQLFromCondition returns the WHERE clause sql and arguments corresponding to
	// the given condition.`,
	SQLFromCondition(c *models.Condition) (string, models.SQLParams)

	Search(condition q.ImageMixinCondition) ImageMixinSet
	// SearchAll returns a RecordSet with all items of the table, regardless of the
	// current RecordSet query. It is mainly meant to be used on an empty RecordSet.
	SearchAll() ImageMixinSet

	SearchByName(name string, op operator.Operator, additionalCond q.ImageMixinCondition, limit int) ImageMixinSet
	// SearchCount fetch from the database the number of records that match the RecordSet conditions.
	SearchCount() int
	// SearchDomain execute a search on the given domain.
	SearchDomain(domain domains.Domain) CommonMixinSet
	// SearchRead retrieves database records according to the filters defined in params.
	SearchRead(params webtypes.SearchParams) []models.RecordData

	Sorted(less func(ImageMixinSet, ImageMixinSet) bool) ImageMixinSet
	// SortedByField returns a new record set with the same records as rc but sorted by the given field.
	// If reverse is true, the sort is done in reversed order
	SortedByField(namer models.FieldName, reverse bool) ImageMixinSet
	// SortedDefault returns a new record set with the same records as rc but sorted according
	// to the default order of this model
	SortedDefault() ImageMixinSet
	// Subtract returns a RecordSet with the Records that are in this
	// RecordCollection but not in the given 'other' one.
	// The result is guaranteed to be a set of unique records.
	Subtract(other ImageMixinSet) ImageMixinSet
	// Sudo returns a new RecordSet with the given userID
	// or the superuser ID if not specified
	Sudo(userID ...int64) ImageMixinSet
	// Union returns a new RecordSet that is the union of this RecordSet and the given
	// "other" RecordSet. The result is guaranteed to be a set of unique records.
	Union(other ImageMixinSet) ImageMixinSet
	// Unlink deletes the given records in the database.
	Unlink() int64
	// WebReadGroup returns the result of a read_group (and optionally search for and read records inside each
	// group), and the total number of groups matching the search domain.
	WebReadGroup(params webtypes.WebReadGroupParams) webtypes.WebReadGroupResult
	// WebReadGroupPrivate performs a read_group and optionally a web_search_read for each group.
	WebReadGroupPrivate(params webtypes.WebReadGroupParams) []models.FieldMap
	// WebSearchRead performs a search_read and a search_count.
	WebSearchRead(params webtypes.SearchParams) webtypes.SearchReadResult
	// WithContext returns a copy of the current RecordSet with
	// its context extended by the given key and value.
	WithContext(key string, value interface{}) ImageMixinSet
	// WithEnv returns a copy of the current RecordSet with the given Environment.
	WithEnv(env models.Environment) ImageMixinSet
	// WithNewContext returns a copy of the current RecordSet with its context
	// replaced by the given one.
	WithNewContext(context *types.Context) ImageMixinSet

	Write(data ImageMixinData) bool
	// Super returns a RecordSet with a modified callstack so that call to the current
	// method will execute the next method layer.
	//
	// This method is meant to be used inside a method layer function to call its parent,
	// such as:
	//
	//    func (rs h.MyRecordSet) MyMethod() string {
	//        res := rs.Super().MyMethod()
	//        res += " ok!"
	//        return res
	//    }
	//
	// Calls to a different method than the current method will call its next layer only
	// if the current method has been called from a layer of the other method. Otherwise,
	// it will be the same as calling the other method directly.
	Super() ImageMixinSet
	// ModelData returns a new ImageMixinData object populated with the values
	// of the given FieldMap.
	ModelData(fMap models.FieldMap) ImageMixinData
	// Records returns a slice with all the records of this RecordSet, as singleton RecordSets
	Records() []ImageMixinSet
	// First returns the values of the first Record of the RecordSet as a pointer to a ImageMixinData.
	//
	// If this RecordSet is empty, it returns an empty ImageMixinData.
	First() ImageMixinData
	// All returns the values of all Records of the RecordCollection as a slice of ImageMixinData pointers.
	All() []ImageMixinData
}

// ImageMixinData is used to hold values of an ImageMixin object instance
// when creating or updating a ImageMixinSet.
type ImageMixinData interface {
	models.RecordData
	// Get returns the value of the given field.
	//
	// The field can be either its name or is JSON name.
	Get(field models.FieldName) interface{}
	// Has returns true if a value is set for the given field.
	//
	// The field can be either its name or is JSON name.
	Has(field models.FieldName) bool
	// Set sets the given field with the given value.
	// If the field already exists, then it is updated with value.
	// Otherwise, a new entry is inserted.
	//
	// It returns the given ImageMixinData so that calls can be chained
	Set(field models.FieldName, value interface{}) ImageMixinData
	// Unset removes the value of the given field if it exists.
	//
	// It returns the given ModelData so that calls can be chained
	Unset(field models.FieldName) ImageMixinData
	// Copy returns a copy of this ImageMixinData
	Copy() ImageMixinData
	// MergeWith updates this ImageMixinData with the given other ImageMixinData
	// If a field of the other ImageMixinData already exists here, the value is overridden,
	// otherwise, the field is inserted.
	MergeWith(other ImageMixinData)
	// Keys returns the ImageMixinData keys as a slice of strings
	Keys() (res []string)
	// OrderedKeys returns the keys of this ImageMixinData ordered.
	//
	// This has the convenient side effect of having shorter paths come before longer paths,
	// which is particularly useful when creating or updating related records.
	OrderedKeys() []string
	// FieldNames returns the ImageMixinData keys as a slice of FieldNames.
	FieldNames() models.FieldNames
	// ID returns the value of the ID field.
	// If this ID is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	ID() int64
	// HasID returns true if ID is set in this ImageMixinData
	HasID() bool
	// SetID sets the ID field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetID(value int64) ImageMixinData
	// UnsetID removes the value of the ID field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetID() ImageMixinData

	// Image1024 returns the value of the Image1024 field.
	// If this Image1024 is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	Image1024() string
	// HasImage1024 returns true if Image1024 is set in this ImageMixinData
	HasImage1024() bool
	// SetImage1024 sets the Image1024 field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetImage1024(value string) ImageMixinData
	// UnsetImage1024 removes the value of the Image1024 field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetImage1024() ImageMixinData

	// Image128 returns the value of the Image128 field.
	// If this Image128 is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	Image128() string
	// HasImage128 returns true if Image128 is set in this ImageMixinData
	HasImage128() bool
	// SetImage128 sets the Image128 field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetImage128(value string) ImageMixinData
	// UnsetImage128 removes the value of the Image128 field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetImage128() ImageMixinData

	// Image1920 returns the value of the Image1920 field.
	// If this Image1920 is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	Image1920() string
	// HasImage1920 returns true if Image1920 is set in this ImageMixinData
	HasImage1920() bool
	// SetImage1920 sets the Image1920 field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetImage1920(value string) ImageMixinData
	// UnsetImage1920 removes the value of the Image1920 field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetImage1920() ImageMixinData

	// Image256 returns the value of the Image256 field.
	// If this Image256 is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	Image256() string
	// HasImage256 returns true if Image256 is set in this ImageMixinData
	HasImage256() bool
	// SetImage256 sets the Image256 field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetImage256(value string) ImageMixinData
	// UnsetImage256 removes the value of the Image256 field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetImage256() ImageMixinData

	// Image512 returns the value of the Image512 field.
	// If this Image512 is not set in this ImageMixinData, then
	// the Go zero value for the type is returned.
	Image512() string
	// HasImage512 returns true if Image512 is set in this ImageMixinData
	HasImage512() bool
	// SetImage512 sets the Image512 field with the given value.
	// It returns this ImageMixinData so that calls can be chained.
	SetImage512(value string) ImageMixinData
	// UnsetImage512 removes the value of the Image512 field if it exists.
	// It returns this ImageMixinData so that calls can be chained.
	UnsetImage512() ImageMixinData
}

// A ImageMixinGroupAggregateRow holds a row of results of a query with a group by clause
type ImageMixinGroupAggregateRow interface {
	// Values() returns the values of the actual query
	Values() ImageMixinData
	// Count is the number of lines aggregated into this one
	Count() int
	// Condition can be used to query the aggregated rows separately if needed
	Condition() q.ImageMixinCondition
}
