import React from "react";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEMS } from "../graphql/Menuqueries";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";
import Navbar from "../components/NavBar";

const Dashboard: React.FC = () => {
    const { data: menuData, loading: menuLoading, error: menuError } = useQuery(GET_MENU_ITEMS);
    const { data: userData, loading: userLoading, error: userError } = useQuery(GET_AUTHENTICATED_USER);
    

    if (menuLoading || userLoading) return <p className="text-center text-gray-600">Loading...</p>;
    if (menuError || userError) return <p className="text-center text-red-500">Error loading data.</p>;

    console.log("Fetched data:", menuData, userData);

    return (
        <div>
            <Navbar userId={userData?.getAuthenticatedUser?.id} />
            <div className="container mx-auto p-8">
                <h1 className="text-3xl font-semibold text-gray-800 mb-6">Menu</h1>
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">
                    Welcome { userData?.getAuthenticatedUser?.userType?.charAt(0).toUpperCase() + userData?.getAuthenticatedUser?.userType?.slice(1) || "Employee"} {userData?.getAuthenticatedUser?.firstName || "Unknown"}!
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {menuData?.getAllMenuItems?.map((item: any) => (
                        <div key={item.id} className="bg-white p-4 rounded-lg shadow-md">
                            <h2 className="text-xl font-semibold text-gray-700">{item.name}</h2>
                            <p className="text-gray-600">{item.description}</p>
                            <p className="text-gray-900 font-bold">${item.price.toFixed(2)}</p>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
