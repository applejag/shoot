package game

type Projectiles struct {
	items *Set2[Projectile]
}

func (ps *Projectiles) update() {
	for i, isSet := range ps.items.iter() {
		if !isSet {
			continue
		}
		keep := ps.items.ref(i).update()
		if !keep {
			ps.items.remove(i)
		}
	}
}

func (ps Projectiles) render() {
	for i, isSet := range ps.items.iter() {
		if !isSet {
			continue
		}
		ps.items.ref(i).render()
	}
}
