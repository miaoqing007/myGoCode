func CompareTwoStringArrayIsSame(a,b []string)bool{
if len(am) != len(bm){
  return false
}
   for _, v := range a {
		  if _, ok := am[v]; !ok {
		  	am[v] = 0
	  	} else {
		  	am[v]++
	  	}
  	}
	for _, v := range b {
		if _, ok := bm[v]; !ok {
			bm[v] = 0
		} else {
			bm[v]++
		}
	}
	for k, v := range am {
		if z, ok := bm[k]; ok && z == v {
			continue
		}
		fmt.Println("不相同")
		return false
	}
	fmt.Println("相同")
  return true
}
