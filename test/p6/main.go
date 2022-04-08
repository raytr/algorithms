package main

type Friend struct {
	Email   string
	Friends []*Friend
}

func (this *Friend) canBeFriend(f *Friend) bool {
	visited := make(map[string]bool)

	return dfs(this, f, visited)
}

func dfs(this, target *Friend, visited map[string]bool) bool {
	if this.Email == target.Email {
		return true
	}
	for _, curFriend := range this.Friends {
		if !visited[curFriend.Email] {
			visited[curFriend.Email] = true
			res := dfs(curFriend, target, visited)
			if res {
				return true
			}
			return false
		}
	}
	return false
}
