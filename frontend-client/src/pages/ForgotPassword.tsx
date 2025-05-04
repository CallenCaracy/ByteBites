import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { FORGOT_PASSWORD } from "../graphql/Userqueries";
import "../styles/main.css";
import logo from "../assets/ByteBitesLogo/logo.png";
import bg from "../assets/forgotpasswordbg.jpg";
import { useMutation } from "@apollo/client";

const ForgotPassword = () => {
    const [email, setEmail] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<{ message: string } | null>(null);
    const [successMessage, setSuccessMessage] = useState<string | null>(null);
    const navigate = useNavigate();
    const [forgotPassword] = useMutation(FORGOT_PASSWORD);

    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        setLoading(true);
        setError(null);
        setSuccessMessage(null);

        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            setError({ message: "Please enter a valid email address." });
            return;
        }

        try {
            const { data } = await forgotPassword({ variables: { email } });
            if (data?.forgotPassword?.success) {
                setSuccessMessage(data?.forgotPassword?.message || "If your email exists, check for the reset link.");
                setTimeout(() => navigate("/login"), 5000); 
            } else {
                setError(data.forgotPassword.message);
            }
        } catch (err: any) {
            console.error("GraphQL mutation error:", err);
            setError({ message: err?.message ?? "Request failed. Please try again." });
        }
    };

    return (
        <div className="py-16 flex justify-center items-center min-h-screen bg-gray-100">
            <div className="flex flex-col lg:flex-row bg-white rounded-lg shadow-lg overflow-hidden w-full max-w-4xl">
                <div className="hidden lg:block w-1/2">
                    <img src={bg} alt="Background" className="w-full h-full object-cover" />
                </div>
                <form onSubmit={handleSubmit} className="flex flex-col justify-center p-8 w-full lg:w-1/2">
                    <div className="flex items-center justify-center mt-1 mb-1 w-full">
                        <img src={logo} alt="Logo" className="h-20 w-20 object-contain" />
                    </div>
                    <h2 className="text-3xl font-semibold text-gray-700 text-center mt-1 mb-1">Forgot Password</h2>
                    <p className="text-lg text-gray-600 text-center mt-1 mb-1">Enter your email to reset your password.</p>
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Email Address</label>
                        <input
                            className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none"
                            type="email"
                            value={email}
                            placeholder="Enter your email"
                            onChange={(e) => setEmail(e.target.value)}
                            required
                            autoComplete="email"
                        />
                    </div>
                    <div className="mt-8">
                        <button
                            type="submit"
                            className="bg-gray-700 text-white font-bold py-2 px-4 w-full rounded hover:bg-gray-600"
                            disabled={loading}
                        >
                            {loading ? "Sending..." : "Send Reset Link"}
                        </button>
                        {successMessage && <p className="text-green-500 bg-green-100 text-sm mt-2 p-2 rounded">{successMessage}</p>}
                    </div>
                    {error && (
                        <p className="text-red-500 text-sm mt-2">
                            {error.message || "An error occurred. Please try again."}
                        </p>
                    )}
                    <p className="text-center text-gray-600 text-sm mt-4">
                        Remembered your password? <span className="text-blue-500 cursor-pointer" onClick={() => navigate("/login")}>Login here</span>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default ForgotPassword;
