package institution

// Institution structure
type Institution struct {
    Id string
    Name string
    Rate float32
}

// Institution constructor
func New(id string, name string, rate float32) *Institution {
    return &Institution{Id:id, Name:name, Rate:rate}
}
