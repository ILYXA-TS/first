package main

// func main() {
// 	// es := []int64{1, 2, 3, 4, 5}
// 	// es1 := []float64{1.2, 1.3, 2.1}
// 	//es2 := []string{"a", "b", "c"}

// 	// fmt.Println(contains(es, 4))
// 	// fmt.Println(contains(es1, 4.1))
// 	// fmt.Println(contains(es1, 1.2))
// 	// fmt.Println(contains(es2, "c"))

// 	// anyy(es)
// 	// anyy(es1)
// 	// anyy(es2)

// 	// custom generic
// 	// fmt.Println(union(es))
// 	// fmt.Println(union(es1))

// 	// uniontype()

// }

// func ss[V int64 | float64](man []V) V {

// 	var num V

// 	for _, q := range man {
// 		num += q
// 	}
// 	return num
// }

// func contains[T comparable](slice []T, num T) bool {
// 	for _, er := range slice {
// 		if num == er {
// 			return true
// 		}
// 	}
// 	return false
// }

// func anyy[T any](entity ...T) {
// 	fmt.Println(entity)

// }

// не должно быть методов для создания своих дженериков

// ~ знак приближения типов
// type Numb interface {
// 	~int64 | float64
// }

// func union[V Numb](numbers []V) V {
// 	var num V

// 	for _, q := range numbers {
// 		num += q
// 	}
// 	return num
// }

// // обобщеные типы

// type Numbers[T Numb] []T

// func uniontype() {

// 	var numint Numbers[int64]

// 	numint = append(numint, []int64{1, 2, 3, 4, 5, 6}...)

// 	numfloat := Numbers[float64]{1, 2, 3, 4, 5, 6, 7, 8.2}

// 	fmt.Println(numfloat,"\n",numint)
// }

// type Ggg int64

// func (s Ggg) inpositive() bool {
// 	return s > 0
// }

// func typeinpro() {
// 	es := []Ggg{1, 2, 3, 4, 5, 6, 7}

// 	union(es)
// }
