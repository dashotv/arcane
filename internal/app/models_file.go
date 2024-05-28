package app

func (c *Connector) FileGet(id string) (*File, error) {
	m := &File{}
	err := c.File.Find(id, m)
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) FileList(page, limit int) ([]*File, error) {
	skip := (page - 1) * limit
	list, err := c.File.Query().Limit(limit).Skip(skip).Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}
