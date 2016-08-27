# Simple Go & eddystone-url example

### Simple implementation of Eddystone-URL specification using Go ###

With Go installed, clone the project and run:

1. ```$ go get github.com/suapapa/go_eddystone```
2. ```$ go get github.com/currantlabs/gatt```
3. ```$ go run main.go```  

Depends of your Operating System you may need to execute the script with root previleges. So run: ```$ sudo go run main.go```

For Mac users, you will need to run ```$ GODEBUG=cgocheck=0 go run main.go``` due an issue at **PayPal/GATT** project [https://github.com/paypal/gatt/issues/76](https://github.com/paypal/gatt/issues/76).  

Done. You may be able to see your desired URL using Physical Web app on your Mobile Device.

It works in any computer with Bluetooth 4.0 hardware.