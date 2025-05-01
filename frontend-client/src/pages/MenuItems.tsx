import { useNavigate, useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEM_BY_ID } from "../graphql/Menuqueries";
import Navbar from "../components/NavBar";
import placeholder from "../assets/placeholder.jpg";

const MenuItem: React.FC = () => {
    const { menuId, userId } = useParams();
    const navigate = useNavigate();

    const { data, loading, error } = useQuery(GET_MENU_ITEM_BY_ID, {
        variables: { id: menuId },
        skip: !menuId,
    });

    const [orderData, setOrderData] = useState(() => ({
        orderId: userId || "",
        menuItemId: menuId || "",
        quantity: 1,
        price: 0,
        customizations: [] as string[],
    }));

    const [availableCustomizations, setAvailableCustomizations] = useState<string[]>([]);

    useEffect(() => {
        if (data && data.getMenuItemById) {
            const item = data.getMenuItemById;
            setOrderData((prev) => ({
                ...prev,
                price: item.price || 0,
            }));

            setAvailableCustomizations(["Spicy", "Extra Sauce", "No Sauce"]);
        }
    }, [data]);

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching menu data: {error.message}</p>;
    if (!data || !data.getMenuItemById) return <p>No menu data found.</p>;

    const item = data.getMenuItemById;

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setOrderData((prev) => ({ ...prev, [name]: Number(value) }));
    };

    const handleCheckboxChange = (option: string) => {
        setOrderData((prev) => {
            const current = Array.isArray(prev.customizations) ? prev.customizations : [];
            const updated = current.includes(option)
                ? current.filter((o) => o !== option)
                : [...current, option];
            return { ...prev, customizations: updated };
        });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log("Order Data submitted:", orderData);
        navigate("/dashboard"); // or any other page
    };

    return (
        <>
            <Navbar />
            <div className="p-8 max-w-200 mx-auto bg-white rounded-2xl shadow-[0_4px_10px_0_rgba(0,0,0,0.5)] space-y-5 mt-10">
                <img
                    src={placeholder}
                    alt={item.name}
                    className="w-full rounded-xl object-cover h-64"
                />

                <div className="flex justify-between items-center">
                    <h2 className="text-gray-900 text-2xl font-bold">{item.name}</h2>
                    <p className="text-xl font-semibold text-green-700 whitespace-nowrap">
                        ₱{item.price.toFixed(2)}
                    </p>
                </div>

                <p className="text-gray-600">{item.description}</p>
                <p className="text-sm text-gray-500">Category: {item.category}</p>
                <p className="text-sm text-gray-500">
                    Status:{" "}
                    <span
                        className={
                            item.availability_status
                                ? "text-green-600 font-semibold"
                                : "text-red-600 font-semibold"
                        }
                    >
                        {item.availability_status ? "Available" : "Out of stock"}
                    </span>
                </p>

                <form className="mt-6 space-y-4" onSubmit={handleSubmit}>
                    <div>
                        <label
                            htmlFor="quantity"
                            className="block text-sm font-semibold text-gray-600"
                        >
                            Quantity
                        </label>
                        <input
                            type="number"
                            name="quantity"
                            id="quantity"
                            value={orderData.quantity}
                            onChange={handleInputChange}
                            min="1"
                            className="w-full p-3 border rounded-lg text-black"
                            required
                        />
                    </div>

                    <div>
                        <label className="block text-sm font-semibold text-gray-600 mb-2">
                            Customizations
                        </label>
                        <div className="flex flex-wrap gap-4">
                            {availableCustomizations.map((option) => (
                                <label
                                    key={option}
                                    className="flex items-center space-x-2 text-black"
                                >
                                    <input
                                        type="checkbox"
                                        className="h-5 w-5 accent-white border-2 border-gray-400 rounded-sm checked:bg-white checked:border-blue-600"
                                        checked={orderData.customizations.includes(option)}
                                        onChange={() => handleCheckboxChange(option)}
                                    />
                                    <span className="text-black">{option}</span>
                                </label>
                            ))}
                        </div>
                    </div>

                    <div className="pt-4 flex gap-4">
                        <button
                            type="submit"
                            className="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition"
                        >
                            Submit Order
                        </button>
                        <button
                            type="button"
                            onClick={() => navigate("/dashboard")}
                            className="px-6 py-3 bg-gray-400 text-white rounded-lg hover:bg-gray-500 transition"
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </>
    );
};

export default MenuItem;
