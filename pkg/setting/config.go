package setting

type ServerSettings struct {
	HttpPort string
}

type DatabaseSettings struct {
	DBType    string
	Username  string
	Password  string
	Host      string
	CharSet   string
	ParseTime bool
	DBName    string
}

type AppSettings struct {
	LogSavePath       string
	LogFileName       string
	LogFileExt        string
	MaxSize           int
	MaxAge            int
	MaxBackups        int
	Compress          bool
	EncoderType       string
	EncoderConfigType string
}

func (s *Setting) ReadConfig(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
