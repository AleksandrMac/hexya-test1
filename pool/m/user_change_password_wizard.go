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
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/pool/q"
)

// UserChangePasswordWizardSet is an autogenerated type to handle UserChangePasswordWizard objects.
type UserChangePasswordWizardSet interface {
	models.RecordSet
	// UserChangePasswordWizardSetHexyaFunc is a dummy function to uniquely match interfaces.
	UserChangePasswordWizardSetHexyaFunc()
	// ForceLoad reloads the cache for the given fields and updates the ids of this UserChangePasswordWizardSet.
	//
	// If no fields are given, all DB columns of the UserChangePasswordWizard model are retrieved.
	//
	// It also returns this UserChangePasswordWizardSet.
	ForceLoad(fields ...models.FieldName) UserChangePasswordWizardSet
	// CreateDate is a getter for the value of the "CreateDate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	CreateDate() dates.DateTime
	// SetCreateDate is a setter for the value of the "CreateDate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetCreateDate panics if the RecordSet is empty.
	SetCreateDate(value dates.DateTime)
	// CreateUID is a getter for the value of the "CreateUID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	CreateUID() int64
	// SetCreateUID is a setter for the value of the "CreateUID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetCreateUID panics if the RecordSet is empty.
	SetCreateUID(value int64)
	// DisplayName is a getter for the value of the "DisplayName" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	DisplayName() string
	// SetDisplayName is a setter for the value of the "DisplayName" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetDisplayName panics if the RecordSet is empty.
	SetDisplayName(value string)
	// ID is a getter for the value of the "ID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	ID() int64
	// SetID is a setter for the value of the "ID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetID panics if the RecordSet is empty.
	SetID(value int64)
	// LastUpdate is a getter for the value of the "LastUpdate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	LastUpdate() dates.DateTime
	// SetLastUpdate is a setter for the value of the "LastUpdate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetLastUpdate panics if the RecordSet is empty.
	SetLastUpdate(value dates.DateTime)
	// Users is a getter for the value of the "Users" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Users() UserChangePasswordWizardLineSet
	// SetUsers is a setter for the value of the "Users" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetUsers panics if the RecordSet is empty.
	SetUsers(value UserChangePasswordWizardLineSet)
	// WriteDate is a getter for the value of the "WriteDate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	WriteDate() dates.DateTime
	// SetWriteDate is a setter for the value of the "WriteDate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetWriteDate panics if the RecordSet is empty.
	SetWriteDate(value dates.DateTime)
	// WriteUID is a getter for the value of the "WriteUID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	WriteUID() int64
	// SetWriteUID is a setter for the value of the "WriteUID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetWriteUID panics if the RecordSet is empty.
	SetWriteUID(value int64)
	// AddDomainLimitOffset adds the given domain, limit, offset
	// and order to the current RecordSet query.
	AddDomainLimitOffset(domain domains.Domain, limit int, offset int, order string) UserChangePasswordWizardSet
	// AddModifiers adds the modifiers attribute nodes to given xml doc.
	AddModifiers(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// AddNameToRelations returns the given RecordData after getting the name of all 2one relation ids
	AddNamesToRelations(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData

	Aggregates(fieldNames ...models.FieldName) []UserChangePasswordWizardGroupAggregateRow
	// Browse returns a new RecordSet with only the records with the given ids.
	// Note that this function is just a shorcut for Search on a list of ids.
	Browse(ids []int64) UserChangePasswordWizardSet
	// BrowseOne returns a new RecordSet with only the record with the given id.
	// Note that this function is just a shorcut for Search on a given id.
	BrowseOne(id int64) UserChangePasswordWizardSet

	CartesianProduct(others ...UserChangePasswordWizardSet) []UserChangePasswordWizardSet
	// ChangePasswordButton is called when the user clicks on 'Apply' button in the popup.
	// It updates the user's password.`,
	ChangePasswordButton()
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
	// ComputeDisplayName updates the DisplayName field with the result of NameGet
	ComputeDisplayName() *models.ModelData
	// ComputeLastUpdate returns the last datetime at which the record has been updated.
	ComputeLastUpdate() *models.ModelData

	Copy(overrides UserChangePasswordWizardData) UserChangePasswordWizardSet

	CopyData(overrides UserChangePasswordWizardData) UserChangePasswordWizardData

	Create(data UserChangePasswordWizardData) UserChangePasswordWizardSet

	DefaultGet() UserChangePasswordWizardData
	// Enqueue queues the execution of the given method with the given arguments on this recordset.
	// description will be the name given to the job.
	Enqueue(description string, method models.Methoder, arguments ...interface{}) QueueJobSet
	// Equals returns true if this RecordSet is the same as other
	// i.e. they are of the same model and have the same ids
	Equals(other UserChangePasswordWizardSet) bool
	// ExecuteO2MActions executes the actions on one2many fields given by
	// the list of triplets received from the client
	ExecuteO2MActions(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Fetch query the database with the current filter and returns a RecordSet
	// with the queries ids.
	//
	// Fetch is lazy and only return ids. Use Load() instead if you want to fetch all fields.
	Fetch() UserChangePasswordWizardSet
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

	Filtered(test func(UserChangePasswordWizardSet) bool) UserChangePasswordWizardSet
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
	GetRecord(externalID string) UserChangePasswordWizardSet
	// GetToolbar returns a toolbar populated with the actions linked to this model
	GetToolbar() webtypes.Toolbar
	// GroupBy returns a new RecordSet grouped with the given GROUP BY expressions.
	GroupBy(exprs ...models.FieldName) UserChangePasswordWizardSet
	// Intersect returns a new RecordCollection with only the records that are both
	// in this RecordCollection and in the other RecordSet.
	Intersect(other UserChangePasswordWizardSet) UserChangePasswordWizardSet
	// Limit returns a new RecordSet with only the first 'limit' records.
	Limit(limit int) UserChangePasswordWizardSet
	// Load looks up cache for fields of the RecordCollection and
	// query database for missing values.
	// fields are the fields to retrieve in the expression format,
	// i.e. "User.Profile.Age" or "user_id.profile_id.age".
	// If no fields are given, all DB columns of the RecordCollection's
	// model are retrieved.
	Load(fields ...models.FieldName) UserChangePasswordWizardSet
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

	New(data UserChangePasswordWizardData) UserChangePasswordWizardSet
	// NormalizeM2MData converts the list of triplets received from the client into the final list of ids
	// to keep in the Many2Many relationship of this model through the given field.
	NormalizeM2MData(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Offset returns a new RecordSet with only the records starting at offset
	Offset(offset int) UserChangePasswordWizardSet
	// Onchange returns the values that must be modified according to each field's Onchange
	// method in the pseudo-record given as params.Values`,
	Onchange(params models.OnchangeParams) models.OnchangeResult
	// OrderBy returns a new RecordSet ordered by the given ORDER BY expressions.
	// Each expression contains a field name and optionally one of "asc" or "desc", such as:
	//
	// rs.OrderBy("Company", "Name desc")
	OrderBy(exprs ...string) UserChangePasswordWizardSet
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

	Search(condition q.UserChangePasswordWizardCondition) UserChangePasswordWizardSet
	// SearchAll returns a RecordSet with all items of the table, regardless of the
	// current RecordSet query. It is mainly meant to be used on an empty RecordSet.
	SearchAll() UserChangePasswordWizardSet

	SearchByName(name string, op operator.Operator, additionalCond q.UserChangePasswordWizardCondition, limit int) UserChangePasswordWizardSet
	// SearchCount fetch from the database the number of records that match the RecordSet conditions.
	SearchCount() int
	// SearchDomain execute a search on the given domain.
	SearchDomain(domain domains.Domain) CommonMixinSet
	// SearchRead retrieves database records according to the filters defined in params.
	SearchRead(params webtypes.SearchParams) []models.RecordData

	Sorted(less func(UserChangePasswordWizardSet, UserChangePasswordWizardSet) bool) UserChangePasswordWizardSet
	// SortedByField returns a new record set with the same records as rc but sorted by the given field.
	// If reverse is true, the sort is done in reversed order
	SortedByField(namer models.FieldName, reverse bool) UserChangePasswordWizardSet
	// SortedDefault returns a new record set with the same records as rc but sorted according
	// to the default order of this model
	SortedDefault() UserChangePasswordWizardSet
	// Subtract returns a RecordSet with the Records that are in this
	// RecordCollection but not in the given 'other' one.
	// The result is guaranteed to be a set of unique records.
	Subtract(other UserChangePasswordWizardSet) UserChangePasswordWizardSet
	// Sudo returns a new RecordSet with the given userID
	// or the superuser ID if not specified
	Sudo(userID ...int64) UserChangePasswordWizardSet
	// Union returns a new RecordSet that is the union of this RecordSet and the given
	// "other" RecordSet. The result is guaranteed to be a set of unique records.
	Union(other UserChangePasswordWizardSet) UserChangePasswordWizardSet
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
	WithContext(key string, value interface{}) UserChangePasswordWizardSet
	// WithEnv returns a copy of the current RecordSet with the given Environment.
	WithEnv(env models.Environment) UserChangePasswordWizardSet
	// WithNewContext returns a copy of the current RecordSet with its context
	// replaced by the given one.
	WithNewContext(context *types.Context) UserChangePasswordWizardSet

	Write(data UserChangePasswordWizardData) bool
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
	Super() UserChangePasswordWizardSet
	// ModelData returns a new UserChangePasswordWizardData object populated with the values
	// of the given FieldMap.
	ModelData(fMap models.FieldMap) UserChangePasswordWizardData
	// Records returns a slice with all the records of this RecordSet, as singleton RecordSets
	Records() []UserChangePasswordWizardSet
	// First returns the values of the first Record of the RecordSet as a pointer to a UserChangePasswordWizardData.
	//
	// If this RecordSet is empty, it returns an empty UserChangePasswordWizardData.
	First() UserChangePasswordWizardData
	// All returns the values of all Records of the RecordCollection as a slice of UserChangePasswordWizardData pointers.
	All() []UserChangePasswordWizardData
}

// UserChangePasswordWizardData is used to hold values of an UserChangePasswordWizard object instance
// when creating or updating a UserChangePasswordWizardSet.
type UserChangePasswordWizardData interface {
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
	// It returns the given UserChangePasswordWizardData so that calls can be chained
	Set(field models.FieldName, value interface{}) UserChangePasswordWizardData
	// Unset removes the value of the given field if it exists.
	//
	// It returns the given ModelData so that calls can be chained
	Unset(field models.FieldName) UserChangePasswordWizardData
	// Copy returns a copy of this UserChangePasswordWizardData
	Copy() UserChangePasswordWizardData
	// MergeWith updates this UserChangePasswordWizardData with the given other UserChangePasswordWizardData
	// If a field of the other UserChangePasswordWizardData already exists here, the value is overridden,
	// otherwise, the field is inserted.
	MergeWith(other UserChangePasswordWizardData)
	// Keys returns the UserChangePasswordWizardData keys as a slice of strings
	Keys() (res []string)
	// OrderedKeys returns the keys of this UserChangePasswordWizardData ordered.
	//
	// This has the convenient side effect of having shorter paths come before longer paths,
	// which is particularly useful when creating or updating related records.
	OrderedKeys() []string
	// FieldNames returns the UserChangePasswordWizardData keys as a slice of FieldNames.
	FieldNames() models.FieldNames
	// CreateDate returns the value of the CreateDate field.
	// If this CreateDate is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	CreateDate() dates.DateTime
	// HasCreateDate returns true if CreateDate is set in this UserChangePasswordWizardData
	HasCreateDate() bool
	// SetCreateDate sets the CreateDate field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetCreateDate(value dates.DateTime) UserChangePasswordWizardData
	// UnsetCreateDate removes the value of the CreateDate field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetCreateDate() UserChangePasswordWizardData

	// CreateUID returns the value of the CreateUID field.
	// If this CreateUID is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	CreateUID() int64
	// HasCreateUID returns true if CreateUID is set in this UserChangePasswordWizardData
	HasCreateUID() bool
	// SetCreateUID sets the CreateUID field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetCreateUID(value int64) UserChangePasswordWizardData
	// UnsetCreateUID removes the value of the CreateUID field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetCreateUID() UserChangePasswordWizardData

	// DisplayName returns the value of the DisplayName field.
	// If this DisplayName is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	DisplayName() string
	// HasDisplayName returns true if DisplayName is set in this UserChangePasswordWizardData
	HasDisplayName() bool
	// SetDisplayName sets the DisplayName field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetDisplayName(value string) UserChangePasswordWizardData
	// UnsetDisplayName removes the value of the DisplayName field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetDisplayName() UserChangePasswordWizardData

	// ID returns the value of the ID field.
	// If this ID is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	ID() int64
	// HasID returns true if ID is set in this UserChangePasswordWizardData
	HasID() bool
	// SetID sets the ID field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetID(value int64) UserChangePasswordWizardData
	// UnsetID removes the value of the ID field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetID() UserChangePasswordWizardData

	// LastUpdate returns the value of the LastUpdate field.
	// If this LastUpdate is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	LastUpdate() dates.DateTime
	// HasLastUpdate returns true if LastUpdate is set in this UserChangePasswordWizardData
	HasLastUpdate() bool
	// SetLastUpdate sets the LastUpdate field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetLastUpdate(value dates.DateTime) UserChangePasswordWizardData
	// UnsetLastUpdate removes the value of the LastUpdate field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetLastUpdate() UserChangePasswordWizardData

	// Users returns the value of the Users field.
	// If this Users is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	Users() UserChangePasswordWizardLineSet
	// HasUsers returns true if Users is set in this UserChangePasswordWizardData
	HasUsers() bool
	// SetUsers sets the Users field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetUsers(value UserChangePasswordWizardLineSet) UserChangePasswordWizardData
	// UnsetUsers removes the value of the Users field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetUsers() UserChangePasswordWizardData

	// CreateUsers stores the related UserChangePasswordWizardLineData to be used to create
	// a related record on the fly for Users.
	//
	// This method can be called multiple times to create multiple records
	CreateUsers(related UserChangePasswordWizardLineData) UserChangePasswordWizardData
	// WriteDate returns the value of the WriteDate field.
	// If this WriteDate is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	WriteDate() dates.DateTime
	// HasWriteDate returns true if WriteDate is set in this UserChangePasswordWizardData
	HasWriteDate() bool
	// SetWriteDate sets the WriteDate field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetWriteDate(value dates.DateTime) UserChangePasswordWizardData
	// UnsetWriteDate removes the value of the WriteDate field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetWriteDate() UserChangePasswordWizardData

	// WriteUID returns the value of the WriteUID field.
	// If this WriteUID is not set in this UserChangePasswordWizardData, then
	// the Go zero value for the type is returned.
	WriteUID() int64
	// HasWriteUID returns true if WriteUID is set in this UserChangePasswordWizardData
	HasWriteUID() bool
	// SetWriteUID sets the WriteUID field with the given value.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	SetWriteUID(value int64) UserChangePasswordWizardData
	// UnsetWriteUID removes the value of the WriteUID field if it exists.
	// It returns this UserChangePasswordWizardData so that calls can be chained.
	UnsetWriteUID() UserChangePasswordWizardData
}

// A UserChangePasswordWizardGroupAggregateRow holds a row of results of a query with a group by clause
type UserChangePasswordWizardGroupAggregateRow interface {
	// Values() returns the values of the actual query
	Values() UserChangePasswordWizardData
	// Count is the number of lines aggregated into this one
	Count() int
	// Condition can be used to query the aggregated rows separately if needed
	Condition() q.UserChangePasswordWizardCondition
}
