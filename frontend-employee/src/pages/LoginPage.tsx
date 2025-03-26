import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import "../styles/login.css";
import logo from '../assets/ByteBitesLogo/logo.png';
import bg from '../assets/loginwallpaper.jpg';

const LoginPage: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        navigate('/dashboard');
    };

    return (
        <div className="py-16 flex justify-center items-center min-h-screen bg-gray-100">
            <div className="flex flex-col lg:flex-row bg-white rounded-lg shadow-lg overflow-hidden w-full max-w-4xl"> 
            <div className="hidden lg:block w-1/2">
                <img src={bg} alt="Background" className="w-full h-full object-cover" />
            </div>
                <form onSubmit={handleLogin} className="flex flex-col justify-center p-8 w-full lg:w-1/2">
                <div className="flex items-center justify-center mt-1 mb-1 w-full">
                    <img src={logo} alt="Logo" className="h-20 w-20 object-contain" />
                </div>
                <h2 className="text-3xl font-semibold text-gray-700 text-center mt-1 mb-1">ByteBites</h2>
                <p className="text-lg text-gray-600 text-center mt-1 mb-1">Welcome back!</p>
                    <div className="mt-4">
                        <label className="block text-gray-700 text-sm font-bold mb-2">Email Address</label>
                        <input className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none" 
                            type="email"
                            value={email}
                            placeholder="Enter your email"
                            onChange={(e) => setEmail(e.target.value)}
                            required />
                    </div>
                    <div className="mt-4">
                        <div className="flex justify-between">
                            <label className="block text-gray-700 text-sm font-bold mb-2">Password</label>
                            <a href="#" className="text-xs text-gray-500">Forget Password?</a>
                        </div>
                        <input className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none" 
                            type="password"
                            value={password}
                            placeholder="Enter your password"
                            onChange={(e) => setPassword(e.target.value)}
                            required />
                    </div>
                    <div className="mt-8">
                        <button className="bg-gray-700 text-white font-bold py-2 px-4 w-full rounded hover:bg-gray-600">Login</button>
                    </div>
                    <div className="mt-4 flex items-center justify-between">
                        <span className="border-b w-1/5 md:w-1/4"></span>
                        <span className="text-xs text-gray-500 uppercase">or sign up</span>
                        <span className="border-b w-1/5 md:w-1/4"></span>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default LoginPage;