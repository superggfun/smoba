package doTask

func (m *Account) Subscribe() ([]string, error) {
	l := make([]string, 0, 10)
	u, err := m.GetNews()
	if err != nil {
		return nil, err
	}
	for key := range u {
		b, err := m.AddFriend(key)
		if err != nil {
			return nil, err
		} else if !b {
			m.DelFriend(key)
			m.AddFriend(key)
			l = append(l, key)
		} else {
			m.DelFriend(key)
			l = append(l, key)
		}
	}
	return l, nil
}
