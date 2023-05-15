package datastore

import "fmt"

func (d *Datastore) GetUsers() ([]string, error) {

	usersKey := fmt.Sprintf("%s:%s", d.cfg.ValueAsStr("REDIS_KEY_PREFIX"), UsersKey)

	storedUsers, err := d.client.SMembers(d.ctx, usersKey).Result()
	if err != nil {
		return storedUsers, err
	}

	return storedUsers, nil
}
