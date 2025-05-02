import type React from "react"

import { useState } from "react"
import { ArrowLeft, CreditCard, Truck, Wifi, Mail, Phone, Utensils, Info, ShoppingBag, AlertCircle } from "lucide-react"

export default function PaymentService() {
  const [selectedPayment, setSelectedPayment] = useState<string | null>(null)
  const [onlineMethod, setOnlineMethod] = useState<string | null>(null)
  const [phoneNumber, setPhoneNumber] = useState("")
  const [email, setEmail] = useState("")
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const [paymentInfo, setPaymentInfo] = useState({
    cardNumber: "",
    cardHolder: "",
    expiryDate: "",
    cvv: "",
  })

  // You can set this to an empty array to test the empty cart functionality
  const items = [
    { name: "Premium Steak", qty: 2, price: 59.98 },
    { name: "Premium Burger", qty: 1, price: 19.99 },
    { name: "Red Wine", qty: 1, price: 129.99 },
  ]

  // This will work even if items is undefined, null, or empty
  const isCartEmpty = !items || items.length === 0

  const handlePhoneNumberChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value.replace(/\D/g, "")
    if (value.length <= 11) {
      setPhoneNumber(value)
    }
  }

  const handlePaymentInfoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setPaymentInfo((prev) => ({ ...prev, [name]: value }))
  }

  const validateCreditCardInfo = () => {
    return (
      paymentInfo.cardNumber.trim() !== "" &&
      paymentInfo.cardHolder.trim() !== "" &&
      paymentInfo.expiryDate.trim() !== "" &&
      paymentInfo.cvv.trim() !== ""
    )
  }

  const validateOnlinePaymentInfo = () => {
    if (onlineMethod === "gcash") {
      return phoneNumber.length === 11
    } else if (onlineMethod === "paymaya") {
      return email.trim() !== "" || phoneNumber.length === 11
    } else if (onlineMethod === "paypal") {
      return email.trim() !== ""
    }
    return false
  }

  const handlePayNowClick = () => {
    // Clear any previous error messages
    setErrorMessage(null)

    // Check if cart is empty
    if (isCartEmpty) {
      setErrorMessage("There's no order selected. Order now!")
      setTimeout(() => setErrorMessage(null), 3000)
      return
    }

    // Check if payment method is selected
    if (!selectedPayment) {
      setErrorMessage("Please choose payment option first.")
      setTimeout(() => setErrorMessage(null), 3000)
      return
    }

    // Validate payment information based on selected method
    if (selectedPayment === "credit" && !validateCreditCardInfo()) {
      setErrorMessage("Please fill out the required information first.")
      setTimeout(() => setErrorMessage(null), 3000)
      return
    }

    if (selectedPayment === "online") {
      if (!onlineMethod) {
        setErrorMessage("Please select an online payment method.")
        setTimeout(() => setErrorMessage(null), 3000)
        return
      }

      if (!validateOnlinePaymentInfo()) {
        setErrorMessage("Please fill out the required information first.")
        setTimeout(() => setErrorMessage(null), 3000)
        return
      }
    }

    // If all validations pass, proceed with payment
    alert("Payment processing...")
  }

  // Safely calculate totals even if items is not defined
  const subtotal = items?.reduce((sum, item) => sum + item.price, 0) || 0
  const taxRate = 0.1
  const tax = subtotal * taxRate
  const total = subtotal + tax

  const renderPaymentDetails = () => {
    if (selectedPayment === "credit") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <h3 className="font-medium mb-4">Enter Credit Card Details</h3>
          <div className="space-y-4">
            <div>
              <label htmlFor="cardNumber" className="block text-sm text-gray-600 mb-1">
                Card Number <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                id="cardNumber"
                name="cardNumber"
                value={paymentInfo.cardNumber}
                onChange={handlePaymentInfoChange}
                placeholder="1234 5678 9012 3456"
                className="w-full p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                required
              />
            </div>
            <div>
              <label htmlFor="cardHolder" className="block text-sm text-gray-600 mb-1">
                Cardholder Name <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                id="cardHolder"
                name="cardHolder"
                value={paymentInfo.cardHolder}
                onChange={handlePaymentInfoChange}
                placeholder="John Doe"
                className="w-full p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                required
              />
            </div>
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label htmlFor="expiryDate" className="block text-sm text-gray-600 mb-1">
                  Expiry Date <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  id="expiryDate"
                  name="expiryDate"
                  value={paymentInfo.expiryDate}
                  onChange={handlePaymentInfoChange}
                  placeholder="MM/YY"
                  className="w-full p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                  required
                />
              </div>
              <div>
                <label htmlFor="cvv" className="block text-sm text-gray-600 mb-1">
                  CVV <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  id="cvv"
                  name="cvv"
                  value={paymentInfo.cvv}
                  onChange={handlePaymentInfoChange}
                  placeholder="123"
                  className="w-full p-2 border border-gray-300 border-gray-300rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                  required
                />
              </div>
            </div>
          </div>
        </div>
      )
    } else if (selectedPayment === "cash") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <div className="bg-blue-50 p-4 rounded-md border border-blue-200">
            <p className="text-blue-800">Cash on Delivery selected. Pay when your order arrives.</p>
          </div>
        </div>
      )
    } else if (selectedPayment === "online") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <h3 className="font-medium mb-4">Select Online Payment Method</h3>
          <div className="grid grid-cols-3 gap-3">
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "gcash" ? "border-green-500 bg-green-50" : "hover:bg-gray-50"
              }`}
              onClick={() => setOnlineMethod(onlineMethod === "gcash" ? null : "gcash")}
            >
              <span className="font-medium">GCash</span>
            </button>
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "paymaya" ? "border-green-500 bg-green-50" : "hover:bg-gray-50"
              }`}
              onClick={() => setOnlineMethod(onlineMethod === "paymaya" ? null : "paymaya")}
            >
              <span className="font-medium">PayMaya</span>
            </button>
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "paypal" ? "border-green-500 bg-green-50" : "hover:bg-gray-50"
              }`}
              onClick={() => setOnlineMethod(onlineMethod === "paypal" ? null : "paypal")}
            >
              <span className="font-medium">PayPal</span>
            </button>
          </div>

          {onlineMethod === "gcash" && (
            <div className="mt-4">
              <label htmlFor="gcashNumber" className="block text-sm text-gray-600 mb-1">
                GCash Phone Number <span className="text-red-500">*</span>
              </label>
              <div className="flex items-center">
                <Phone className="h-5 w-5 text-gray-400 mr-2" />
                <input
                  type="text"
                  id="gcashNumber"
                  value={phoneNumber}
                  onChange={handlePhoneNumberChange}
                  placeholder="09XXXXXXXXX"
                  className="flex-1 p-2 border rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                  maxLength={11}
                  required
                />
              </div>
              <p className="text-xs text-gray-500 mt-1">Enter your 11-digit GCash number</p>
            </div>
          )}

          {onlineMethod === "paymaya" && (
            <div className="mt-4 space-y-4">
              <div>
                <label htmlFor="paymayaEmail" className="block text-sm text-gray-600 mb-1">
                  PayMaya Email
                </label>
                <div className="flex items-center">
                  <Mail className="h-5 w-5 text-gray-400 mr-2" />
                  <input
                    type="email"
                    id="paymayaEmail"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="email@example.com"
                    className="flex-1 p-2 border rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                  />
                </div>
                <p className="text-xs text-gray-500 mt-1">Enter either email or phone number</p>
              </div>
              <div>
                <label htmlFor="paymayaPhone" className="block text-sm text-gray-600 mb-1">
                  Or PayMaya Phone Number
                </label>
                <div className="flex items-center">
                  <Phone className="h-5 w-5 text-gray-400 mr-2" />
                  <input
                    type="text"
                    id="paymayaPhone"
                    value={phoneNumber}
                    onChange={handlePhoneNumberChange}
                    placeholder="09XXXXXXXXX"
                    className="flex-1 p-2 border rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                    maxLength={11}
                  />
                </div>
              </div>
            </div>
          )}

          {onlineMethod === "paypal" && (
            <div className="mt-4">
              <label htmlFor="paypalEmail" className="block text-sm text-gray-600 mb-1">
                PayPal Email <span className="text-red-500">*</span>
              </label>
              <div className="flex items-center">
                <Mail className="h-5 w-5 text-gray-400 mr-2" />
                <input
                  type="email"
                  id="paypalEmail"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="email@example.com"
                  className="flex-1 p-2 border rounded-md focus:ring-2 focus:ring-green-500 focus:border-green-500"
                  required
                />
              </div>
            </div>
          )}
        </div>
      )
    }
    return (
      <div className="mt-6 pt-6 border-t border-gray-300">
        <div className="flex items-start p-4 bg-gray-50 rounded-md border border-gray-200">
          <Info className="h-5 w-5 text-blue-500 mt-0.5 mr-3 flex-shrink-0" />
          <p className="text-gray-600">Please choose a payment option to proceed with your order.</p>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-white p-4 md:p-8 max-w-6xl mx-auto relative">
      {/* Error Message */}
      {errorMessage && (
        <div className="fixed top-4 left-1/2 transform -translate-x-1/2 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md shadow-md flex items-center z-50">
          <AlertCircle className="h-5 w-5 mr-2" />
          <span>{errorMessage}</span>
        </div>
      )}

      {/* Header */}
      <header className="flex items-center justify-between mb-8">
        <button className="flex items-center rounded-full cursor-pointer transition-all duration-200 bg-transparent border-none hover:bg-gray-300">
          <ArrowLeft className="h-5 w-5" />
        </button>
        <div className="w-24"></div> {/* Spacer for alignment */}
      </header>

      <div className="flex flex-col md:flex-row gap-6">
        {/* Payment Method Section */}
        <div className="w-full md:w-3/5 border border-gray-300 rounded-lg p-6">
          <h2 className="text-xl font-bold mb-1">Payment Method</h2>
          <p className="text-gray-500 mb-6">Select your preferred payment method</p>

          <div className="space-y-4">
            <label
              className="flex items-center space-x-3 cursor-pointer p-2 rounded-md hover:bg-gray-50"
              onClick={() => {
                setSelectedPayment(selectedPayment === "credit" ? null : "credit")
                setOnlineMethod(null)
              }}
            >
              <div
                className={`h-5 w-5 rounded-full border flex items-center justify-center ${selectedPayment === "credit" ? "border-green-500" : "border-gray-300"}`}
              >
                {selectedPayment === "credit" && <div className="h-3 w-3 rounded-full bg-green-500"></div>}
              </div>
              <CreditCard className="h-5 w-5 text-gray-600" />
              <span className="text-gray-800">Credit Card</span>
            </label>

            <label
              className="flex items-center space-x-3 cursor-pointer p-2 rounded-md hover:bg-gray-50"
              onClick={() => {
                setSelectedPayment(selectedPayment === "cash" ? null : "cash")
                setOnlineMethod(null)
              }}
            >
              <div
                className={`h-5 w-5 rounded-full border flex items-center justify-center ${selectedPayment === "cash" ? "border-green-500" : "border-gray-300"}`}
              >
                {selectedPayment === "cash" && <div className="h-3 w-3 rounded-full bg-green-500"></div>}
              </div>
              <Truck className="h-5 w-5 text-gray-600" />
              <span className="text-gray-800">Cash on Delivery</span>
            </label>

            <label
              className="flex items-center space-x-3 cursor-pointer p-2 rounded-md hover:bg-gray-50"
              onClick={() => {
                setSelectedPayment(selectedPayment === "online" ? null : "online")
                setOnlineMethod(null)
              }}
            >
              <div
                className={`h-5 w-5 rounded-full border flex items-center justify-center ${selectedPayment === "online" ? "border-green-500" : "border-gray-300"}`}
              >
                {selectedPayment === "online" && <div className="h-3 w-3 rounded-full bg-green-500"></div>}
              </div>
              <Wifi className="h-5 w-5 text-gray-600" />
              <span className="text-gray-800">Online Payment</span>
            </label>
          </div>

          {renderPaymentDetails()}
        </div>

        {/* Order Summary Section */}
        <div className="w-full md:w-2/5 border border-gray-300 rounded-lg p-6">
          <h2 className="text-xl font-bold flex items-center mb-6">
            <Utensils className="h-6 w-6 mr-2" />
            Order Summary
          </h2>

          {isCartEmpty ? (
            <div className="text-center py-8 bg-amber-50 rounded-lg border border-amber-200">
              <ShoppingBag className="h-12 w-12 mx-auto text-amber-500 mb-2" />
              <p className="text-amber-800 font-medium text-lg">No food selected!</p>
              <p className="text-amber-700 mt-1">Order now!</p>
            </div>
          ) : (
            <>
              <div className="space-y-4 mb-6">
                {items.map((item, index) => (
                  <div key={index} className="flex justify-between">
                    <div>
                      <p className="font-medium">{item.name}</p>
                      <p className="text-gray-500 text-sm">Qty: {item.qty}</p>
                    </div>
                    <p className="font-medium">${item.price.toFixed(2)}</p>
                  </div>
                ))}
              </div>

              <div className="border-t pt-4 border-gray-300 space-y-2">
                <div className="flex justify-between">
                  <p className="text-gray-600">Subtotal</p>
                  <p className="font-medium">${subtotal.toFixed(2)}</p>
                </div>
                <div className="flex justify-between">
                  <p className="text-gray-600">Tax (10%)</p>
                  <p className="font-medium">${tax.toFixed(2)}</p>
                </div>
                <div className="flex justify-between font-bold text-lg mt-4">
                  <p>Total</p>
                  <p>${total.toFixed(2)}</p>
                </div>
              </div>
            </>
          )}

          <button
            className="w-full bg-green-500 hover:bg-green-600 text-white font-medium py-3 rounded-md mt-6 transition-colors"
            onClick={handlePayNowClick}
          >
            Pay Now
          </button>
        </div>
      </div>
    </div>
  )
}