import React, { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import logo from "../assets/ByteBitesLogo/logo-transparent.png";
import { useApolloClient, useMutation } from "@apollo/client";
import { SIGN_OUT_USER } from "../graphql/Userqueries";
import { supabase } from "../utils/supabaseClient";

const Navbar: React.FC = () => {
    const navigate = useNavigate();
    const isAuthenticated = Boolean(localStorage.getItem("accessToken"));

    const [signOut] = useMutation(SIGN_OUT_USER);
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);
    const [userId, setUserId] = useState<string | null>(null);
    const [logoutInProgress, setLogoutInProgress] = useState(false);
    const client = useApolloClient();

    useEffect(() => {
        const token = localStorage.getItem("accessToken");
        if (token) {
            try {
                const payload = JSON.parse(atob(token.split('.')[1]));
                setUserId(payload.sub || null);
            } catch (e) {
                console.error("Invalid token payload", e);
            }
        }
    }, []);

    const handleLogout = async () => {
        if (logoutInProgress) return;

        setLogoutInProgress(true);

        try {
            const { data } = await signOut();
            if (data?.signOut) {
                localStorage.removeItem("accessToken");
                localStorage.removeItem("refreshToken");
                localStorage.removeItem("expiresAt");

                await supabase.auth.signOut();
                await client.resetStore();

                console.log("User signed out cleanly.");
                navigate("/login");
            } else {
                console.error("SignOut mutation failed");
            }
        } catch (err) {
            console.error("Logout error:", err);
        } finally {
            setLogoutInProgress(false);
        }
    };

    const handleViewAccountClick = () => {
        if (userId) navigate(`/account/${userId}`);
    };

    const handleHomeClick = () => {
        navigate(`/dashboard`);
    };

    return (
        <nav className="bg-blue-950 text-white p-4 shadow-md">
            <div className="container mx-auto flex justify-between items-center">
                <Link to="/" className="flex items-center space-x-2">
                    <img src={logo} alt="Logo" className="h-15 w-15 object-contain" />
                    <span className="text-xl font-semibold text-white">ByteBites</span>
                </Link>
                <ul className="hidden md:flex space-x-6">
                    <li>
                        <button
                            onClick={handleHomeClick}
                            className="block px-4 py-2 text-white hover:bg-gray-600 w-full text-left"
                            >Home
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