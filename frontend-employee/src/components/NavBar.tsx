import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import logo from "../assets/ByteBitesLogo/logo-transparent.png";
import { useApolloClient, useMutation } from "@apollo/client";
import { SIGN_OUT_USER } from "../graphql/Userqueries";

const Navbar: React.FC = () => {
    const navigate = useNavigate();
    const isAuthenticated = Boolean(localStorage.getItem("accessToken"));

    const [signOut, { loading, error }] = useMutation(SIGN_OUT_USER);

    const [isDropdownOpen, setIsDropdownOpen] = useState(false);

    const [userId, setUserId] = useState<string | null>(null);

    const client = useApolloClient();

    useEffect(() => {
        const token = localStorage.getItem("accessToken");
        if (token) {
            const payloadRaw = token.split('.')[1];
            const decoded = atob(payloadRaw);
            console.log("Decoded JWT payload string:", decoded);
    
            try {
                const payload = JSON.parse(decoded);
                console.log("Parsed payload object:", payload);
    
                if (payload.sub) {
                    setUserId(payload.sub);
                } else {
                    console.warn("No userId in payload! Check the token structure.");
                }
            } catch (e) {
                console.error("Failed to parse JWT payload:", e);
            }
        }
    }, []);
    

    const handleLogout = async () => {
        try {
            const { data } = await signOut();
            if (data?.signOut) {
                if (loading) return <p className="text-center text-gray-600">Signing out...</p>;
                if (error) return <p className="text-center text-red-500">Error signing out.</p>;
                console.log("Logout successful");

                localStorage.removeItem("accessToken");
                localStorage.removeItem("refreshToken");

                client.resetStore();
                navigate("/login");
            } else {
                console.error("Logout failed");
            }
        } catch (err) {
            console.error("Error signing out:", err);
        }
    }

    const handleViewAccountClick = () => {
        if (userId) {
            navigate(`/account/${userId}`);
        }
    };

    const handleHomeClick = () => {
        if (userId) {
            navigate(`/dashboard`);
        }
    };

    return (
        <nav className="bg-gray-800 text-white p-4 shadow-md">
            <div className="container mx-auto flex justify-between items-center">
                <Link to="/" className="flex items-center space-x-2">
                    <img src={logo} alt="Logo" className="h-15 w-15 object-contain" />
                    <span className="text-xl font-semibold">ByteBites</span>
                </Link>
                <ul className="hidden md:flex space-x-6">
                    <li>
                        <button
                            onClick={handleHomeClick}
                            className="block px-4 py-2 text-white hover:bg-gray-600 w-full text-left"
                            >Home
                        </button>
                    </li>
                    <li>
                        <button
                            onClick={() => navigate('/payment')}
                            className="block px-4 py-2 text-white hover:bg-gray-600 w-full text-left"
                            >Proceed to Payment
                        </button>
                    </li>
                    {isAuthenticated ? (
                        <li className="relative">
                            <button
                                onClick={() => setIsDropdownOpen(!isDropdownOpen)}
                                className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
                            >
                                Account
                            </button>
                            {isDropdownOpen && (
                                <div className="absolute left-0 mt-2 w-48 bg-gray-800 rounded shadow-lg z-10">
                                    <ul>
                                        <li>
                                            <button
                                                onClick={handleViewAccountClick}
                                                className="block px-4 py-2 text-white hover:bg-gray-600 w-full text-left"
                                            >
                                                Account
                                            </button>
                                        </li>
                                        <li>
                                            <button
                                                onClick={handleLogout}
                                                className="block px-4 py-2 text-white hover:bg-gray-600 w-full text-left"
                                            >
                                                Logout
                                            </button>
                                        </li>
                                    </ul>
                                </div>
                            )}
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