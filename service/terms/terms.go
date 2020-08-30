package terms

import (
	"simple-core/database"
	"simple-core/graph/model"
	"simple-core/service/errmsg"

	json "github.com/json-iterator/go"
)

func Get(id int64, termType int, cols ...string) (*model.Term, error) {
	t := &database.Terms{Id: id, Type: termType}
	has, err := database.Engine().Cols(cols...).Get(t)
	if err != nil {
		return nil, errmsg.InternalError
	}

	if !has {
		return nil, errmsg.TermNotFoundError
	}

	term := &model.Term{
		ID:    id,
		Name:  t.Name,
		Meta:  nil,
		Count: t.Count,
	}

	return term, nil
}

func GetList(termType, offset, row int, cols ...string) ([]*model.Term, error) {
	var termList []database.Terms
	err := database.Engine().Cols(cols...).Where("type=?", termType).Limit(row, offset).Find(&termList)
	if err != nil {
		return nil, errmsg.InternalError
	}

	terms := make([]*model.Term, len(termList))
	meta := &model.TermMeta{}

	for i, t := range termList {
		if t.Meta != "" {
			err = json.UnmarshalFromString(t.Meta, meta)
			if err != nil {
				return nil, errmsg.InternalError
			}
		}
		terms[i] = &model.Term{
			ID:    t.Id,
			Name:  t.Name,
			Meta:  meta,
			Count: t.Count,
		}
	}

	return terms, nil
}

func GetNonNullList(termType, offset, row int, cols ...string) ([]*model.Term, error) {
	var termList []database.Terms
	err := database.Engine().Cols(cols...).Where("type=?",
		termType).And("count!=?", 0).Limit(row, offset).Find(&termList)
	if err != nil {
		return nil, errmsg.InternalError
	}

	terms := make([]*model.Term, len(termList))
	meta := &model.TermMeta{}

	for i, t := range termList {
		if t.Meta != "" {
			err = json.UnmarshalFromString(t.Meta, meta)
			if err != nil {
				return nil, errmsg.InternalError
			}
		}
		terms[i] = &model.Term{
			ID:    t.Id,
			Name:  t.Name,
			Meta:  meta,
			Count: t.Count,
		}
	}

	return terms, nil
}

func Add(termType int, name string, termMeta *model.TermMeta) (bool, error) {
	meta := ""
	var err error
	if termMeta != nil {
		meta, err = json.MarshalToString(termMeta)
		if err != nil {
			return false, errmsg.InternalError
		}
	}

	t := &database.Terms{
		Name:  name,
		Count: 0,
		Meta:  meta,
		Type:  termType,
	}

	_, err = database.Engine().InsertOne(t)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

func Alter(id int64, name string, termMeta *model.TermMeta) (bool, error) {
	t := &database.Terms{Id: id}
	has, err := database.Engine().Cols("version").Get(t)
	if err != nil {
		return false, errmsg.InternalError
	}

	if !has {
		return false, errmsg.TermNotFoundError
	}

	if name != "" {
		t.Name = name
	}

	t = &database.Terms{Version: t.Version, Name: t.Name}

	if termMeta != nil {
		meta, err := json.MarshalToString(termMeta)
		if err != nil {
			return false, errmsg.InternalError
		}
		t.Meta = meta
	}

	_, err = database.Engine().ID(id).Update(t)
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}

func Delete(id int64) (bool, error) {
	_, err := database.Engine().ID(id).Delete(new(database.Terms))
	if err != nil {
		return false, errmsg.InternalError
	}

	return true, nil
}
