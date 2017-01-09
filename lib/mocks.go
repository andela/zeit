package lib

func getMockConfig() *Config {
    return &Config{
        CurrentEntry: "KGDFDFD",
        Projects:     []KeyValue{KeyValue{Name: "Skilltree", ID: "HJ8E3ENM"}, KeyValue{Name: "Zhisi", ID: "HDUGESW"}},
        Tags:         []KeyValue{},
        NewTags:      []KeyValue{},
        Entries:      []BasicEntry{},
    }
}

func getMockUser() User {
    return User{
        Id:    "FGHJJSDS2444",
        Token: "DFDFGHJB77828SDS",
    }
}
