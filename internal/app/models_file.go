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

func (c *Connector) FileCreateOrUpdate(path string) (*File, error) {
	var f *File
	list, err := c.File.Query().Where("path", path).Desc("path").Limit(1).Run()
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		f = &File{Path: path}
	} else {
		f = list[0]
	}

	return f, nil
}
