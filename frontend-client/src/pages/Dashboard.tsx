import React from "react";
import { useQuery } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { GET_MENU_ITEMS } from "../graphql/Menuqueries";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";
import { useSubscription } from "@apollo/client";
import { MENU_ITEM_CREATED } from "../graphql/Menuqueries";
import Navbar from "../components/NavBar";
import "../styles/main.css";
import placeholderpic from "../assets/placeholder.jpg";

const Dashboard: React.FC = () => {
    const navigate = useNavigate();
    const { data: menuData, loading: menuLoading, error: menuError } = useQuery(GET_MENU_ITEMS);
    const { data: userData, loading: userLoading, error: userError } = useQuery(GET_AUTHENTICATED_USER);

    useSubscription(MENU_ITEM_CREATED, {
        onData: ({ client, data }) => {
          const newItem = data?.data?.menuItemCreated;
          if (!newItem) return;
    
          console.log("Received new menu item:", newItem);
    
          const existing = client.readQuery({ query: GET_MENU_ITEMS });
          if (existing) {
            console.log("Existing menu items:", existing.getAllMenuItems);
          }
    
          if (existing && !existing.getAllMenuItems?.some((item: any) => item.id === newItem.id)) {
            console.log("New item detected, updating cache...");
            client.writeQuery({
              query: GET_MENU_ITEMS,
              data: {
                getAllMenuItems: [newItem, ...existing.getAllMenuItems],
              },
            });
          }
        },
        onError: (error) => {
          console.error("Subscription error:", error);
        },
      });
    

    if (menuLoading || userLoading) return <p className="text-center text-gray-600">Loading...</p>;
    if (menuError || userError) return <p className="text-center text-red-500">Error loading data.</p>;

    console.log("Fetched data:", menuData, userData);

    return (
      <div className="h-screen flex flex-col">
      <Navbar />
      <div className="container p-8 flex-1 overflow-y-auto scrollbar-hide">
          <h1 className="text-3xl font-semibold text-gray-800 mb-6">Menu</h1>
          <h2 className="text-3xl font-semibold text-gray-800 mb-6">
              Welcome {userData?.getAuthenticatedUser?.userType?.charAt(0).toUpperCase() + userData?.getAuthenticatedUser?.userType?.slice(1) || "Employee"} {userData?.getAuthenticatedUser?.firstName || "Unknown"}!
          </h2>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {menuData?.getAllMenuItems?.map((item: any) => (
                  <div
                      key={item.id}
                      className="group bg-white p-4 rounded-lg shadow-md cursor-pointer hover:bg-blue-950 transition"
                      onClick={() => navigate(`/menu-item/${userData?.getAuthenticatedUser?.id}/${item.id}`)}
                  >
                        <div className="flex justify-between items-center">
                          <h2 className="text-xl font-semibold text-gray-700 group-hover:text-white">
                            {item.name}
                          </h2>
                          <p className="text-xl font-semibold text-gray-700 group-hover:text-white">
                            {item?.discount ? `${item.discount.toFixed(2)}% Off` : "No Discount"}
                          </p>
                        </div>
                      <img
                          src={item?.image_url || placeholderpic} 
                          alt="Menu Picture"
                          className="w-128 h-48 rounded-lg object-cover mx-auto mb-3 text-gray-600 group-hover:text-white"
                      />
                      <p className="text-gray-600 group-hover:text-gray-300">
                          {item.description}
                      </p>
                      <div className="flex items-center space-x-4 justify-start">
                        {item.discount > 0 ? (
                            <>
                                <p className="text-lg text-gray-500 line-through">₱{item.price.toFixed(2)}</p>
                                <p className="text-2xl font-bold text-green-700">₱{item.discounted_price.toFixed(2)}</p>
                            </>
                        ) : (
                            <p className="text-2xl font-bold text-green-700">₱{item.price.toFixed(2)}</p>
                        )}
                    </div>
                  </div>
              ))}
          </div>
      </div>
  </div>
);
};

export default Dashboard;
