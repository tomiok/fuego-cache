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

//BulkAdd will get all the entries and return if the operation was successful or not and the number of errors.
func (c *cache) BulkAdd(entries []entry) BulkResponse {
	var res BulkResponse
	return res
}

//BulkDelete will delete all the keys in the cache and return if the response showing if any error occurred.
func (c *cache) BulkDelete(keys []interface{}) BulkResponse {
	var res BulkResponse
	return res
}

type BulkGetResponse struct {
	Value string `json:"value"`
	Err   bool   `json:"err"`
}

type BulkResponse struct {
	Success bool `json:"success"`
	Errs    int  `json:"errs"`
}
