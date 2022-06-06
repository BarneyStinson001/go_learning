package main

import "fmt"

func printSquare(n int,m int){
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func printHalfTriangle(n int){
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printTriangle(m int){
	fmt.Println("")
	for i:=0;i<m;i++{
		for j:=0;j<=2*m-2;j++{
			if j<=m-1+i && j>=m-1-i{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
func printFrameTriangle(m int) {
	fmt.Println("")

	for i:=0;i<m-1;i++{
		for j:=0;j<=2*m-2;j++{
			if j==m-1-i {
				fmt.Print("*")
			}else if j==m-1+i {
				fmt.Print("*")
			}else{
				fmt.Print(" ")

			}
	}
		fmt.Println("")
	}
	for i:=0;i<2*m-1;i++{
		fmt.Print("*")
	}
}

func printPlusSheet(n int){
	for i:=1;i<=n;i++{
		row_res:=""
		sep:=""
		for j:=1;j<=i;j++{
			fmt.Print(sep)
			fmt.Printf("%d * %d = %d",i,j,i*j)
			sep="\t"
		}
		fmt.Println(row_res)
	}
}


func main()  {
	printSquare(5,6)
	fmt.Println("========================")
	printHalfTriangle(5)
	fmt.Println("========================")
	printTriangle(5)
	fmt.Println("========================")
	printFrameTriangle(5)
	fmt.Println("========================")
	printPlusSheet(10)

}
