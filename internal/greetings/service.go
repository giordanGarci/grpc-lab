package greetings

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SayHello(name string) string {
	return "Hello, " + name + "!"
}
