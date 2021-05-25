// 3つの変数a,b,cについて、a=b, b=c, c=aとなるような関数をポインタで実装しましょう。
package testcode

func Swap1(a, b, c *int) {
	*a, *b, *c = *b, *c, *a
}
