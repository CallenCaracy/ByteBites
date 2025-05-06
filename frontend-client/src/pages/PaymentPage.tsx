import type React from "react";

import { useState } from "react";
import {
  ArrowLeft,
  CreditCard,
  Truck,
  Wifi,
  Mail,
  Phone,
  Info,
  ShoppingBag,
  AlertCircle,
} from "lucide-react";
import { useMutation, useQuery } from "@apollo/client";
import { CREATE_TRANSACTION, CREATE_RECEIPT } from "../graphql/Paymentqueries";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";
import { v4 as uuidv4 } from "uuid";
import Navbar from "../components/NavBar";

export default function PaymentPage() {
  const [selectedPayment, setSelectedPayment] = useState<string | null>(null);
  const [onlineMethod, setOnlineMethod] = useState<string | null>(null);
  const [phoneNumber, setPhoneNumber] = useState("");
  const [email, setEmail] = useState("");
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);
  const [paymentInfo, setPaymentInfo] = useState({
    cardNumber: "",
    cardHolder: "",
    expiryDate: "",
    cvv: "",
  });

  // Get authenticated user
  const { data: userData } = useQuery(GET_AUTHENTICATED_USER);

  // GraphQL mutations
  const [createTransaction, { loading: transactionLoading }] =
    useMutation(CREATE_TRANSACTION);
  const [createReceipt, { loading: receiptLoading }] =
    useMutation(CREATE_RECEIPT);
  // const [updateTransactionStatus, { loading: updateStatusLoading }] = useMutation(UPDATE_TRANSACTION_STATUS)

  // You can set this to an empty array to test the empty cart functionality
  const items = [{ name: "Premium Steak", qty: 2, price: 12345678.99 }];

  // This will work even if items is undefined, null, or empty
  const isCartEmpty = !items || items.length === 0;

  const handlePhoneNumberChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value.replace(/\D/g, "");
    if (value.length <= 11) {
      setPhoneNumber(value);
    }
  };

  const handlePaymentInfoChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setPaymentInfo((prev) => ({ ...prev, [name]: value }));
  };

  const validateCreditCardInfo = () => {
    return (
      paymentInfo.cardNumber.trim() !== "" &&
      paymentInfo.cardHolder.trim() !== "" &&
      paymentInfo.expiryDate.trim() !== "" &&
      paymentInfo.cvv.trim() !== ""
    );
  };

  const validateOnlinePaymentInfo = () => {
    if (onlineMethod === "gcash") {
      return phoneNumber.length === 11;
    } else if (onlineMethod === "paymaya") {
      return email.trim() !== "" || phoneNumber.length === 11;
    } else if (onlineMethod === "paypal") {
      return email.trim() !== "";
    }
    return false;
  };

  const handlePayNowClick = async () => {
    // Clear any previous messages
    setErrorMessage(null);
    setSuccessMessage(null);

    // Check if cart is empty
    if (isCartEmpty) {
      setErrorMessage("There's no order selected. Order now!");
      setTimeout(() => setErrorMessage(null), 3000);
      return;
    }

    // Check if payment method is selected
    if (!selectedPayment) {
      setErrorMessage("Please choose payment option first.");
      setTimeout(() => setErrorMessage(null), 3000);
      return;
    }

    // Validate payment information based on selected method
    if (selectedPayment === "credit" && !validateCreditCardInfo()) {
      setErrorMessage("Please fill out the required information first.");
      setTimeout(() => setErrorMessage(null), 3000);
      return;
    }

    if (selectedPayment === "online") {
      if (!onlineMethod) {
        setErrorMessage("Please select an online payment method.");
        setTimeout(() => setErrorMessage(null), 3000);
        return;
      }

      if (!validateOnlinePaymentInfo()) {
        setErrorMessage("Please fill out the required information first.");
        setTimeout(() => setErrorMessage(null), 3000);
        return;
      }
    }

    // If all validations pass, proceed with payment
    try {
      // Instead of creating a random UUID, we should use an actual order ID
      // If you don't have an actual order system yet, we can add more error handling
      // to make the API call more robust
      const mockOrderId = uuidv4();

      // Get user ID from authenticated user
      const userId = userData?.getAuthenticatedUser?.id;

      if (!userId) {
        setErrorMessage("User authentication required");
        return;
      }

      // Map payment method to GraphQL enum
      let paymentMethodEnum;
      if (selectedPayment === "credit") {
        paymentMethodEnum = "credit_card";
      } else if (selectedPayment === "online") {
        paymentMethodEnum = onlineMethod?.toLowerCase() || "other";
      } else {
        paymentMethodEnum = "cash";
      }

      // Add error handling and logging for better debugging
      console.log("Creating transaction with:", {
        orderId: mockOrderId,
        userId: userId,
        amountPaid: total,
        paymentMethod: paymentMethodEnum,
      });

      // Create transaction - remove the .replace(/"/g, '') as it's not needed
      // and might be causing issues with the enum parsing
      const transactionResult = await createTransaction({
        variables: {
          orderId: mockOrderId,
          userId: userId,
          amountPaid: total,
          paymentMethod: paymentMethodEnum,
          status: "completed",
        },
      });

      if (transactionResult.data?.createTransaction) {
        // Create receipt
        const receiptResult = await createReceipt({
          variables: {
            transactionId: transactionResult.data.createTransaction.id,
            userId: userId,
            amount: total,
            paymentMethod: paymentMethodEnum,
          },
        });

        if (receiptResult.data?.createReceipt) {
          setSuccessMessage("Payment processed successfully!");
          // Reset form or redirect as needed
        }
      }
    } catch (error) {
      console.error("Payment error:", error);
      setErrorMessage("Payment processing failed. Please try again.");
    }
  };

  // Safely calculate totals even if items is not defined
  const subtotal = items?.reduce((sum, item) => sum + item.price, 0) || 0;
  const taxRate = 0.1;
  const tax = subtotal * taxRate;
  const total = subtotal + tax;

  const renderPaymentDetails = () => {
    if (selectedPayment === "credit") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <h3 className="font-medium mb-4">Enter Credit Card Details</h3>
          <div className="space-y-4">
            <div>
              <label
                htmlFor="cardNumber"
                className="block text-sm text-gray-600 mb-1">
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
              <label
                htmlFor="cardHolder"
                className="block text-sm text-gray-600 mb-1">
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
                <label
                  htmlFor="expiryDate"
                  className="block text-sm text-gray-600 mb-1">
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
                <label
                  htmlFor="cvv"
                  className="block text-sm text-gray-600 mb-1">
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
      );
    } else if (selectedPayment === "cash") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <div className="bg-blue-50 p-4 rounded-md border border-blue-200">
            <p className="text-blue-800">
              Cash on Delivery selected. Pay when your order arrives.
            </p>
          </div>
        </div>
      );
    } else if (selectedPayment === "online") {
      return (
        <div className="mt-6 border-t border-gray-300 pt-6">
          <h3 className="font-medium mb-4">Select Online Payment Method</h3>
          <div className="grid grid-cols-3 gap-3">
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "gcash"
                  ? "border-green-500 bg-green-50"
                  : "hover:bg-gray-50"
              }`}
              onClick={() =>
                setOnlineMethod(onlineMethod === "gcash" ? null : "gcash")
              }>
              <span className="font-medium">GCash</span>
            </button>
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "paymaya"
                  ? "border-green-500 bg-green-50"
                  : "hover:bg-gray-50"
              }`}
              onClick={() =>
                setOnlineMethod(onlineMethod === "paymaya" ? null : "paymaya")
              }>
              <span className="font-medium">PayMaya</span>
            </button>
            <button
              className={`p-3 border border-gray-300 rounded-md flex flex-col items-center justify-center transition-colors ${
                onlineMethod === "paypal"
                  ? "border-green-500 bg-green-50"
                  : "hover:bg-gray-50"
              }`}
              onClick={() =>
                setOnlineMethod(onlineMethod === "paypal" ? null : "paypal")
              }>
              <span className="font-medium">PayPal</span>
            </button>
          </div>

          {onlineMethod === "gcash" && (
            <div className="mt-4">
              <label
                htmlFor="gcashNumber"
                className="block text-sm text-gray-600 mb-1">
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
              <p className="text-xs text-gray-500 mt-1">
                Enter your 11-digit GCash number
              </p>
            </div>
          )}

          {onlineMethod === "paymaya" && (
            <div className="mt-4 space-y-4">
              <div>
                <label
                  htmlFor="paymayaEmail"
                  className="block text-sm text-gray-600 mb-1">
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
                <p className="text-xs text-gray-500 mt-1">
                  Enter either email or phone number
                </p>
              </div>
              <div>
                <label
                  htmlFor="paymayaPhone"
                  className="block text-sm text-gray-600 mb-1">
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
              <label
                htmlFor="paypalEmail"
                className="block text-sm text-gray-600 mb-1">
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
      );
    }
    return null;
  };

  return (
    <>
      <Navbar />
      <div className="max-w-6xl mx-auto p-6">
        <div className="flex items-center mb-6">
          <button
            onClick={() => window.history.back()}
            className="flex items-center text-gray-600 hover:text-gray-900">
            <ArrowLeft className="mr-2 h-5 w-5" />
            <span>Back</span>
          </button>
          <h1 className="text-2xl font-bold text-center flex-grow">Payment</h1>
        </div>

        {/* Main content */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Payment methods */}
          <div className="lg:col-span-2 space-y-6">
            <div className="bg-white p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-semibold mb-4">Payment Method</h2>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <button
                  className={`p-4 border rounded-md flex flex-col items-center justify-center transition-colors ${
                    selectedPayment === "credit"
                      ? "border-green-500 bg-green-50"
                      : "hover:bg-gray-50"
                  }`}
                  onClick={() =>
                    setSelectedPayment(
                      selectedPayment === "credit" ? null : "credit"
                    )
                  }>
                  <CreditCard className="h-6 w-6 mb-2" />
                  <span>Credit Card</span>
                </button>
                <button
                  className={`p-4 border rounded-md flex flex-col items-center justify-center transition-colors ${
                    selectedPayment === "cash"
                      ? "border-green-500 bg-green-50"
                      : "hover:bg-gray-50"
                  }`}
                  onClick={() =>
                    setSelectedPayment(
                      selectedPayment === "cash" ? null : "cash"
                    )
                  }>
                  <Truck className="h-6 w-6 mb-2" />
                  <span>Cash on Delivery</span>
                </button>
                <button
                  className={`p-4 border rounded-md flex flex-col items-center justify-center transition-colors ${
                    selectedPayment === "online"
                      ? "border-green-500 bg-green-50"
                      : "hover:bg-gray-50"
                  }`}
                  onClick={() =>
                    setSelectedPayment(
                      selectedPayment === "online" ? null : "online"
                    )
                  }>
                  <Wifi className="h-6 w-6 mb-2" />
                  <span>Online Payment</span>
                </button>
              </div>

              {renderPaymentDetails()}
            </div>
          </div>

          {/* Order summary */}
          <div className="space-y-6">
            <div className="bg-white p-6 rounded-lg shadow-md">
              <h2 className="text-xl font-semibold mb-4">Order Summary</h2>
              {isCartEmpty ? (
                <div className="flex flex-col items-center justify-center py-6 text-gray-500">
                  <ShoppingBag className="h-12 w-12 mb-2" />
                  <p>Your cart is empty</p>
                </div>
              ) : (
                <>
                  <div className="space-y-4 mb-6">
                    {items.map((item, index) => (
                      <div key={index} className="flex justify-between">
                        <div>
                          <span className="font-medium">{item.name}</span>
                          <span className="text-gray-500 text-sm ml-2">
                            x{item.qty}
                          </span>
                        </div>
                        <span>${item.price.toFixed(2)}</span>
                      </div>
                    ))}
                  </div>
                  <div className="border-t border-gray-200 pt-4 space-y-2">
                    <div className="flex justify-between">
                      <span>Subtotal</span>
                      <span>${subtotal.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between">
                      <span>Tax (10%)</span>
                      <span>${tax.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between font-bold text-lg">
                      <span>Total</span>
                      <span>${total.toFixed(2)}</span>
                    </div>
                  </div>
                </>
              )}

              {errorMessage && (
                <div className="mt-4 p-3 bg-red-50 border border-red-200 rounded-md flex items-start">
                  <AlertCircle className="h-5 w-5 text-red-500 mr-2 flex-shrink-0 mt-0.5" />
                  <p className="text-red-700 text-sm">{errorMessage}</p>
                </div>
              )}

              {successMessage && (
                <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-md flex items-start">
                  <Info className="h-5 w-5 text-green-500 mr-2 flex-shrink-0 mt-0.5" />
                  <p className="text-green-700 text-sm">{successMessage}</p>
                </div>
              )}

              <button
                onClick={handlePayNowClick}
                disabled={transactionLoading || receiptLoading}
                className="w-full mt-6 bg-green-600 text-white py-3 rounded-md font-medium hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 disabled:opacity-50">
                {transactionLoading || receiptLoading
                  ? "Processing..."
                  : "Pay Now"}
              </button>
            </div>

            <div className="bg-blue-50 p-4 rounded-lg border border-blue-200">
              <div className="flex items-start">
                <Info className="h-5 w-5 text-blue-500 mr-2 flex-shrink-0 mt-0.5" />
                <div className="text-sm text-blue-700">
                  <p className="font-medium mb-1">Secure Payment</p>
                  <p>All transactions are secure and encrypted.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
