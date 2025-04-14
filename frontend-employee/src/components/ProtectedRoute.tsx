import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { Check_Token } from "../graphql/Userqueries";

const ProtectedRoute: React.FC = () => {
  const token = localStorage.getItem("accessToken");
  console.log("Check", token)
  
  const { data, loading, error } = useQuery(Check_Token, {
    variables: { token },
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
    return <Navigate to="/login" replace />;
  }

  if (data && data.checkToken) {
    return <Outlet />;
  }

  return <Navigate to="/login" replace />;
};

export default ProtectedRoute;
