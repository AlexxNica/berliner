package berliner

// type context struct {
// 	stream stream
// 	farm renderFarm
// }

// func (c context) start() {
// 	make status
// 	return status
// 	go func() {
// 		postCh, messages := c.stream.start()
// 		posts := []Post
// 		for post := range postCh {
// 			posts = append(posts, post)
// 		}
// 		messages := c.farm.render(posts)
// 	}()
// }

// Add Sources, wraps source in stream, calls add child on toplevel stream, returns wrapped source
// Add Filter, calls add filter to stream