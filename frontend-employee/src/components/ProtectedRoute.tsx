import React, { useEffect, useState } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { CHECK_TOKEN } from "../graphql/Userqueries";
import { supabase } from "../utils/supabaseClient";

const ProtectedRoute: React.FC = () => {
  const [token, setToken] = useState(localStorage.getItem("accessToken"));
  const [shouldLogout, setShouldLogout] = useState(false);

  useEffect(() => {
    const checkToken = () => setToken(localStorage.getItem("accessToken"));
    window.addEventListener("storage", checkToken);
    return () => window.removeEventListener("storage", checkToken);
  }, []);

  const { data, loading, error } = useQuery(CHECK_TOKEN, {
    skip: !token,
    fetchPolicy: 'network-only',
  });

  useEffect(() => {
    if (error) {
      console.error("Token verification error:", error);
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
      localStorage.removeItem("expiresAt");
      supabase.auth.signOut().then(() => {
        console.log("User signed out due to token error.");
        setShouldLogout(true);
      });
    }
  }, [error]);

  if (shouldLogout || !token) {
    return <Navigate to="/login" replace />;
  }

  if (loading) {
    return <div>Loading...</div>;
  }

  if (data?.checkToken) {
    return <Outlet />;
  }

  return <Navigate to="/login" replace />;
};

export default ProtectedRoute;