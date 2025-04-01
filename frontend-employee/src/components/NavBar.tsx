import React from "react";
import { Link, useNavigate } from "react-router-dom";
import logo from "../assets/ByteBitesLogo/logo.png";

const Navbar: React.FC = () => {
    const navigate = useNavigate();
    const isAuthenticated = Boolean(localStorage.getItem("accessToken"));

    const handleLogout = () => {
        localStorage.removeItem("accessToken");
        localStorage.removeItem("refreshToken");
        navigate("/login");
    };

    return (
        <nav className="bg-gray-800 text-white p-4 shadow-md">
            <div className="container mx-auto flex justify-between items-center">
                <Link to="/" className="flex items-center space-x-2">
                    <img src={logo} alt="Logo" className="h-10 w-10 object-contain" />
                    <span className="text-xl font-semibold">ByteBites</span>
                </Link>
                <ul className="hidden md:flex space-x-6">
                    <li><Link to="/dashboard" className="hover:text-gray-400">Home</Link></li>
                    <li><Link to="/menu" className="hover:text-gray-400">Menu</Link></li>
                    <li><Link to="/about" className="hover:text-gray-400">About</Link></li>
                    <li><Link to="/account" className="hover:text-gray-400">Account</Link></li>
                    {isAuthenticated ? (
                        <li>
                            <button 
                                onClick={handleLogout} 
                                className="bg-red-500 px-4 py-2 rounded hover:bg-red-600"
                            >
                                Logout
                            </button>
                        </li>
                    ) : (
                        <li><Link to="/login" className="hover:text-gray-400">Login</Link></li>
                    )}
                </ul>
                <div className="md:hidden">
                    <button className="focus:outline-none">☰</button>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;