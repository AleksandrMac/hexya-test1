// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/partner_title"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// PartnerTitleModel is a strongly typed model definition that is used
// to extend the PartnerTitle model or to get a PartnerTitleSet through
// its NewSet() function.
//
// To get the unique instance of this type, call PartnerTitle().
type PartnerTitleModel struct {
	*models.Model
}

// NewSet returns a new PartnerTitleSet instance in the given Environment
func (md PartnerTitleModel) NewSet(env models.Environment) m.PartnerTitleSet {
	return partner_title.PartnerTitleSet{
		RecordCollection: env.Pool("PartnerTitle"),
	}
}

// Create creates a new PartnerTitle record and returns the newly created
// PartnerTitleSet instance.
func (md PartnerTitleModel) Create(env models.Environment, data m.PartnerTitleData) m.PartnerTitleSet {
	return partner_title.PartnerTitleSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new PartnerTitleSet instance
// with the records found.
func (md PartnerTitleModel) Search(env models.Environment, cond q.PartnerTitleCondition) m.PartnerTitleSet {
	return partner_title.PartnerTitleSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md PartnerTitleModel) Browse(env models.Environment, ids []int64) m.PartnerTitleSet {
	return partner_title.PartnerTitleSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md PartnerTitleModel) BrowseOne(env models.Environment, id int64) m.PartnerTitleSet {
	return partner_title.PartnerTitleSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty PartnerTitleData instance.
//
// Optional field maps if given will be used to populate the data.
func (md PartnerTitleModel) NewData(fm ...models.FieldMap) m.PartnerTitleData {
	return &partner_title.PartnerTitleData{
		ModelData: models.NewModelData(PartnerTitle(), fm...),
	}
}

// Fields returns the Field Collection of the PartnerTitle Model
func (md PartnerTitleModel) Fields() partner_title.FieldsCollection {
	return partner_title.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the PartnerTitle Model
func (md PartnerTitleModel) Methods() partner_title.MethodsCollection {
	return partner_title.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md PartnerTitleModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = PartnerTitleModel{}

// Coalesce takes a list of PartnerTitleSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md PartnerTitleModel) Coalesce(lst ...m.PartnerTitleSet) m.PartnerTitleSet {
	var last m.PartnerTitleSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// PartnerTitle returns the unique instance of the PartnerTitleModel type
// which is used to extend the PartnerTitle model or to get a PartnerTitleSet through
// its NewSet() function.
func PartnerTitle() PartnerTitleModel {
	return PartnerTitleModel{
		Model: models.Registry.MustGet("PartnerTitle"),
	}
}