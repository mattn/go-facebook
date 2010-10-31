package facebook

type Friends []Object

func parseFriends(value []interface{}) (friends Friends) {
	friends = make(Friends, len(value))
	for i, v := range value {
		wp := v.(map[string]interface{})
		friends[i] = parseObject(wp)
	}
	return
}
