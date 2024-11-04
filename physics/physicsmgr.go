package physics

type PhysicsManager struct {
	objects []PhysicsObject
}

func NewManager() *PhysicsManager {
	return &PhysicsManager{
		objects: make([]PhysicsObject, 0, 50),
	}
}

func (p *PhysicsManager) RegisterObject(object PhysicsObject) {
	p.objects = append(p.objects, object)
}

func (p *PhysicsManager) CheckCollisions() {
	//ingenuous O(n^2) solution
	for i, e1 := range p.objects {
		for j, e2 := range p.objects {
			if i != j && CheckCollision(e1.GetCollider(), e2.GetCollider()) {
				e1.OnCollision(e2)
				e2.OnCollision(e1)
			}
		}
	}
}
