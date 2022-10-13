package chdir

const (
	ERRORST = ": No such file or directory"
)

//return previous directory
func PrevDirctory(req []byte) string {
	var res []byte
	j := len(req) - 1

	//fetching the previous directory
	for j > 0 {
		if req[j] == '/' {
			break
		}
		j--
	}

	if j != 0 { //current path had more directory, got the previous one
		res = req[:j]
	} else { //current path is root, so return root
		res = append(res, '/')
	}
	return string(res)
}

//This is the main function which is called from main file. Takes two string arg, first is current path dir, second one is
//where to go, it return exact dirctory position
func OutPutPath(currPath, newPath string) string {
	var res string
	initLen, goPrev, orgiNewPath := len(newPath), false, newPath

	//looping the newPath req
	for len(newPath) > 0 {
		temp := []byte(newPath)

		//requested new path starting with root
		if hasPrefix(newPath, "/") {
			//if output path has / is suffix no need add it again
			if !hasSuffix(res, "/") {

				res = res + "/"

			}

			//removing from the requested new path
			newPath = string(temp[1:])
			goPrev = false
		} else if hasPrefix(newPath, "...") { // wrong path dirction is checking, returning error
			res = orgiNewPath + ERRORST
			break
		} else if checkError(newPath) { // wrong path dirction is checking, returning error, ex : ..klm
			res = orgiNewPath + ERRORST
			break
		} else if hasPrefix(newPath, "../") { // if present , current path should decreased by 1 directory
			tcurrPath := PrevDirctory([]byte(currPath))

			//setting output and current path as decreased current directory
			res, currPath = tcurrPath, tcurrPath

			//removing from the requested new path
			newPath = string(temp[3:])

			//to identify last loop aldreay move to previous directory
			goPrev = true
		} else if hasPrefix(newPath, "..") { // if present , current path should decreased by 1 directory
			tcurrPath := PrevDirctory([]byte(currPath))

			//setting output and current path as decreased current directory
			res, currPath = tcurrPath, tcurrPath

			//removing from the requested new path
			newPath = string(temp[2:])

			//to identify last loop aldreay move to previous directory
			goPrev = true
		} else if hasPrefix(newPath, ".") {
			res = currPath

			//removing from the requested new path
			newPath = string(temp[1:])
			goPrev = false
		} else {
			// requested new path does not start with root
			if initLen == len(newPath) {
				// currPath is not root then adding / at end of current path , so requested new path will be added after currPath
				if currPath[len(currPath)-1] != '/' {
					currPath = currPath + "/"
				}
				res = res + currPath
			} else if goPrev == true { //moved previous directory in last loop, so add / to add last to append new directory name
				res = res + "/"
				goPrev = false
			}
			res = res + string(temp[0])
			currPath = res

			//removing from the requested new path
			newPath = string(temp[1:])
		}

	}
	tempRes := []byte(res)
	if len(tempRes) > 1 && tempRes[len(tempRes)-1] == '/' {
		res = string(tempRes[:len(tempRes)-1])
	}
	return res

}

//checking path has error direction or not ex : ..klm
func checkError(req string) bool {
	res := false
	if len(req) > 2 && hasPrefix(req, "..") && req[2] != '/' {
		res = true
	}
	return res
}

//checking key is present in resquest at last or not ... mainly /, if there return true
func hasSuffix(req, key string) bool {

	res := false
	if len(req) > 0 && req[len(req)-1] == key[len(key)-1] {
		res = true
	}
	return res
}

//checking key is present in resquest at front or not , if there return true
func hasPrefix(req, key string) bool {

	res := true
	if len(req) >= len(key) {
		for i := 0; i < len(key); i++ {
			if req[i] != key[i] {
				res = false
				break
			}
		}
	} else {
		res = false
	}
	return res
}
