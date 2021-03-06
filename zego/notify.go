package zego

func (a Auth) SendEvent(event, msg string) (*Resource, error) {

	path := "apps/notify.json"
	param := `{"event": "` + event + `", "app_id": "0", "body":"` + msg + `"}`
	resource, err := api(a, "POST", path, param)
	if err != nil {
		return nil, err
	}

	return resource, nil

}
