package main

import (
	"fmt"
	// "strconv"
	"os"
)



func users(id string) string{



type Person struct{
	name string
	surname string
	amount int
	passcode int 
}


id543:=Person{
	name:"john",
	surname:"bele",
	amount:540,
	passcode:439890,
}

id293:=Person{
	name:"alin",
	surname:"salm",
	amount:1045,
	passcode:145237,
}

id764:=Person{
	name:"robert",
	surname:"barc",
	amount:10456,
	passcode:969696,
}

id154:=Person{
	name:"raal",
	surname:"manch",
	amount:459,
	passcode:958421,
}

id203:=Person{
	name:"janve",
	surname:"milan",
	amount:4545,
	passcode:784589,
}

id199:=Person{
	name:"saud",
	surname:"ektas",
	amount:45669,
	passcode:363671,
}

var result string=fmt.Sprint("id",id)

//id543
if result=="id543"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)


if passcodeuser==id543.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id543.name)
	start1:
	fmt.Print("1.Widthraw,2.Deposit,check balance")
	fmt.Scan(&account)
	if account==1{
		withdraw1:
		fmt.Printf("balance is :%d , what amount do you want to withdraw\n",id543.amount)
	var widhdrawcommand int 
		fmt.Scan(&widhdrawcommand)
		if widhdrawcommand>=id543.amount{
			fmt.Println("full amount cannot be withdrawn")
			goto withdraw1
		} 
		id543.amount=id543.amount-widhdrawcommand
		fmt.Printf("Current balance is %d\n ",id543.amount)

		
		goto start1 
		
	}
	if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id543.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id543.amount=id543.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id543.amount)
	goto start1
	}
	 if account==3{
		fmt.Println(id543.amount)
		goto start1
	}
	
}else{

	
	fmt.Println("passcode wrong")
		os.Exit(0)
	}

	
}
//id293
if result=="id293"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)


if passcodeuser==id293.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id293.name)
	start2:
	fmt.Print("1.Widthraw,2.Deposit,check balance")
	
	fmt.Scan(&account)
	if account==1{
		withdraw2:
		fmt.Printf("balance is :%d, what amound do you want to widhdraw\n",id293.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	if widhdrawcommand>=id293.amount{
			fmt.Println("full amount cannot be withdrawn")
			goto withdraw2
		} 
	id293.amount=id293.amount-widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id293.amount)
	goto start2
	}
	if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id293.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id293.amount=id293.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id293.amount)
	goto start2
	}
	if account==3{
		fmt.Println(id293.amount)
		goto start2
	}
}else{
	fmt.Println("passcode wrong")
		os.Exit(0)
	}

	
}

//id764
if result=="id764"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)


if passcodeuser==id764.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id764.name)
	start3:
	fmt.Print("1.Widthraw,2.Deposit,check balance")
	fmt.Scan(&account)
	if account==1{
		 withdraw3:
		fmt.Printf("balance is :%d , what amount do you want to withdraw\n",id764.amount)
	var widhdrawcommand int 
		fmt.Scan(&widhdrawcommand)
		if widhdrawcommand>=id764.amount{
			fmt.Println("full amount cannot be withdrawn")
			goto withdraw3
		} 
		id764.amount=id764.amount-widhdrawcommand
		fmt.Printf("Current balance is %d\n ",id764.amount)
		goto start3
	}
	if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id764.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id764.amount=id764.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\f ",id764.amount)
	goto start3
	}
	if account==3{
		fmt.Println(id764.amount)
		goto start3
	}
	



}else{
	fmt.Println("passcode wrong")
		os.Exit(0)
	}

	
}
//id154
if result=="id154"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)


if passcodeuser==id154.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id154.name)
	start4:
	fmt.Scan(&account)
		if account==1{
		fmt.Printf("balance is :%d , what amount do you want to withdraw\n",id154.amount)
	var widhdrawcommand int 
		fmt.Scan(&widhdrawcommand)
		
		id154.amount=id154.amount-widhdrawcommand
		fmt.Printf("Current balance is %d\n ",id154.amount)
		goto start4
	}
	if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id154.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id154.amount=id154.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id154.amount)
	goto start4
	}
	if account==3{
		fmt.Println(id764.amount)
		goto start4
	}
	
}else{
	fmt.Println("passcode wrong")
		os.Exit(0)
	}

	
}
if result=="id203"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)

//id203
if passcodeuser==id203.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id203.name)
	start5:
fmt.Print("1.Widthraw,2.Deposit,check balance")
	fmt.Scan(&account)
	
	if account==1{
		withdraw4:
		fmt.Printf("balance is :%d , what amount do you want to withdraw\n",id203.amount)
	var widhdrawcommand int 
		fmt.Scan(&widhdrawcommand)
		if widhdrawcommand>=id203.amount{
			fmt.Println("full amount cannot be withdrawn")
			goto withdraw4
		} 
		id203.amount=id203.amount-widhdrawcommand
		fmt.Printf("Current balance is %d\n ",id203.amount)
		goto start5

		
	}
	if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id203.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id203.amount=id203.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id203.amount)
	goto start5
	}
	if account==3{
		fmt.Println(id764.amount)
		goto start5
	}
}else{
	fmt.Println("passcode wrong")
		os.Exit(0)
	}


}
//id199
if result=="id199"{
	fmt.Println("enterpasscode")
	var passcodeuser int
	fmt.Scan(&passcodeuser)


if passcodeuser==id199.passcode{
	var account int 
	fmt.Printf("wellcome ,%s",id199.name)
start6:
fmt.Print("1.Widthraw,2.Deposit,check balance")
	fmt.Scan(&account)
	if account==1{
		withdraw5:
		fmt.Printf("balance is :%d , what amount do you want to withdraw\n",id199.amount)
	var widhdrawcommand int 
		fmt.Scan(&widhdrawcommand)
		if widhdrawcommand>=id199.amount{
			fmt.Println("full amount cannot be withdrawn")
			goto withdraw5
		} 
		id199.amount=id199.amount-widhdrawcommand
		fmt.Printf("Current balance is %d\n ",id199.amount)

	goto start6
}
if account==2{
		fmt.Printf("balance is :%d, what amound do you want to deposit\n",id199.amount)
	var widhdrawcommand int 
	fmt.Scan(&widhdrawcommand)
	id199.amount=id199.amount+widhdrawcommand
	fmt.Printf("Current balance is %d\n ",id199.amount)
	goto start6
	}
	if account==3{
		fmt.Println(id199.amount)
		goto start6
	}	else{
	fmt.Println("passcode wrong")
		os.Exit(0)
	}

	
}



}

return os.DevNull
}
func main(){
	fmt.Println("enter userid ")
	var enterdata string
	fmt.Scan(&enterdata)
	users(enterdata)

	

}