import React, { useState } from "react";
import { useMutation } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { SIGN_UP_MUTATION } from "../graphql/Userqueries";
import "../styles/main.css";
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
        birthDate: "",
        userType: "",
        pfp: "https://hzjjmfwrtvqjwxunfcue.supabase.co/storage/v1/object/public/pictures/pfp/defualtpic.jpg",
        gender: "",
    });    
    const navigate = useNavigate();
    const [signUp, { loading, error }] = useMutation(SIGN_UP_MUTATION);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;

        if (name === "email") {
            const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
            setEmailError(emailRegex.test(value) ? "" : "Invalid email format");
        }

        if (name === "phone") {
            if (!/^\d*$/.test(value)) return;
            if (value.length > 11) return;
        }
    
        if (name === "role") return;       
    
        setFormData({ ...formData, [name]: value });
    };

    const [emailError, setEmailError] = useState("");

    const validateEmail = () => {
        const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
        
        if (!emailRegex.test(formData.email)) {
            setEmailError("Invalid email format");
        } else {
            setEmailError("");
        }
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
                        birthDate: formData.birthDate,
                        userType: formData.userType,
                        pfp: formData.pfp,
                        gender: formData.gender,
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
                    
                    {/* Email */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Email</label>
                        <input
                            type="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            onBlur={validateEmail}
                            required
                            placeholder="Enter your email"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        />
                        <p className="text-red-500 text-sm">{emailError}</p>
                    </div>

                    {/* First Name */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">First Name</label>
                        <input
                            type="text"
                            name="firstName"
                            value={formData.firstName}
                            onChange={handleChange}
                            required
                            placeholder="Enter your First Name"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        />
                    </div>

                    {/* Last Name */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Last Name</label>
                        <input
                            type="text"
                            name="lastName"
                            value={formData.lastName}
                            onChange={handleChange}
                            required
                            placeholder="Enter your Last Name"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        />
                    </div>

                    {/* Address */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Address</label>
                        <input
                            type="text"
                            name="address"
                            value={formData.address}
                            onChange={handleChange}
                            placeholder="Enter your Address"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        />
                    </div>

                    {/* Phone */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Phone</label>
                        <input
                            type="text"
                            name="phone"
                            value={formData.phone}
                            onChange={handleChange}
                            maxLength={11}
                            placeholder="Enter your phone number"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                            autoComplete="tel"
                        />
                    </div>

                    {/* Age */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Age</label>
                        <input
                            type="date"
                            name="birthDate"
                            value={formData.birthDate}
                            onChange={handleChange}
                            required
                            placeholder="Enter your birth date"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        />
                    </div>

                    {/* User Type */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">User Type</label>
                        <select
                            name="userType"
                            value={formData.userType}
                            onChange={handleChange}
                            required
                            title="Select User Type"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        >
                            <option value="">Select user type</option>
                            <option value="staff">Staff</option>
                            <option value="manager">Manager</option>
                            <option value="chef">Chef</option>
                        </select>
                    </div>

                    {/* Gender */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Gender</label>
                        <select
                            name="gender"
                            value={formData.gender}
                            onChange={handleChange}
                            title="Select Gender"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                        >
                            <option value="">Select Gender</option>
                            <option value="male">Male</option>
                            <option value="female">Female</option>
                            <option value="other">Other</option>
                        </select>
                    </div>

                    {/* Password */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Password</label>
                        <input
                            type="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                            placeholder="Enter your password"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
                            autoComplete="new-password"
                        />
                    </div>

                    {/* Confirm Password */}
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Confirm Password</label>
                        <input
                            type="password"
                            name="confirmPassword"
                            value={formData.confirmPassword}
                            onChange={handleChange}
                            required
                            placeholder="Enter your confirm password"
                            className="bg-gray-200 text-gray-700 border border-gray-300 rounded py-2 px-4 w-full"
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
                            {error.message.includes("User already registered") 
                                ? "This email is already registered. Try logging in." 
                                : "Registration failed. Please try again."
                            }
                        </p>
                    )}

                    <p className="text-</p>center text-gray-600 text-sm mt-4">
                        Already have an account? <span className="text-blue-500 cursor-pointer" onClick={() => navigate("/login")}>Login here</span>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default RegisterPage;
