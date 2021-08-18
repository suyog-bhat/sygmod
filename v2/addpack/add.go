package addpack

import "errors"

func Addnums(a int,b int) (int,error){
	if a<0{
		err:=errors.New("negative number")
		return 0,err
	}
	return a+b,nil

}
