package replicator

type Op struct {
	Id        interface{}
	Operation string
	Namespace string
	Data      map[string]interface{}
	Timestamp bson.MongoTimestamp
}

func (this *Op) IsInsert() bool {
	return this.Operation == "i"
}

func (this *Op) IsUpdate() bool {
	return this.Operation == "u"
}

func (this *Op) IsDelete() bool {
	return this.Operation == "d"
}

func (this *Op) ParseNamespace() []string {
	return strings.SplitN(this.Namespace, ".", 2)
}

func (this *Op) GetDatabase() string {
	return this.ParseNamespace()[0]
}

func (this *Op) GetCollection() string {
	return this.ParseNamespace()[1]
}

func (this *Op) FetchData(session *mgo.Session) error {
	if this.IsDelete() {
		return nil
	}
	collection := session.DB(this.GetDatabase()).C(this.GetCollection())
	doc := make(map[string]interface{})
	err := collection.FindId(this.Id).One(doc)
	if err == nil {
		this.Data = doc
		return nil
	} else {
		return err
	}
}

func (this *Op) ParseLogEntry(entry OpLogEntry) {
	this.Operation = entry["op"].(string)
	this.Timestamp = entry["ts"].(bson.MongoTimestamp)
	// only parse inserts, deletes, and updates
	if this.IsInsert() || this.IsDelete() || this.IsUpdate() {
		var objectField OpLogEntry
		if this.IsUpdate() {
			objectField = entry["o2"].(OpLogEntry)
		} else {
			objectField = entry["o"].(OpLogEntry)
		}
		this.Id = objectField["_id"]
		this.Namespace = entry["ns"].(string)
	}
}
