package cache

//BulkGet will return all the keys and return the value if it is found, otherwise a fake response with an error.
func (c *cache) BulkGet(keys []interface{}) []BulkGetResponse {
	var res []BulkGetResponse
	for _, k := range keys {
		val, err := c.GetOne(k)
		var getResponse BulkGetResponse
		if err != nil {
			getResponse = BulkGetResponse{
				Value: responseNil,
				Err:   true,
			}
		} else {
			getResponse = BulkGetResponse{
				Value: val,
				Err:   false,
			}
		}
		res = append(res, getResponse)
	}

	return res
}

//BulkSet will get all the entries and return if the operation was successful or not and the number of errors.
func (c *cache) BulkSet(entries []entry) BulkResponse {
	var res BulkResponse
	count := 0
	for _, e := range entries {
		_, err := c.SetOne(e.key, e.object.value, int(e.object.ttl))
		if err != nil {
			count = 1
			break
		}
	}
	res.Err = count > 0
	return res
}

//BulkDelete will delete all the keys in the cache and return if the response showing if any error occurred.
func (c *cache) BulkDelete(keys []interface{}) BulkResponse {
	var res BulkResponse
	for _, key := range keys {
		c.DeleteOne(key)
	}
	return res
}

type BulkGetResponse struct {
	Value string `json:"value,omitempty"`
	Err   bool   `json:"err"`
}

type BulkResponse struct {
	Err     bool   `json:"err"`
	Message string `json:"message,omitempty"`
}
