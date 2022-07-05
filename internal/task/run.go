package task

type TaskInfo struct {
	Url  string
	Uuid int64
	pre  *TaskInfo
	next *TaskInfo
}

func (s *Server) run() {

}

func (g *Server) Add(s *TaskInfo) {
	if g.task != nil {
		s.next = g.task.next
		g.task.pre = s
	}
	g.task = s
}

func (g *Server) Del(s *TaskInfo) {
	if s == nil {
		return
	}
	if s.pre == nil { //第一位
		g.task = s.next
		if s.next != nil { //有第二位
			s.next.pre = nil
		}
	} else if s.next == nil { //末尾
		s.pre.next = nil
	} else { //中间
		s.pre.next = s.next
		s.next.pre = s.pre
	}
}
