import { useState, useEffect } from "react";
import { ArrowLeft, CreditCard, Truck, Smartphone } from "lucide-react";
import { useNavigate } from "react-router-dom";
import styles from './Payment.module.css';

type PaymentMethod = "cod" | "credit" | "online" | null;
type OnlinePaymentType = "gcash" | "paymaya" | "paypal" | null;

const PaymentService = () => {
  const navigate = useNavigate();
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [selectedPayment, setSelectedPayment] = useState<PaymentMethod>(null);
  const [onlinePaymentType, setOnlinePaymentType] = useState<OnlinePaymentType>(null);

  useEffect(() => {
    // Check authentication status
    const checkAuth = () => {
      const token = localStorage.getItem('token'); // or your auth token
      if (!token) {
        navigate('/login', { replace: true });
        return;
      }
      setIsAuthenticated(true);
    };

    checkAuth();
  }, [navigate]);

  // Don't render anything while checking auth
  if (!isAuthenticated) {
    return null;
  }

  const handleBack = () => {
    if (onlinePaymentType) {
      setOnlinePaymentType(null);
    } else if (selectedPayment) {
      setSelectedPayment(null);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        {(selectedPayment || onlinePaymentType) && (
          <button onClick={handleBack} className={styles.backButton}>
            <ArrowLeft className={styles.backIcon} />
          </button>
        )}
        <h2 className={styles.title}>
          {!selectedPayment
            ? "Select Payment Method"
            : selectedPayment === "online" && !onlinePaymentType
            ? "Choose Online Payment"
            : "Payment Details"}
        </h2>
      </div>
          <div className="bg-sky-950 p-4 rounded-lg">
      <h1 className="text-white text-2xl font-bold">Alfred</h1>
    </div>
      <div className={styles.content}>
        {!selectedPayment && (
          <div className={styles.content}>
            <button onClick={() => setSelectedPayment("cod")} className={styles.paymentButton}>
              <Truck className={styles.buttonIcon} />
              <span className={styles.buttonText}>Cash on Delivery</span>
            </button>
            <button onClick={() => setSelectedPayment("credit")} className={styles.paymentButton}>
              <CreditCard className={styles.buttonIcon} />
              <span className={styles.buttonText}>Credit Card</span>
            </button>
            <button onClick={() => setSelectedPayment("online")} className={styles.paymentButton}>
              <Smartphone className={styles.buttonIcon} />
              <span className={styles.buttonText}>Online Payment</span>
            </button>
            <button onClick={() => window.history.back()} className={styles.backToMenuButton}>
              Back to Menu
            </button>
          </div>
        )}

        {selectedPayment === "online" && (
          <div className={styles.onlinePaymentContainer}>
            {["gcash", "paymaya", "paypal"].map((method) => (
              <div key={method} className={styles.paymentMethodWrapper}>
                <button
                  onClick={() => setOnlinePaymentType(method as OnlinePaymentType)}
                  className={`${styles.paymentButton} ${onlinePaymentType === method ? styles.selected : ''}`}
                >
                  <span className={styles.buttonText}>
                    {method.charAt(0).toUpperCase() + method.slice(1)}
                  </span>
                </button>
                {onlinePaymentType === method && (
                  <form className={styles.slideDownForm}>
                    <div className={styles.formGroup}>
                      <label className={styles.label}>Phone Number</label>
                      <input
                        type="tel"
                        placeholder="Enter your phone number"
                        className={styles.input}
                      />
                    </div>
                  </form>
                )}
              </div>
            ))}
          </div>
        )}

        {selectedPayment === "credit" && (
          <form className={styles.form}>
            <div className={styles.formGroup}>
              <label className={styles.label}>Card Number</label>
              <div className={styles.inputContainer}>
                <div className={styles.inputWrapper}>
                  <CreditCard className={styles.inputIcon} />
                  <input
                    type="text"
                    placeholder="1234 5678 9012 3456"
                    className={`${styles.input} ${styles.hasIcon}`}
                    maxLength={19}
                  />
                </div>
              </div>
            </div>
            <div className={styles.cardDetailsGrid}>
              <div className={styles.formGroup}>
                <label className={styles.label}>Expiry Date</label>
                <div className={styles.inputContainer}>
                  <div className={styles.inputWrapper}>
                    <input
                      type="text"
                      placeholder="MM/YY"
                      className={`${styles.input} ${styles.inputNoIcon}`}
                      maxLength={5}
                    />
                  </div>
                </div>
              </div>
              <div className={styles.formGroup}>
                <label className={styles.label}>CVC</label>
                <div className={styles.inputContainer}>
                  <div className={styles.inputWrapper}>
                    <input
                      type="text"
                      placeholder="123"
                      className={styles.input}
                      maxLength={3}
                    />
                  </div>
                </div>
              </div>
            </div>
          </form>
        )}

        {selectedPayment === "cod" && (
          <div className={styles.message}>
            <p>You selected Cash on Delivery</p>
            <p className={styles.messageText}>Payment will be collected upon delivery</p>
          </div>
        )}

        {selectedPayment && (
          <button className={styles.submitButton}>
            Confirm Payment
          </button>
        )}
      </div>
    </div>
  );
};

export default PaymentService;