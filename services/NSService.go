package services

type NSService struct {
	NSMap *NSMap `inject:"-"`
}

func NewNSService() *NSService {
	return &NSService{}
}

func (this *NSService) ListAll() []string {
	ret := make([]string, 0)
	this.NSMap.data.Range(func(key, value interface{}) bool {
		ret = append(ret, key.(string))
		return true
	})
	return ret
}
