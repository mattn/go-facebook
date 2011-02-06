package facebook

type Item struct {
	Id           string
	Name         string
	Category     string
	Created_Time string
}

type ItemCollection struct {
	Data []*Item
}

func (c *ItemCollection) Items() []*Item {
	if c.Data == nil {
		c.Data = make([]*Item, 0)
	}
	return c.Data
}

func (c *ItemCollection) Item(index int) *Item {
	if c.Data == nil {
		return nil
	}
	if len(c.Data) > index {
		return c.Data[index]
	}
	return nil
}
