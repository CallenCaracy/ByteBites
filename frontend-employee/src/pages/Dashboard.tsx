import React, { useState } from "react";
import { useQuery } from "@apollo/client";
import { useNavigate } from "react-router-dom";
import { GET_MENU_ITEMS } from "../graphql/Menuqueries";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";
import Navbar from "../components/NavBar";



const Dashboard: React.FC = () => {
    const navigate = useNavigate();
    const { data: menuData, loading: menuLoading, error: menuError } = useQuery(GET_MENU_ITEMS);
    const { data: userData, loading: userLoading, error: userError } = useQuery(GET_AUTHENTICATED_USER);
    
    const [form, setForm] = useState({
        name: "",
        description: "",
        price: "",
        category: "",
        availabilityStatus: "available",
        image: null,
        imageUrl: "",
    });

    const [isModalOpen, setIsModalOpen] = useState(false);

    if (menuLoading || userLoading) return <p className="text-center text-gray-600">Loading...</p>;
    if (menuError || userError) return <p className="text-center text-red-500">Error loading data.</p>;

    console.log("Fetched data:", menuData, userData);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setForm((prev) => ({ ...prev, [name]: value }));
    };

    const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            const reader = new FileReader();
            reader.onloadend = () => {
                setForm((prev) => ({ ...prev, imageUrl: reader.result as string }));
            };
            reader.readAsDataURL(file);
        }
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        // Handle form submission logic here
    };


  


      

    return (
        <div>
            <Navbar />
            <div className="container mx-auto p-8">
                <h1 className="text-3xl font-semibold text-gray-800 mb-6">Menu</h1>
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">
                    Welcome {userData?.getAuthenticatedUser?.userType?.charAt(0).toUpperCase() + userData?.getAuthenticatedUser?.userType?.slice(1) || "Employee"} {userData?.getAuthenticatedUser?.firstName || "Unknown"}!
                </h2>

                

                <button
                    onClick={() => setIsModalOpen(true)}
                    className="bg-blue-600 text-white px-4 py-2 rounded"
                >
                    + Add Menu Item
                </button>

                {isModalOpen && (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-100 bg-opacity-10000"> {/* Changed background color */}
        <div className="bg-white p-6 rounded w-[400px] relative">
            <h2 className="text-xl font-bold mb-4 text-black">Add Menu Item</h2>

            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    name="name"
                    placeholder="Item Name"
                    value={form.name}
                    onChange={handleChange}
                    required
                    className="border p-2 w-full mb-2 text-black"
                />
                <input
                    type="text"
                    name="description"
                    placeholder="Description"
                    value={form.description}
                    onChange={handleChange}
                    required
                    className="border p-2 w-full mb-2 text-black"
                />
                <input
                    type="number"
                    name="price"
                    placeholder="Price"
                    value={form.price}
                    onChange={handleChange}
                    required
                    className="border p-2 w-full mb-2 text-black"
                />
                <input
                    type="text"
                    name="category"
                    placeholder="Category"
                    value={form.category}
                    onChange={handleChange}
                    className="border p-2 w-full mb-2 text-black"
                />
                <select
                    name="availabilityStatus"
                    value={form.availabilityStatus}
                    onChange={handleChange}
                    className="border p-2 w-full mb-2 text-black"
                >
                    <option value="available">Available</option>
                    <option value="unavailable">Unavailable</option>
                </select>
                <input
                    type="file"
                    accept="image/*"
                    onChange={handleImageChange}
                    className="border p-2 w-full mb-4 text-black"
                />
                <div className="flex justify-end gap-2">
                    <button type="submit" className="bg-green-600 text-white px-4 py-2 rounded">
                        Save
                    </button>
                    <button onClick={() => setIsModalOpen(false)} className="text-red-500">
                        Cancel
                    </button>
                </div>
            </form>
        </div>
    </div>
)}




                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {menuData?.getAllMenuItems?.map((item: any) => (
                        <div 
                            key={item.id} 
                            className="bg-white p-4 rounded-lg shadow-md cursor-pointer hover:bg-gray-50 transition"
                            onClick={() => navigate(`/menu-item/${item.id}`)}
                        >
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
