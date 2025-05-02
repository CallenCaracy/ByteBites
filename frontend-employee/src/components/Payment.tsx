import { useState } from "react";
import { ArrowLeft, CreditCard, Truck, Smartphone } from "lucide-react";

type PaymentMethod = "cod" | "credit" | "online" | null;
type OnlinePaymentType = "gcash" | "paymaya" | "paypal" | null;

const Payments = () => {
  const [selectedPayment, setSelectedPayment] = useState<PaymentMethod>(null);
  const [onlinePaymentType, setOnlinePaymentType] =
    useState<OnlinePaymentType>(null);

  // Don't render anything while checking auth

  const handleBack = () => {
    if (onlinePaymentType) {
      setOnlinePaymentType(null);
    } else if (selectedPayment) {
      setSelectedPayment(null);
    }
  };

  return (
    <div className="min-h-[400px] w-full max-w-[28rem] mx-auto p-8 bg-white rounded-2xl shadow-md flex flex-col items-center text-center transition-all duration-300">
      <div className="flex items-center justify-center w-full mb-8 relative text-center">
        {(selectedPayment || onlinePaymentType) && (
          <button
            onClick={handleBack}
            className="absolute left-0 p-2 rounded-full cursor-pointer transition-all duration-200 bg-transparent border-none hover:bg-gray-100"
          >
            <ArrowLeft className="w-5 h-5 text-black" />
          </button>
        )}
        <h2 className="text-2xl font-semibold text-gray-800 text-center w-full mx-auto">
          {!selectedPayment
            ? "Select Payment Method"
            : selectedPayment === "online" && !onlinePaymentType
            ? "Choose Online Payment"
            : "Payment Details"}
        </h2>
      </div>
      <div className="flex flex-col gap-4 w-full">
        {!selectedPayment && (
          <div className="flex flex-col gap-2 w-full p-2">
            <button
              onClick={() => setSelectedPayment("cod")}
              className="flex items-center p-5 border border-gray-200 rounded-2xl bg-[#fafafa] cursor-pointer transition-all duration-200 w-full mb-4 h-16 hover:bg-gray-50 hover:border-blue-600 hover:-translate-y-0.5 hover:shadow-sm"
            >
              <Truck className="w-6 h-6 mr-4 text-gray-800" />
              <span className="text-lg text-gray-800 font-medium">
                Cash on Delivery
              </span>
            </button>
            <button
              onClick={() => setSelectedPayment("credit")}
              className="flex items-center p-5 border border-gray-200 rounded-2xl bg-[#fafafa] cursor-pointer transition-all duration-200 w-full mb-4 h-16 hover:bg-gray-50 hover:border-blue-600 hover:-translate-y-0.5 hover:shadow-sm"
            >
              <CreditCard className="w-6 h-6 mr-4  text-gray-800" />
              <span className="text-lg  text-gray-800 font-medium">
                Credit Card
              </span>
            </button>
            <button
              onClick={() => setSelectedPayment("online")}
              className="flex items-center p-5 border border-gray-200 rounded-2xl bg-[#fafafa] cursor-pointer transition-all duration-200 w-full mb-4 h-16 hover:bg-gray-50 hover:border-blue-600 hover:-translate-y-0.5 hover:shadow-sm"
            >
              <Smartphone className="w-6 h-6 mr-4  text-gray-800" />
              <span className="text-lg  text-gray-800 font-medium">
                Online Payment
              </span>
            </button>
            <button
              onClick={() => window.history.back()}
              className="bg-transparent text-gray-800 p-4 border border-gray-200 rounded-2xl text-lg font-medium cursor-pointer transition-all duration-200 w-full mt-2 h-16 hover:bg-gray-50 hover:border-blue-600 hover:text-blue-600"
            >
              Back to Menu
            </button>
          </div>
        )}

        {selectedPayment === "online" && (
          <div className="w-[90%] mx-auto">
            {["gcash", "paymaya", "paypal"].map((method) => (
              <div key={method} className="flex flex-col w-full mb-2">
                <button
                  onClick={() =>
                    setOnlinePaymentType(method as OnlinePaymentType)
                  }
                  className={`flex items-center p-5 border rounded-2xl bg-white cursor-pointer transition-all duration-200 w-full mb-4 h-16 ${
                    onlinePaymentType === method
                      ? "border-blue-600 bg-gray-50"
                      : "border-gray-200"
                  }`}
                >
                  <span className="text-lg  text-gray-800 font-medium">
                    {method.charAt(0).toUpperCase() + method.slice(1)}
                  </span>
                </button>
                {onlinePaymentType === method && (
                  <form className="flex flex-col gap-4 p-4 mt-2 border border-gray-200 rounded-xl bg-white animate-slideDown">
                    <div className="flex flex-col gap-1.5 w-[90%] mx-auto">
                      <label className="text-sm font-medium text-gray-600 text-left -ml-2.5">
                        Phone Number
                      </label>
                      <input
                        type="tel"
                        placeholder="Enter your phone number"
                        className="p-3 border text-gray-800 border-gray-200 rounded-lg text-sm h-11 w-full transition-all duration-200 -ml-2.5"
                      />
                    </div>
                  </form>
                )}
              </div>
            ))}
          </div>
        )}

        {selectedPayment === "credit" && (
          <form className="flex flex-col gap-4 w-[90%] mx-auto p-5 border border-gray-200 rounded-xl bg-gray-50">
            <div className="flex flex-col gap-1.5 w-[90%] mx-auto">
              <label className="text-sm font-medium text-gray-600 text-left -ml-2.5">
                Card Number
              </label>
              <div className="w-[90%] mx-auto">
                <div className="relative flex items-center w-full">
                  <CreditCard className="absolute left-1 w-5 h-5 text-gray-400" />
                  <input
                    type="text"
                    placeholder="1234 5678 9012 3456"
                    className="pl-10 p-3 border text-gray-800 border-gray-200 rounded-lg text-sm h-11 w-full transition-all duration-200 -ml-2.5"
                    maxLength={19}
                  />
                </div>
              </div>
            </div>
            <div className="grid grid-cols-2 gap-4 w-[90%] mx-auto">
              <div className="flex flex-col gap-2">
                <label className="text-sm font-medium text-gray-600">
                  Expiry Date
                </label>
                <input
                  type="text"
                  placeholder="MM/YY"
                  className="p-3 border text-gray-800 border-gray-200 rounded-lg text-sm h-11 w-full"
                  maxLength={5}
                />
              </div>
              <div className="flex flex-col gap-2">
                <label className="text-sm font-medium text-gray-600">CVC</label>
                <input
                  type="text"
                  placeholder="123"
                  className="p-3 border text-gray-800 border-gray-200 rounded-lg text-sm h-11 w-full"
                  maxLength={3}
                />
              </div>
            </div>
          </form>
        )}

        {selectedPayment === "cod" && (
          <div className="text-center p-4 bg-green-50 border border-green-200 rounded-xl text-green-700">
            <p>You selected Cash on Delivery</p>
            <p className="text-sm text-green-700 mt-2">
              Payment will be collected upon delivery
            </p>
          </div>
        )}

        {selectedPayment && (
          <button className="bg-gray-950 text-white py-3.5 px-6 border-none rounded-xl text-base font-medium cursor-pointer transition-all duration-200 w-full mt-6 hover:bg-blue-700 hover:-translate-y-0.5 active:translate-y-0">
            Confirm Payment
          </button>
        )}
      </div>
    </div>
  );
};

export default Payments;
