package task

type TasksObj struct {
	Url  string
	Uuid int64
	Pre  *TasksObj
	Next *TasksObj
}

func (s *Server) run() {

}

//func (g *Server) Add(s *TasksObj) {
//	if g.task != nil {
//		s.Next = g.task.Next
//		g.task.Pre = s
//	}
//	g.task = s
//}

//func (g *Server) Del(s *TasksObj) {
//	if s == nil {
//		return
//	}
//	if s.Pre == nil { //第一位
//		g.task = s.Next
//		if s.Next != nil { //有第二位
//			s.Next.Pre = nil
//		}
//	} else if s.Next == nil { //末尾
//		s.Pre.Next = nil
//	} else { //中间
//		s.Pre.Next = s.Next
//		s.Next.Pre = s.Pre
//	}
//}
