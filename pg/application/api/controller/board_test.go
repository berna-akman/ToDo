package controller

//type BoardControllerTest struct {
//	suite.Suite
//	controller   BoardController
//	boardService *mocks.BoardService
//}
//
//func (c *BoardControllerTest) SetupTest() {
//	c.boardService = new(mocks.BoardService)
//	c.controller = &boardController{
//		s: c.boardService,
//	}
//}
//
//func Test_RunBoardControllerTestSuite(t *testing.T) {
//	suite.Run(t, new(BoardControllerTest))
//}
//
//func (c *BoardControllerTest) Test_BoardController_GetByID() {
//	t := c.T()
//
//	expected := &board.Board{}
//
//	c.boardService.On("GetByID").Return(expected, nil)
//
//	err := c.controller.GetByID()
//
//	assert.Nil(t, err)
//	assert.Equal(t, expected, err)
//}

//func TestNewBoardController(t *testing.T) {
//	type args struct {
//		s      presentation.BoardService
//		router *mux.Router
//	}
//	tests := []struct {
//		name string
//		args args
//		want BoardController
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewBoardController(tt.args.s, tt.args.router); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewBoardController() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_boardController_Create(t *testing.T) {
//	type fields struct {
//		s presentation.BoardService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &boardController{
//				s: tt.fields.s,
//			}
//			if err := c.Create(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_boardController_Delete(t *testing.T) {
//	type fields struct {
//		s presentation.BoardService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &boardController{
//				s: tt.fields.s,
//			}
//			if err := c.Delete(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_boardController_GetAll(t *testing.T) {
//	type fields struct {
//		s presentation.BoardService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &boardController{
//				s: tt.fields.s,
//			}
//			if err := c.GetAll(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_boardController_GetByID(t *testing.T) {
//	type fields struct {
//		s presentation.BoardService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &boardController{
//				s: tt.fields.s,
//			}
//			if err := c.GetByID(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_boardController_Update(t *testing.T) {
//	type fields struct {
//		s presentation.BoardService
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &boardController{
//				s: tt.fields.s,
//			}
//			if err := c.Update(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
//				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
