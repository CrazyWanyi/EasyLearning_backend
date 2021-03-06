package model

import (
	"easy_learning/db"
	"errors"
	"github.com/globalsign/mgo/bson"
	"time"
)

// Class 班级模型
type Class struct {
	Id          bson.ObjectId   `bson:"_id"`
	TeacherId   bson.ObjectId   `bson:"teacherId"`
	Classname   string          `bson:"classname"`
	Description string          `bson:"description"`
	StudentList []bson.ObjectId `bson:"studentList"`
	ExamList    []bson.ObjectId `bson:"examList"`
	CreatedAt   time.Time       `bson:"createdAt"`
	UpdatedAt   time.Time       `bson:"updatedAt"`
}


// CreateUser 创建班级
func (class *Class) CreateClass() error {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("class")

	class.Id = bson.NewObjectId()
	class.CreatedAt = time.Now()
	class.UpdatedAt = class.CreatedAt

	class.StudentList = []bson.ObjectId{}
	class.ExamList = []bson.ObjectId{}

	return client.Insert(class)
}

// FindClassById 通过 Id 查找班级
func FindClassById(id string) (class Class, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("class")
	if !bson.IsObjectIdHex(id) {
		return Class{}, errors.New("error")
	}
	if err = client.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&class); err == nil {
		return class, nil
	} else {
		return Class{}, err
	}
}

// FindClassByClassname 通过 Classname 查找班级
func FindClassByClassname(classname string) (class Class, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("class")

	if err = client.Find(bson.M{"classname": classname}).One(&class); err == nil {
		return class, nil
	} else {
		return Class{}, err
	}
}

// InsertStudentList
func InsertStudentList(uid string, cid string) (err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("class")


	selector := bson.M{"_id": bson.ObjectIdHex(cid)}
	update := bson.M{"$addToSet": bson.M{"studentList": bson.ObjectIdHex(uid)}}
	err = client.Update(selector, update)
	return err
}