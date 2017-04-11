package rpcserver

type ObjectId struct {
	id int `json:"id"`
}

func buildOjectIds(ids []int) []ObjectId {
	res := make([]ObjectId, len(ids))
	for i, id := range ids {
		res[i] = ObjectId{id}
	}
	return res
}