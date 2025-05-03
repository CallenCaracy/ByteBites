import React from "react";
import { useParams } from "react-router-dom";

const PaymentPage: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="bg-white p-8 rounded-lg shadow-lg text-center">
                <h1 className="text-2xl font-bold mb-4">Payment Page</h1>
                <p className="text-gray-600">Order ID: <span className="font-mono">{userId}</span></p>
                <p className="mt-4 text-green-600">This is just a placeholder. Hook up your payment logic here.</p>
            </div>
        </div>
    );
};

export default PaymentPage;
