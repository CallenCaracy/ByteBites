import React, { useEffect, useState } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { CHECK_TOKEN } from "../graphql/Userqueries";

const ProtectedRoute: React.FC = () => {
  const [token, setToken] = useState(localStorage.getItem("accessToken"));
  console.log("Check", token)
  
  useEffect(() => {
    const checkToken = () => setToken(localStorage.getItem("accessToken"));
    window.addEventListener("storage", checkToken);
    return () => window.removeEventListener("storage", checkToken);
  }, []);

  const { data, loading, error } = useQuery(CHECK_TOKEN, {
    skip: !token,
    fetchPolicy: 'network-only',
  });

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    console.error("Token verification error:", error);
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
    return <Navigate to="/login" replace />;
  }

  if (data && data.checkToken) {
    return <Outlet />;
  }

  return <Navigate to="/login" replace />;
};

export default ProtectedRoute;