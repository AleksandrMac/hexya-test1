// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/filter"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// FilterModel is a strongly typed model definition that is used
// to extend the Filter model or to get a FilterSet through
// its NewSet() function.
//
// To get the unique instance of this type, call Filter().
type FilterModel struct {
	*models.Model
}

// NewSet returns a new FilterSet instance in the given Environment
func (md FilterModel) NewSet(env models.Environment) m.FilterSet {
	return filter.FilterSet{
		RecordCollection: env.Pool("Filter"),
	}
}

// Create creates a new Filter record and returns the newly created
// FilterSet instance.
func (md FilterModel) Create(env models.Environment, data m.FilterData) m.FilterSet {
	return filter.FilterSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new FilterSet instance
// with the records found.
func (md FilterModel) Search(env models.Environment, cond q.FilterCondition) m.FilterSet {
	return filter.FilterSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md FilterModel) Browse(env models.Environment, ids []int64) m.FilterSet {
	return filter.FilterSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md FilterModel) BrowseOne(env models.Environment, id int64) m.FilterSet {
	return filter.FilterSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty FilterData instance.
//
// Optional field maps if given will be used to populate the data.
func (md FilterModel) NewData(fm ...models.FieldMap) m.FilterData {
	return &filter.FilterData{
		ModelData: models.NewModelData(Filter(), fm...),
	}
}

// Fields returns the Field Collection of the Filter Model
func (md FilterModel) Fields() filter.FieldsCollection {
	return filter.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the Filter Model
func (md FilterModel) Methods() filter.MethodsCollection {
	return filter.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md FilterModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = FilterModel{}

// Coalesce takes a list of FilterSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md FilterModel) Coalesce(lst ...m.FilterSet) m.FilterSet {
	var last m.FilterSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// Filter returns the unique instance of the FilterModel type
// which is used to extend the Filter model or to get a FilterSet through
// its NewSet() function.
func Filter() FilterModel {
	return FilterModel{
		Model: models.Registry.MustGet("Filter"),
	}
}
