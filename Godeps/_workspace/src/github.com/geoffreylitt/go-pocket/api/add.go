package api

// AddOption is the options for retrieve API.
type AddOption struct {
	Url					string				 `json:"url,omitempty"`
	Title				string				 `json:"title,omitempty"`
	Tags				string				 `json:"tags,omitempty"`
}

type addAPIOptionWithAuth struct {
	*AddOption
	authInfo
}

type AddResult struct {
	Status   int
	Complete int
}

// Retrieve returns the in Pocket
func (c *Client) Add(options *AddOption) (*AddResult, error) {
	data := addAPIOptionWithAuth{
		authInfo:       c.authInfo,
		AddOption: 			options,
	}

	res := &AddResult{}
	err := PostJSON("/v3/add", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
