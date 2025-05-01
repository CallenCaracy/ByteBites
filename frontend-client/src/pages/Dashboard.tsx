import React from "react";
import { useQuery } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { GET_MENU_ITEMS } from "../graphql/Menuqueries";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";
import Navbar from "../components/NavBar";
import placeholder from "../assets/placeholder.jpg";

const Dashboard: React.FC = () => {
    const navigate = useNavigate();
    const { data: menuData, loading: menuLoading, error: menuError } = useQuery(GET_MENU_ITEMS);
    const { data: userData, loading: userLoading, error: userError } = useQuery(GET_AUTHENTICATED_USER);
    

    if (menuLoading || userLoading) return <p className="text-center text-gray-600">Loading...</p>;
    if (menuError || userError) return <p className="text-center text-red-500">Error loading data.</p>;

    console.log("Fetched data:", menuData, userData);

    return (
        <div>
            <Navbar />
            <div className="container mx-auto p-8">
                <h1 className="text-3xl font-semibold text-gray-800 mb-6">Menu</h1>
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">
                    Welcome {userData?.getAuthenticatedUser?.firstName || "Unknown"}!
                </h2>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {menuData?.getAllMenuItems?.map((item: any) => (
                        <div 
                            key={item.id} 
                            className="bg-white p-4 rounded-lg shadow-[0_4px_10px_0_rgba(0,0,0,0.4)] cursor-pointer hover:bg-gray-50 transition"
                            onClick={() => navigate(`/menu-item/${userData?.getAuthenticatedUser?.id}/${item.id}`)}
                        >
                            <h2 className="text-xl font-semibold text-gray-700 mb-2">{item.name}</h2>
                            <img
                                src={placeholder}
                                alt="Menu Picture"
                                className="w-128 h-48 rounded-lg object-cover mx-auto mb-3"
                            />
                            <p className="text-xl font-semibold text-green-700 whitespace-nowrap">
                                ₱{item.price.toFixed(2)}
                            </p>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Dashboard;
