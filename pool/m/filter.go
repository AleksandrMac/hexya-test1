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

// FilterSet is an autogenerated type to handle Filter objects.
type FilterSet interface {
	models.RecordSet
	// FilterSetHexyaFunc is a dummy function to uniquely match interfaces.
	FilterSetHexyaFunc()
	// ForceLoad reloads the cache for the given fields and updates the ids of this FilterSet.
	//
	// If no fields are given, all DB columns of the Filter model are retrieved.
	//
	// It also returns this FilterSet.
	ForceLoad(fields ...models.FieldName) FilterSet
	// Action is a getter for the value of the "Action" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Action() int64
	// SetAction is a setter for the value of the "Action" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetAction panics if the RecordSet is empty.
	SetAction(value int64)
	// Active is a getter for the value of the "Active" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Active() bool
	// SetActive is a setter for the value of the "Active" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetActive panics if the RecordSet is empty.
	SetActive(value bool)
	// Context is a getter for the value of the "Context" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Context() string
	// SetContext is a setter for the value of the "Context" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetContext panics if the RecordSet is empty.
	SetContext(value string)
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
	// Domain is a getter for the value of the "Domain" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Domain() string
	// SetDomain is a setter for the value of the "Domain" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetDomain panics if the RecordSet is empty.
	SetDomain(value string)
	// HexyaExternalID is a getter for the value of the "HexyaExternalID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	HexyaExternalID() string
	// SetHexyaExternalID is a setter for the value of the "HexyaExternalID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetHexyaExternalID panics if the RecordSet is empty.
	SetHexyaExternalID(value string)
	// HexyaVersion is a getter for the value of the "HexyaVersion" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	HexyaVersion() int
	// SetHexyaVersion is a setter for the value of the "HexyaVersion" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetHexyaVersion panics if the RecordSet is empty.
	SetHexyaVersion(value int)
	// ID is a getter for the value of the "ID" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	ID() int64
	// SetID is a setter for the value of the "ID" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetID panics if the RecordSet is empty.
	SetID(value int64)
	// IsDefault is a getter for the value of the "IsDefault" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	IsDefault() bool
	// SetIsDefault is a setter for the value of the "IsDefault" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetIsDefault panics if the RecordSet is empty.
	SetIsDefault(value bool)
	// LastUpdate is a getter for the value of the "LastUpdate" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	LastUpdate() dates.DateTime
	// SetLastUpdate is a setter for the value of the "LastUpdate" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetLastUpdate panics if the RecordSet is empty.
	SetLastUpdate(value dates.DateTime)
	// Name is a getter for the value of the "Name" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Name() string
	// SetName is a setter for the value of the "Name" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetName panics if the RecordSet is empty.
	SetName(value string)
	// ResModel is a getter for the value of the "ResModel" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	ResModel() string
	// SetResModel is a setter for the value of the "ResModel" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetResModel panics if the RecordSet is empty.
	SetResModel(value string)
	// Sort is a getter for the value of the "Sort" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	Sort() string
	// SetSort is a setter for the value of the "Sort" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetSort panics if the RecordSet is empty.
	SetSort(value string)
	// User is a getter for the value of the "User" field of the first
	// record in this RecordSet. It returns the Go zero value if the RecordSet is empty.
	User() UserSet
	// SetUser is a setter for the value of the "User" field of this
	// RecordSet. All Records of this RecordSet will be updated. Each call to this
	// method makes an update query in the database.
	//
	// SetUser panics if the RecordSet is empty.
	SetUser(value UserSet)
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
	// ActionArchive sets Active=false on a recordset, by calling ToggleActive to take the
	// corresponding actions according to the model
	ActionArchive()
	// ActionUnarchive sets Active=true on a recordset, by calling ToggleActive to take the
	// corresponding actions according to the model
	ActionUnarchive()
	// AddDomainLimitOffset adds the given domain, limit, offset
	// and order to the current RecordSet query.
	AddDomainLimitOffset(domain domains.Domain, limit int, offset int, order string) FilterSet
	// AddModifiers adds the modifiers attribute nodes to given xml doc.
	AddModifiers(doc *etree.Document, fieldInfos map[string]*models.FieldInfo)
	// AddNameToRelations returns the given RecordData after getting the name of all 2one relation ids
	AddNamesToRelations(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData

	Aggregates(fieldNames ...models.FieldName) []FilterGroupAggregateRow
	// Browse returns a new RecordSet with only the records with the given ids.
	// Note that this function is just a shorcut for Search on a list of ids.
	Browse(ids []int64) FilterSet
	// BrowseOne returns a new RecordSet with only the record with the given id.
	// Note that this function is just a shorcut for Search on a given id.
	BrowseOne(id int64) FilterSet

	CartesianProduct(others ...FilterSet) []FilterSet
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
	// CheckGlobalDefault checks if there is a global default for the ResModel requested.
	//
	// If there is, and the default is different than the record being written
	// (-> we're not updating the current global default), raise an error
	// to avoid users unknowingly overwriting existing global defaults (they
	// have to explicitly remove the current default before setting a new one)
	//
	// This method should only be called if 'vals' is trying to set 'IsDefault'
	CheckGlobalDefault(values FilterData, matchingFilters []FilterData)
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

	Copy(overrides FilterData) FilterSet

	CopyData(overrides FilterData) FilterData

	Create(data FilterData) FilterSet
	// CreateOrReplace creates or updates the filter with the given parameters.
	// Filter is considered the same if it has the same name (case insensitive) and the same user (if it has one).
	CreateOrReplace(vals FilterData) FilterSet

	DefaultGet() FilterData
	// Enqueue queues the execution of the given method with the given arguments on this recordset.
	// description will be the name given to the job.
	Enqueue(description string, method models.Methoder, arguments ...interface{}) QueueJobSet
	// Equals returns true if this RecordSet is the same as other
	// i.e. they are of the same model and have the same ids
	Equals(other FilterSet) bool
	// ExecuteO2MActions executes the actions on one2many fields given by
	// the list of triplets received from the client
	ExecuteO2MActions(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Fetch query the database with the current filter and returns a RecordSet
	// with the queries ids.
	//
	// Fetch is lazy and only return ids. Use Load() instead if you want to fetch all fields.
	Fetch() FilterSet
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

	Filtered(test func(FilterSet) bool) FilterSet
	// FormatRelationFields returns the given data with all relation fields converted to int64 or []int64
	FormatRelationFields(data models.RecordData, fInfos map[string]*models.FieldInfo) models.RecordData
	// GetActionCondition returns a condition for matching filters that are visible in the
	// same context (menu/view) as the given action.
	GetActionCondition(action int64) q.FilterCondition
	// GetFilters returns the filters for the given model and actionID for the current user
	GetFilters(modelName string, actionID int64) []FilterData
	// GetFormviewAction returns an action to open the document.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.`,
	GetFormviewAction() *actions.Action
	// GetFormviewID returns an view id to open the document with.
	// This method is meant to be overridden in addons that want
	// to give specific view ids for example.
	GetFormviewId() string
	// GetRecord returns the Recordset with the given externalID. It panics if the externalID does not exist.
	GetRecord(externalID string) FilterSet
	// GetToolbar returns a toolbar populated with the actions linked to this model
	GetToolbar() webtypes.Toolbar
	// GroupBy returns a new RecordSet grouped with the given GROUP BY expressions.
	GroupBy(exprs ...models.FieldName) FilterSet
	// Intersect returns a new RecordCollection with only the records that are both
	// in this RecordCollection and in the other RecordSet.
	Intersect(other FilterSet) FilterSet
	// Limit returns a new RecordSet with only the first 'limit' records.
	Limit(limit int) FilterSet
	// Load looks up cache for fields of the RecordCollection and
	// query database for missing values.
	// fields are the fields to retrieve in the expression format,
	// i.e. "User.Profile.Age" or "user_id.profile_id.age".
	// If no fields are given, all DB columns of the RecordCollection's
	// model are retrieved.
	Load(fields ...models.FieldName) FilterSet
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

	New(data FilterData) FilterSet
	// NormalizeM2MData converts the list of triplets received from the client into the final list of ids
	// to keep in the Many2Many relationship of this model through the given field.
	NormalizeM2MData(fieldName models.FieldName, info *models.FieldInfo, value interface{}) interface{}
	// Offset returns a new RecordSet with only the records starting at offset
	Offset(offset int) FilterSet
	// Onchange returns the values that must be modified according to each field's Onchange
	// method in the pseudo-record given as params.Values`,
	Onchange(params models.OnchangeParams) models.OnchangeResult
	// OrderBy returns a new RecordSet ordered by the given ORDER BY expressions.
	// Each expression contains a field name and optionally one of "asc" or "desc", such as:
	//
	// rs.OrderBy("Company", "Name desc")
	OrderBy(exprs ...string) FilterSet
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

	Search(condition q.FilterCondition) FilterSet
	// SearchAll returns a RecordSet with all items of the table, regardless of the
	// current RecordSet query. It is mainly meant to be used on an empty RecordSet.
	SearchAll() FilterSet

	SearchByName(name string, op operator.Operator, additionalCond q.FilterCondition, limit int) FilterSet
	// SearchCount fetch from the database the number of records that match the RecordSet conditions.
	SearchCount() int
	// SearchDomain execute a search on the given domain.
	SearchDomain(domain domains.Domain) CommonMixinSet
	// SearchRead retrieves database records according to the filters defined in params.
	SearchRead(params webtypes.SearchParams) []models.RecordData

	Sorted(less func(FilterSet, FilterSet) bool) FilterSet
	// SortedByField returns a new record set with the same records as rc but sorted by the given field.
	// If reverse is true, the sort is done in reversed order
	SortedByField(namer models.FieldName, reverse bool) FilterSet
	// SortedDefault returns a new record set with the same records as rc but sorted according
	// to the default order of this model
	SortedDefault() FilterSet
	// Subtract returns a RecordSet with the Records that are in this
	// RecordCollection but not in the given 'other' one.
	// The result is guaranteed to be a set of unique records.
	Subtract(other FilterSet) FilterSet
	// Sudo returns a new RecordSet with the given userID
	// or the superuser ID if not specified
	Sudo(userID ...int64) FilterSet
	// ToggleActive toggles the Active field of this object if it exists.
	ToggleActive()
	// Union returns a new RecordSet that is the union of this RecordSet and the given
	// "other" RecordSet. The result is guaranteed to be a set of unique records.
	Union(other FilterSet) FilterSet
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
	WithContext(key string, value interface{}) FilterSet
	// WithEnv returns a copy of the current RecordSet with the given Environment.
	WithEnv(env models.Environment) FilterSet
	// WithNewContext returns a copy of the current RecordSet with its context
	// replaced by the given one.
	WithNewContext(context *types.Context) FilterSet

	Write(data FilterData) bool
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
	Super() FilterSet
	// ModelData returns a new FilterData object populated with the values
	// of the given FieldMap.
	ModelData(fMap models.FieldMap) FilterData
	// Records returns a slice with all the records of this RecordSet, as singleton RecordSets
	Records() []FilterSet
	// First returns the values of the first Record of the RecordSet as a pointer to a FilterData.
	//
	// If this RecordSet is empty, it returns an empty FilterData.
	First() FilterData
	// All returns the values of all Records of the RecordCollection as a slice of FilterData pointers.
	All() []FilterData
}

// FilterData is used to hold values of an Filter object instance
// when creating or updating a FilterSet.
type FilterData interface {
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
	// It returns the given FilterData so that calls can be chained
	Set(field models.FieldName, value interface{}) FilterData
	// Unset removes the value of the given field if it exists.
	//
	// It returns the given ModelData so that calls can be chained
	Unset(field models.FieldName) FilterData
	// Copy returns a copy of this FilterData
	Copy() FilterData
	// MergeWith updates this FilterData with the given other FilterData
	// If a field of the other FilterData already exists here, the value is overridden,
	// otherwise, the field is inserted.
	MergeWith(other FilterData)
	// Keys returns the FilterData keys as a slice of strings
	Keys() (res []string)
	// OrderedKeys returns the keys of this FilterData ordered.
	//
	// This has the convenient side effect of having shorter paths come before longer paths,
	// which is particularly useful when creating or updating related records.
	OrderedKeys() []string
	// FieldNames returns the FilterData keys as a slice of FieldNames.
	FieldNames() models.FieldNames
	// Action returns the value of the Action field.
	// If this Action is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Action() int64
	// HasAction returns true if Action is set in this FilterData
	HasAction() bool
	// SetAction sets the Action field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetAction(value int64) FilterData
	// UnsetAction removes the value of the Action field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetAction() FilterData

	// Active returns the value of the Active field.
	// If this Active is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Active() bool
	// HasActive returns true if Active is set in this FilterData
	HasActive() bool
	// SetActive sets the Active field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetActive(value bool) FilterData
	// UnsetActive removes the value of the Active field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetActive() FilterData

	// Context returns the value of the Context field.
	// If this Context is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Context() string
	// HasContext returns true if Context is set in this FilterData
	HasContext() bool
	// SetContext sets the Context field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetContext(value string) FilterData
	// UnsetContext removes the value of the Context field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetContext() FilterData

	// CreateDate returns the value of the CreateDate field.
	// If this CreateDate is not set in this FilterData, then
	// the Go zero value for the type is returned.
	CreateDate() dates.DateTime
	// HasCreateDate returns true if CreateDate is set in this FilterData
	HasCreateDate() bool
	// SetCreateDate sets the CreateDate field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetCreateDate(value dates.DateTime) FilterData
	// UnsetCreateDate removes the value of the CreateDate field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetCreateDate() FilterData

	// CreateUID returns the value of the CreateUID field.
	// If this CreateUID is not set in this FilterData, then
	// the Go zero value for the type is returned.
	CreateUID() int64
	// HasCreateUID returns true if CreateUID is set in this FilterData
	HasCreateUID() bool
	// SetCreateUID sets the CreateUID field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetCreateUID(value int64) FilterData
	// UnsetCreateUID removes the value of the CreateUID field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetCreateUID() FilterData

	// DisplayName returns the value of the DisplayName field.
	// If this DisplayName is not set in this FilterData, then
	// the Go zero value for the type is returned.
	DisplayName() string
	// HasDisplayName returns true if DisplayName is set in this FilterData
	HasDisplayName() bool
	// SetDisplayName sets the DisplayName field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetDisplayName(value string) FilterData
	// UnsetDisplayName removes the value of the DisplayName field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetDisplayName() FilterData

	// Domain returns the value of the Domain field.
	// If this Domain is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Domain() string
	// HasDomain returns true if Domain is set in this FilterData
	HasDomain() bool
	// SetDomain sets the Domain field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetDomain(value string) FilterData
	// UnsetDomain removes the value of the Domain field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetDomain() FilterData

	// HexyaExternalID returns the value of the HexyaExternalID field.
	// If this HexyaExternalID is not set in this FilterData, then
	// the Go zero value for the type is returned.
	HexyaExternalID() string
	// HasHexyaExternalID returns true if HexyaExternalID is set in this FilterData
	HasHexyaExternalID() bool
	// SetHexyaExternalID sets the HexyaExternalID field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetHexyaExternalID(value string) FilterData
	// UnsetHexyaExternalID removes the value of the HexyaExternalID field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetHexyaExternalID() FilterData

	// HexyaVersion returns the value of the HexyaVersion field.
	// If this HexyaVersion is not set in this FilterData, then
	// the Go zero value for the type is returned.
	HexyaVersion() int
	// HasHexyaVersion returns true if HexyaVersion is set in this FilterData
	HasHexyaVersion() bool
	// SetHexyaVersion sets the HexyaVersion field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetHexyaVersion(value int) FilterData
	// UnsetHexyaVersion removes the value of the HexyaVersion field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetHexyaVersion() FilterData

	// ID returns the value of the ID field.
	// If this ID is not set in this FilterData, then
	// the Go zero value for the type is returned.
	ID() int64
	// HasID returns true if ID is set in this FilterData
	HasID() bool
	// SetID sets the ID field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetID(value int64) FilterData
	// UnsetID removes the value of the ID field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetID() FilterData

	// IsDefault returns the value of the IsDefault field.
	// If this IsDefault is not set in this FilterData, then
	// the Go zero value for the type is returned.
	IsDefault() bool
	// HasIsDefault returns true if IsDefault is set in this FilterData
	HasIsDefault() bool
	// SetIsDefault sets the IsDefault field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetIsDefault(value bool) FilterData
	// UnsetIsDefault removes the value of the IsDefault field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetIsDefault() FilterData

	// LastUpdate returns the value of the LastUpdate field.
	// If this LastUpdate is not set in this FilterData, then
	// the Go zero value for the type is returned.
	LastUpdate() dates.DateTime
	// HasLastUpdate returns true if LastUpdate is set in this FilterData
	HasLastUpdate() bool
	// SetLastUpdate sets the LastUpdate field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetLastUpdate(value dates.DateTime) FilterData
	// UnsetLastUpdate removes the value of the LastUpdate field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetLastUpdate() FilterData

	// Name returns the value of the Name field.
	// If this Name is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Name() string
	// HasName returns true if Name is set in this FilterData
	HasName() bool
	// SetName sets the Name field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetName(value string) FilterData
	// UnsetName removes the value of the Name field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetName() FilterData

	// ResModel returns the value of the ResModel field.
	// If this ResModel is not set in this FilterData, then
	// the Go zero value for the type is returned.
	ResModel() string
	// HasResModel returns true if ResModel is set in this FilterData
	HasResModel() bool
	// SetResModel sets the ResModel field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetResModel(value string) FilterData
	// UnsetResModel removes the value of the ResModel field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetResModel() FilterData

	// Sort returns the value of the Sort field.
	// If this Sort is not set in this FilterData, then
	// the Go zero value for the type is returned.
	Sort() string
	// HasSort returns true if Sort is set in this FilterData
	HasSort() bool
	// SetSort sets the Sort field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetSort(value string) FilterData
	// UnsetSort removes the value of the Sort field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetSort() FilterData

	// User returns the value of the User field.
	// If this User is not set in this FilterData, then
	// the Go zero value for the type is returned.
	User() UserSet
	// HasUser returns true if User is set in this FilterData
	HasUser() bool
	// SetUser sets the User field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetUser(value UserSet) FilterData
	// UnsetUser removes the value of the User field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetUser() FilterData

	// CreateUser stores the related UserData to be used to create
	// a related record on the fly for User.
	//
	// This method can be called multiple times to create multiple records
	CreateUser(related UserData) FilterData
	// WriteDate returns the value of the WriteDate field.
	// If this WriteDate is not set in this FilterData, then
	// the Go zero value for the type is returned.
	WriteDate() dates.DateTime
	// HasWriteDate returns true if WriteDate is set in this FilterData
	HasWriteDate() bool
	// SetWriteDate sets the WriteDate field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetWriteDate(value dates.DateTime) FilterData
	// UnsetWriteDate removes the value of the WriteDate field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetWriteDate() FilterData

	// WriteUID returns the value of the WriteUID field.
	// If this WriteUID is not set in this FilterData, then
	// the Go zero value for the type is returned.
	WriteUID() int64
	// HasWriteUID returns true if WriteUID is set in this FilterData
	HasWriteUID() bool
	// SetWriteUID sets the WriteUID field with the given value.
	// It returns this FilterData so that calls can be chained.
	SetWriteUID(value int64) FilterData
	// UnsetWriteUID removes the value of the WriteUID field if it exists.
	// It returns this FilterData so that calls can be chained.
	UnsetWriteUID() FilterData
}

// A FilterGroupAggregateRow holds a row of results of a query with a group by clause
type FilterGroupAggregateRow interface {
	// Values() returns the values of the actual query
	Values() FilterData
	// Count is the number of lines aggregated into this one
	Count() int
	// Condition can be used to query the aggregated rows separately if needed
	Condition() q.FilterCondition
}
