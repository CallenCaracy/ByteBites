import { useNavigate, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEM_BY_ID } from "../graphql/Menuqueries";
import { GET_INVENTORY } from "../graphql/Kitchenqueries";
import Navbar from "../components/NavBar";

const MenuItem: React.FC = () => {
    const { menuId } = useParams();
    const navigate = useNavigate();

    const { data: inventoryData, loading: inventoryLoading, error: inventoryError } = useQuery(GET_INVENTORY, {
    variables: { menuId: menuId }
    });

    const { data: menuData, loading: menuLoading, error: menuError } = useQuery(GET_MENU_ITEM_BY_ID, {
    variables: { id: menuId },
    skip: !menuId,
    });

    if (menuLoading) return <p>Loading menu item...</p>;
    if (menuError) return <p>Error fetching menu item: {menuError.message}</p>;

    if (inventoryLoading) return <p>Loading inventory...</p>;
    if (inventoryError) return <p>Error fetching inventory: {inventoryError.message}</p>;

    if (!menuData || !menuData.getMenuItemById || !inventoryData || !inventoryData.inventory) {
    return <p>No data found.</p>;
    }

    const item = menuData.getMenuItemById;
    const inventory = inventoryData.inventory;

    return (
        <>
          <Navbar />
          <div className="p-8 max-w-4xl mx-auto bg-white rounded-xl shadow-[0_4px_10px_0_rgba(0,0,0,0.5)] space-y-6 mt-10">
      
            {/* Back Button */}
            <button
              onClick={() => navigate("/dashboard")}
              className="flex items-center text-white focus:outline-none focus:ring-2 focus:ring-blue-500 px-4 py-2 rounded-md text-sm"
            >
              <svg xmlns="http://www.w3.org/2000/svg" className="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth="2">
                <path strokeLinecap="round" strokeLinejoin="round" d="M15 19l-7-7 7-7" />
              </svg>
              Back
            </button>
      
            {/* Full-width Image */}
            <img
              src={item.image_url}
              alt={item.name}
              className="w-full rounded-xl object-cover h-72 shadow-md"
            />
      
            {/* Info Section: Two Columns */}
            <div className="grid md:grid-cols-2 gap-8 mt-6">
              
              {/* Left: Text Info */}
              <div className="space-y-3">
                <h2 className="text-gray-900 text-3xl font-semibold">{item.name}</h2>
                <p className="text-gray-700 text-lg">{item.description}</p>
                <p className="text-sm text-gray-600">
                  <span className="font-medium">Category:</span> {item.category}
                </p>
                <p className="text-sm text-gray-600">
                  <span className="font-medium">Status:</span>{" "}
                  <span className={item.availability_status ? "text-green-600" : "text-red-600"}>
                    {item.availability_status ? "Available" : "Out of stock"}
                  </span>
                </p>
              </div>
      
              {/* Right: Price and Inventory */}
              <div className="space-y-3 text-right">
                <p className="text-2xl font-bold text-green-700">
                  ₱{item.price.toFixed(2)}
                </p>
                <p className="text-sm text-gray-600">
                  Quantity Available:{" "}
                  <span className={`${inventory.availableServings <= inventory.lowStockThreshold ? "text-red-500" : "text-green-600"}`}>
                    {inventory.availableServings}
                  </span>
                </p>
                <p className="text-sm text-gray-600">
                  Low Stock Threshold:{" "}
                  <span className="text-yellow-500">{inventory.lowStockThreshold}</span>
                </p>
              </div>
            </div>
      
            {/* Action Buttons: Centered Below */}
            <div className="flex justify-center gap-4 pt-6">
              <button className="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:ring-2 focus:ring-blue-400">
                Update
              </button>
              <button className="px-6 py-2 bg-red-600 text-white rounded-md hover:bg-red-700 focus:ring-2 focus:ring-red-400">
                Delete
              </button>
            </div>
          </div>
        </>
      );
  };
  
  export default MenuItem;  