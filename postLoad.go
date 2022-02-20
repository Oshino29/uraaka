package main

func (ppp *Posts) Load(n int) {
    *ppp = append(*ppp, Post{Text: []byte("测试用post text"), Time: "2022-99-99"})
}