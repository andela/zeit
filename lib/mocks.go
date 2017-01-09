package lib

func getMockConfig() *Config {
	return &Config{
		CurrentEntry: "KGDFDFD",
		Projects:     []KeyValue{KeyValue{Name: "Skilltree", ID: "HJ8E3ENM"}},
		Tags:         []KeyValue{KeyValue{Name: "Core", ID: "HUFD33JND"}},
		NewTags:      []KeyValue{},
		Entries:      []string{"6DFDFDFD"},
	}
}

func getMockUser() User {
	return User{
		Id:    "FGHJJSDS2444",
		Token: "DFDFGHJB77828SDS",
	}
}
