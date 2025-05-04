import React, { useState } from "react";
import { useQuery, useMutation } from "@apollo/client";
import {
    GET_MENU_ITEM_BY_ID,
    DELETE_MENU_ITEM,
    UPDATE_MENU_ITEM,
} from "../graphql/Menuqueries";
import { useNavigate, useParams } from "react-router-dom";
import { GET_INVENTORY } from "../graphql/Kitchenqueries";
import Navbar from "../components/NavBar";

const MenuItem: React.FC = () => {
    const { menuId } = useParams();
    const navigate = useNavigate();
    const [isUpdateModalOpen, setIsUpdateModalOpen] = useState(false);
    const [formData, setFormData] = useState({
        name: "",
        description: "",
        price: 0,
        category: "",
        availability_status: "available",  // Default value for availability
        image_url: "",
    });

    const { data, loading, error } = useQuery(GET_MENU_ITEM_BY_ID, {
        variables: { id: menuId },
        skip: !menuId,
        onCompleted: (data) => {
            const item = data.getMenuItemById;
            setFormData({
                name: item.name,
                description: item.description,
                price: item.price,
                category: item.category,
                availability_status: item.availability_status ? "available" : "unavailable",
                image_url: "https://your-static-image-url.com/your-image.jpg",
            });
        },
    });

    const [deleteMenuItem] = useMutation(DELETE_MENU_ITEM);
    const [updateMenuItem] = useMutation(UPDATE_MENU_ITEM);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching menu data: {error.message}</p>;
    if (!data || !data.getMenuItemById) return <p>No menu data found.</p>;

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

    const handleFormChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
    ) => {
        const { name, value } = e.target;

        setFormData((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            setFormData((prev) => ({
                ...prev,
                image_url: URL.createObjectURL(file),
            }));
        }
    };

    const handleUpdate = async (e: React.FormEvent) => {
        e.preventDefault();
    
        // Input validation
        const newErrors: { [key: string]: string } = {};
        if (!formData.name) newErrors.name = "Name is required.";
        if (!formData.description) newErrors.description = "Description is required.";
        if (!formData.price) newErrors.price = "Price is required.";
        else if (parseFloat(formData.price.toString()) < 0) newErrors.price = "Price cannot be negative.";
        if (!formData.category) newErrors.category = "Category is required.";
        if (!formData.availability_status) newErrors.availability_status = "Availability status is required.";
    
        // If there are errors, do not submit
        if (Object.keys(newErrors).length > 0) {
            alert(Object.values(newErrors).join("\n"));
            return;
        }
    
        try {
            await updateMenuItem({
                variables: {
                    id: menuId,
                    input: {
                        name: formData.name,
                        description: formData.description,
                        price: parseFloat(formData.price.toString()),
                        category: formData.category,
                        availability_status: formData.availability_status === "available",
                        image_url: formData.image_url,
                    },
                },
            });
    
            setIsUpdateModalOpen(false);
            window.location.reload();
        } catch (err) {
            alert("Failed to update menu item.");
        }
    };
    
    

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
              <button
                            onClick={() => setIsUpdateModalOpen(true)}
                            className="bg-black text-yellow-500 px-4 py-2 rounded hover:bg-gray-800"
                        >
                            Update
                        </button>
                            // ! here is button for delettion
                                                            <button
                                onClick={async () => {
                                    try {
                                    const response = await deleteMenuItem({ variables: { id: menuId } });
                                    if (response.data.deleteMenuItem) {
                                        alert("Menu item deleted successfully.");
                                        navigate("/dashboard");
                                    } else {
                                        alert("Failed to delete menu item.");
                                    }
                                    } catch (err) {
                                    alert("Failed to delete menu item.");
                                    }
                                }}
                                className="bg-black text-red-500 px-4 py-2 rounded hover:bg-gray-800"
                                >
                                Delete
                            </button>
            </div>
            {isUpdateModalOpen && (
                <div className="fixed inset-0 flex items-center justify-center bg-gray-100 bg-opacity-90 z-50">
                    <div className="bg-white p-6 rounded w-[400px] relative">
                        <h2 className="text-xl font-bold mb-4 text-black">Update Menu Item</h2>

                        <form onSubmit={handleUpdate}>
                            <input
                                type="text"
                                name="name"
                                placeholder="Item Name"
                                value={formData.name}
                                onChange={handleFormChange}
                                required
                                className="border p-2 w-full mb-2 text-black"
                            />
                            <textarea
                                name="description"
                                placeholder="Description"
                                value={formData.description}
                                onChange={handleFormChange}
                                required
                                className="border p-2 w-full mb-2 text-black"
                            />
                            <input
                                type="number"
                                name="price"
                                placeholder="Price"
                                value={formData.price}
                                onChange={handleFormChange}
                                required
                                className="border p-2 w-full mb-2 text-black"
                            />
                            <input
                                type="text"
                                name="category"
                                placeholder="Category"
                                value={formData.category}
                                onChange={handleFormChange}
                                className="border p-2 w-full mb-2 text-black"
                            />

                            <select
                                name="availability_status"
                                value={formData.availability_status}
                                onChange={handleFormChange}
                                className="border p-2 w-full mb-2 text-black"
                            >
                                <option value="available">Available</option>
                                <option value="unavailable">Unavailable</option>
                            </select>

                            <input
                                type="file"
                                name="image_url"
                                onChange={handleImageChange}
                                accept="image/*"
                                className="border p-2 w-full mb-2 text-black"
                            />
                            
                           

                            <div className="flex justify-end gap-2">
                                <button
                                    type="button"
                                    onClick={() => setIsUpdateModalOpen(false)}
                                    className="text-red-500"
                                >
                                    Cancel
                                </button>
                                <button
                                    type="submit"
                                    className="bg-green-600 text-white px-4 py-2 rounded"
                                >
                                    Save
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
          </div>
        </>
      );
  };
  
  export default MenuItem; 
