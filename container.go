package container

//go:generate go run generate/main.go -package container -template arraylist -destination arraylist.gen.go -types NUMERIC,string,[]byte
//go:generate go run generate/main.go -package container -template unrolledlist -destination unrolledlist.gen.go -types NUMERIC,string,[]byte
