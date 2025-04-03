import React, { useState } from "react";
import { useMutation } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { SIGN_UP_MUTATION } from "../graphql/Userqueries";
import "../styles/login.css";
import logo from "../assets/ByteBitesLogo/logo.png";
import bg from "../assets/loginwallpaper.jpg";

const RegisterPage: React.FC = () => {
    const [formData, setFormData] = useState({
        email: "",
        firstName: "",
        lastName: "",
        role: "employee",
        address: "",
        phone: "",
        password: "",
        confirmPassword: "",
    });
    const navigate = useNavigate();
    const [signUp, { loading, error }] = useMutation(SIGN_UP_MUTATION);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;

        if (name === "phone") {
            if (!/^\d*$/.test(value)) return;
            if (value.length > 11) return;
        }
    
        if (name === "role") return;
    
        setFormData({ ...formData, [name]: value });
    };

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if (formData.password !== formData.confirmPassword) {
            alert("Passwords do not match!");
            return;
        }

        if (formData.phone.length !== 11) {
            alert("Phone number must be exactly 11 digits long.");
            return;
        }

        if (formData.password.length < 8) {
            alert("Password must be at least 8 characters long.");
            return;
        }

        try {
            const { data } = await signUp({
                variables: {
                    input: {
                        email: formData.email,
                        password: formData.password,
                        firstName: formData.firstName,
                        lastName: formData.lastName,
                        role: formData.role,
                        address: formData.address,
                        phone: formData.phone,
                    },
                },
            });
    
            if (data?.signUp?.id) {
                alert("Registration successful!");
                navigate("/login");
            } else {
                alert("Registration failed.");
            }
        } catch (err) {
            console.error("Registration error:", err);
            if ((err as any)?.graphQLErrors && (err as any).graphQLErrors.length > 0) {
                (err as any).graphQLErrors.forEach((e: any) => console.error("GraphQL error:", e.message));
            }
            alert("Something went wrong. Please try again.");
        }
    };

    return (
        <div className="py-16 flex justify-center items-center min-h-screen bg-gray-100">
            <div className="flex flex-col lg:flex-row bg-white rounded-lg shadow-lg overflow-hidden w-full max-w-4xl">
                <div className="hidden lg:block w-1/2">
                    <img src={bg} alt="Background" className="w-full h-full object-cover" />
                </div>
                <form onSubmit={handleRegister} className="flex flex-col justify-center p-8 w-full lg:w-1/2">
                    <div className="flex items-center justify-center mt-1 mb-1 w-full">
                        <img src={logo} alt="Logo" className="h-20 w-20 object-contain" />
                    </div>
                    <h2 className="text-3xl font-semibold text-gray-700 text-center mt-1 mb-1">ByteBites</h2>
                    <p className="text-lg text-gray-600 text-center mt-1 mb-1">Hello Employee!<br></br>Let's get started.</p>
                    {Object.keys(formData).map((key) => (
                        key !== "confirmPassword" && (
                        <div key={key} className="mt-4">
                            <label className="block text-gray-700 text-sm font-bold mb-2">
                                {key.charAt(0).toUpperCase() + key.slice(1)}
                            </label>
                            <input
                                className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none"
                                type={key === "password" ? "password" : "text"}
                                name={key}
                                value={formData[key as keyof typeof formData]}
                                placeholder={`Enter your ${key}`}
                                onChange={handleChange}
                                required
                                autoComplete="new-password"
                            />
                        </div>)
                    ))}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Confirm Password</label>
                        <input
                            className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none"
                            type="password"
                            name="confirmPassword"
                            value={formData.confirmPassword}
                            placeholder="Confirm your password"
                            onChange={handleChange}
                            required
                            autoComplete="new-password"
                        />
                    </div>
                    <div className="mt-8">
                        <button
                            type="submit"
                            className="bg-gray-700 text-white font-bold py-2 px-4 w-full rounded hover:bg-gray-600"
                            disabled={loading}
                        >
                            {loading ? "Registering..." : "Sign Up"}
                        </button>
                    </div>
                    {error && (
                        <p className="text-red-500 text-sm mt-2">
                            {error.message.includes("duplicate key value violates unique constraint") 
                                ? "This email is already registered. Try logging in." 
                                : "Registration failed. Please try again."
                            }
                        </p>
                    )}
                    <p className="text-center text-gray-600 text-sm mt-4">
                        Already have an account? <span className="text-blue-500 cursor-pointer" onClick={() => navigate("/login")}>Login here</span>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default RegisterPage;