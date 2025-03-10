package rand

import "fmt"

func GenerateUUID() string {
	segment1 := GenerateAlphaFromPool(8, ProvideUUIDPool())
	segment2 := GenerateAlphaFromPool(4, ProvideUUIDPool())
	segment3 := GenerateAlphaFromPool(4, ProvideUUIDPool())
	segment4 := GenerateAlphaFromPool(4, ProvideUUIDPool())
	segment5 := GenerateAlphaFromPool(12, ProvideUUIDPool())
	return fmt.Sprintf("%s-%s-%s-%s-%s", segment1, segment2, segment3, segment4, segment5)
}
