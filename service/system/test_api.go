package system

type TestApi struct {
}

func (testApi *TestApi) SetTest(num int) (int, error) {
	//fmt.Println(num)
	return num, nil

}
