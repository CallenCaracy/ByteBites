import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../styles/login.css";
import logo from "../assets/ByteBitesLogo/logo.png";
import bg from "../assets/forgotpasswordbg.jpg";

const ForgotPassword = () => {
    const [email, setEmail] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<{ message: string } | null>(null);
    const navigate = useNavigate();

    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        setLoading(true);
        try {
            // Add logic for sending the reset password email here
            // For example: await sendResetPasswordEmail(email);
            setLoading(false);
            // Optionally navigate back to the login page or show a success message
            navigate("/login");
        } catch (err) {
            setLoading(false);
            setError(err as { message: string });
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
