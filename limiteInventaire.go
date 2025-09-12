func (personnage *Character) IsInventoryFull() bool {
    count := 0
    for _, item := range personnage.Inventory {
        if item != "" && item != "..." {
            count++
        }
    }
    return count >= 10
}