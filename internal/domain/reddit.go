package domain

type Threads struct {
	Data struct {
		After    string `json:"after"`
		Children []struct {
			Data struct {
				CrosspostParentList []struct {
					Permalink string `json:"permalink"`
				} `json:"crosspost_parent_list"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type Thread struct {
	Data struct {
		Children []struct {
			Data struct {
				Title  string `json:"title,omitempty"`
				Body   string `json:"body,omitempty"`
				ID     string `json:"id"`
				Author string `json:"author"`
			}
		} `json:"children"`
	} `json:"data"`
}
