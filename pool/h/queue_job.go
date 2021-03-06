// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package h

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h/queue_job"
	"github.com/hexya-erp/pool/m"
	"github.com/hexya-erp/pool/q"
)

// ------- MODEL ---------

// QueueJobModel is a strongly typed model definition that is used
// to extend the QueueJob model or to get a QueueJobSet through
// its NewSet() function.
//
// To get the unique instance of this type, call QueueJob().
type QueueJobModel struct {
	*models.Model
}

// NewSet returns a new QueueJobSet instance in the given Environment
func (md QueueJobModel) NewSet(env models.Environment) m.QueueJobSet {
	return queue_job.QueueJobSet{
		RecordCollection: env.Pool("QueueJob"),
	}
}

// Create creates a new QueueJob record and returns the newly created
// QueueJobSet instance.
func (md QueueJobModel) Create(env models.Environment, data m.QueueJobData) m.QueueJobSet {
	return queue_job.QueueJobSet{
		RecordCollection: md.Model.Create(env, data),
	}
}

// Search searches the database and returns a new QueueJobSet instance
// with the records found.
func (md QueueJobModel) Search(env models.Environment, cond q.QueueJobCondition) m.QueueJobSet {
	return queue_job.QueueJobSet{
		RecordCollection: md.Model.Search(env, cond),
	}
}

// Browse returns a new RecordSet with the records with the given ids.
// Note that this function is just a shorcut for Search on a list of ids.
func (md QueueJobModel) Browse(env models.Environment, ids []int64) m.QueueJobSet {
	return queue_job.QueueJobSet{
		RecordCollection: md.Model.Browse(env, ids),
	}
}

// BrowseOne returns a new RecordSet with the record with the given id.
// Note that this function is just a shorcut for Search on the given id.
func (md QueueJobModel) BrowseOne(env models.Environment, id int64) m.QueueJobSet {
	return queue_job.QueueJobSet{
		RecordCollection: md.Model.BrowseOne(env, id),
	}
}

// NewData returns a pointer to a new empty QueueJobData instance.
//
// Optional field maps if given will be used to populate the data.
func (md QueueJobModel) NewData(fm ...models.FieldMap) m.QueueJobData {
	return &queue_job.QueueJobData{
		ModelData: models.NewModelData(QueueJob(), fm...),
	}
}

// Fields returns the Field Collection of the QueueJob Model
func (md QueueJobModel) Fields() queue_job.FieldsCollection {
	return queue_job.FieldsCollection{
		FieldsCollection: md.Model.Fields(),
	}
}

// Methods returns the Method Collection of the QueueJob Model
func (md QueueJobModel) Methods() queue_job.MethodsCollection {
	return queue_job.MethodsCollection{
		MethodsCollection: md.Model.Methods(),
	}
}

// Underlying returns the underlying models.Model instance
func (md QueueJobModel) Underlying() *models.Model {
	return md.Model
}

var _ models.Modeler = QueueJobModel{}

// Coalesce takes a list of QueueJobSet and return the first non-empty one
// if every record set is empty, it will return the last given
func (md QueueJobModel) Coalesce(lst ...m.QueueJobSet) m.QueueJobSet {
	var last m.QueueJobSet
	for _, elem := range lst {
		if elem.Collection().IsNotEmpty() {
			return elem
		}
		last = elem
	}
	return last
}

// QueueJob returns the unique instance of the QueueJobModel type
// which is used to extend the QueueJob model or to get a QueueJobSet through
// its NewSet() function.
func QueueJob() QueueJobModel {
	return QueueJobModel{
		Model: models.Registry.MustGet("QueueJob"),
	}
}
